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

	"github.com/appscode/go/types"
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (c *Controller) getVolumeAndVolumeMountForDefaultUserList(pgbouncer *api.PgBouncer) (*core.Volume, *core.VolumeMount, error) {
	fSecret := c.GetDefaultSecretSpec(pgbouncer)
	_, err := c.Client.CoreV1().Secrets(fSecret.Namespace).Get(fSecret.Name, metav1.GetOptions{})
	if err != nil {
		return nil, nil, err
	}

	secretVolume := &core.Volume{
		Name: "fallback-userlist",
		VolumeSource: core.VolumeSource{
			Secret: &core.SecretVolumeSource{
				SecretName: fSecret.Name,
			},
		},
	}
	//Add to volumeMounts to mount the volume
	secretVolumeMount := &core.VolumeMount{
		Name:      "fallback-userlist",
		MountPath: UserListMountPath,
		ReadOnly:  true,
	}

	return secretVolume, secretVolumeMount, nil
}

func (c *Controller) getVolumeAndVolumeMountForServingServerCertificate(pgbouncer *api.PgBouncer) (*core.Volume, *core.VolumeMount, error) {
	//TODO: this is for issuer only, I'm not sure about clusterIssuer yet
	clientSecret, err := c.Client.CoreV1().Secrets(pgbouncer.Namespace).Get(pgbouncer.Name+api.PgBouncerServingServerSuffix, metav1.GetOptions{})
	if err != nil {
		return nil, nil, err
	}
	secretVolume := &core.Volume{
		Name: clientSecret.Name,
		VolumeSource: core.VolumeSource{
			Secret: &core.SecretVolumeSource{
				SecretName: clientSecret.Name,
			},
		},
	}
	//Add to volumeMounts to mount the volume
	secretVolumeMount := &core.VolumeMount{
		Name:      clientSecret.Name,
		MountPath: ServingServerCertMountPath,
		ReadOnly:  true,
	}

	return secretVolume, secretVolumeMount, nil
}

func (c *Controller) getVolumeAndVolumeMountForServingClientCertificate(pgbouncer *api.PgBouncer) (*core.Volume, *core.VolumeMount, error) {
	//TODO: this is for issuer only, I'm not sure about clusterIssuer yet
	clientSecret, err := c.Client.CoreV1().Secrets(pgbouncer.Namespace).Get(pgbouncer.Name+api.PgBouncerServingClientSuffix, metav1.GetOptions{})
	if err != nil {
		return nil, nil, err
	}
	secretVolume := &core.Volume{
		Name: clientSecret.Name,
		VolumeSource: core.VolumeSource{
			Secret: &core.SecretVolumeSource{
				SecretName:  clientSecret.Name,
				DefaultMode: types.Int32P(0600),
			},
		},
	}
	//Add to volumeMounts to mount the volume
	secretVolumeMount := &core.VolumeMount{
		Name:      clientSecret.Name,
		MountPath: ServingClientCertMountPath,
	}

	return secretVolume, secretVolumeMount, nil
}

func (c *Controller) getVolumeAndVolumeMountForExporterClientCertificate(pgbouncer *api.PgBouncer) (*core.Volume, *core.VolumeMount, error) {
	clientSecret, err := c.Client.CoreV1().Secrets(pgbouncer.Namespace).Get(pgbouncer.Name+api.PgBouncerExporterClientCertSuffix, metav1.GetOptions{})
	if err != nil {
		return nil, nil, err
	}
	secretVolume := &core.Volume{
		Name: clientSecret.Name,
		VolumeSource: core.VolumeSource{
			Secret: &core.SecretVolumeSource{
				SecretName:  clientSecret.Name,
				DefaultMode: types.Int32P(0600),
			},
		},
	}
	//Add to volumeMounts to mount the volume
	secretVolumeMount := &core.VolumeMount{
		Name:      clientSecret.Name,
		MountPath: ServingClientCertMountPath,
	}

	return secretVolume, secretVolumeMount, nil
}
