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

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"time"

	v1alpha1 "kubedb.dev/apimachinery/apis/dba/v1alpha1"
	scheme "kubedb.dev/apimachinery/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// MongoDBModificationRequestsGetter has a method to return a MongoDBModificationRequestInterface.
// A group's client should implement this interface.
type MongoDBModificationRequestsGetter interface {
	MongoDBModificationRequests() MongoDBModificationRequestInterface
}

// MongoDBModificationRequestInterface has methods to work with MongoDBModificationRequest resources.
type MongoDBModificationRequestInterface interface {
	Create(*v1alpha1.MongoDBModificationRequest) (*v1alpha1.MongoDBModificationRequest, error)
	Update(*v1alpha1.MongoDBModificationRequest) (*v1alpha1.MongoDBModificationRequest, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.MongoDBModificationRequest, error)
	List(opts v1.ListOptions) (*v1alpha1.MongoDBModificationRequestList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.MongoDBModificationRequest, err error)
	MongoDBModificationRequestExpansion
}

// mongoDBModificationRequests implements MongoDBModificationRequestInterface
type mongoDBModificationRequests struct {
	client rest.Interface
}

// newMongoDBModificationRequests returns a MongoDBModificationRequests
func newMongoDBModificationRequests(c *DbaV1alpha1Client) *mongoDBModificationRequests {
	return &mongoDBModificationRequests{
		client: c.RESTClient(),
	}
}

// Get takes name of the mongoDBModificationRequest, and returns the corresponding mongoDBModificationRequest object, and an error if there is any.
func (c *mongoDBModificationRequests) Get(name string, options v1.GetOptions) (result *v1alpha1.MongoDBModificationRequest, err error) {
	result = &v1alpha1.MongoDBModificationRequest{}
	err = c.client.Get().
		Resource("mongodbmodificationrequests").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of MongoDBModificationRequests that match those selectors.
func (c *mongoDBModificationRequests) List(opts v1.ListOptions) (result *v1alpha1.MongoDBModificationRequestList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.MongoDBModificationRequestList{}
	err = c.client.Get().
		Resource("mongodbmodificationrequests").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested mongoDBModificationRequests.
func (c *mongoDBModificationRequests) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("mongodbmodificationrequests").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a mongoDBModificationRequest and creates it.  Returns the server's representation of the mongoDBModificationRequest, and an error, if there is any.
func (c *mongoDBModificationRequests) Create(mongoDBModificationRequest *v1alpha1.MongoDBModificationRequest) (result *v1alpha1.MongoDBModificationRequest, err error) {
	result = &v1alpha1.MongoDBModificationRequest{}
	err = c.client.Post().
		Resource("mongodbmodificationrequests").
		Body(mongoDBModificationRequest).
		Do().
		Into(result)
	return
}

// Update takes the representation of a mongoDBModificationRequest and updates it. Returns the server's representation of the mongoDBModificationRequest, and an error, if there is any.
func (c *mongoDBModificationRequests) Update(mongoDBModificationRequest *v1alpha1.MongoDBModificationRequest) (result *v1alpha1.MongoDBModificationRequest, err error) {
	result = &v1alpha1.MongoDBModificationRequest{}
	err = c.client.Put().
		Resource("mongodbmodificationrequests").
		Name(mongoDBModificationRequest.Name).
		Body(mongoDBModificationRequest).
		Do().
		Into(result)
	return
}

// Delete takes name of the mongoDBModificationRequest and deletes it. Returns an error if one occurs.
func (c *mongoDBModificationRequests) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("mongodbmodificationrequests").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *mongoDBModificationRequests) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("mongodbmodificationrequests").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched mongoDBModificationRequest.
func (c *mongoDBModificationRequests) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.MongoDBModificationRequest, err error) {
	result = &v1alpha1.MongoDBModificationRequest{}
	err = c.client.Patch(pt).
		Resource("mongodbmodificationrequests").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
