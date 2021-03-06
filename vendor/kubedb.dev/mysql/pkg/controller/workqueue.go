/*
Copyright The KubeDB Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package controller

import (
	api "kubedb.dev/apimachinery/apis/kubedb/v1alpha1"
	"kubedb.dev/apimachinery/client/clientset/versioned/typed/kubedb/v1alpha1/util"

	"github.com/appscode/go/log"
	core_util "kmodules.xyz/client-go/core/v1"
	"kmodules.xyz/client-go/tools/queue"
)

func (c *Controller) initWatcher() {
	c.myInformer = c.KubedbInformerFactory.Kubedb().V1alpha1().MySQLs().Informer()
	c.myQueue = queue.New("MySQL", c.MaxNumRequeues, c.NumThreads, c.runMySQL)
	c.myLister = c.KubedbInformerFactory.Kubedb().V1alpha1().MySQLs().Lister()
	c.myInformer.AddEventHandler(queue.NewReconcilableHandler(c.myQueue.GetQueue()))
}

func (c *Controller) runMySQL(key string) error {
	log.Debugln("started processing, key:", key)
	obj, exists, err := c.myInformer.GetIndexer().GetByKey(key)
	if err != nil {
		log.Errorf("Fetching object with key %s from store failed with %v", key, err)
		return err
	}

	if !exists {
		log.Debugf("MySQL %s does not exist anymore", key)
	} else {
		// Note that you also have to check the uid if you have a local controlled resource, which
		// is dependent on the actual instance, to detect that a MySQL was recreated with the same name
		mysql := obj.(*api.MySQL).DeepCopy()
		if mysql.DeletionTimestamp != nil {
			if core_util.HasFinalizer(mysql.ObjectMeta, api.GenericKey) {
				if err := c.terminate(mysql); err != nil {
					log.Errorln(err)
					return err
				}
				_, _, err = util.PatchMySQL(c.ExtClient.KubedbV1alpha1(), mysql, func(in *api.MySQL) *api.MySQL {
					in.ObjectMeta = core_util.RemoveFinalizer(in.ObjectMeta, api.GenericKey)
					return in
				})
				return err
			}
		} else {
			mysql, _, err = util.PatchMySQL(c.ExtClient.KubedbV1alpha1(), mysql, func(in *api.MySQL) *api.MySQL {
				in.ObjectMeta = core_util.AddFinalizer(in.ObjectMeta, api.GenericKey)
				return in
			})
			if err != nil {
				return err
			}

			if mysql.Spec.Paused {
				return nil
			}

			if mysql.Spec.Halted {
				if err := c.halt(mysql); err != nil {
					log.Errorln(err)
					c.pushFailureEvent(mysql, err.Error())
					return err
				}
			} else {
				if err := c.create(mysql); err != nil {
					log.Errorln(err)
					c.pushFailureEvent(mysql, err.Error())
					return err
				}
			}
		}
	}
	return nil
}
