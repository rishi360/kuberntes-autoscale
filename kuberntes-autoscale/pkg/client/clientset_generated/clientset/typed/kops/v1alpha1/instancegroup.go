/*
Copyright 2018 The Kubernetes Authors.

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
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	v1alpha1 "k8s.io/kops/pkg/apis/kops/v1alpha1"
	scheme "k8s.io/kops/pkg/client/clientset_generated/clientset/scheme"
)

// InstanceGroupsGetter has a method to return a InstanceGroupInterface.
// A group's client should implement this interface.
type InstanceGroupsGetter interface {
	InstanceGroups(namespace string) InstanceGroupInterface
}

// InstanceGroupInterface has methods to work with InstanceGroup resources.
type InstanceGroupInterface interface {
	Create(*v1alpha1.InstanceGroup) (*v1alpha1.InstanceGroup, error)
	Update(*v1alpha1.InstanceGroup) (*v1alpha1.InstanceGroup, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.InstanceGroup, error)
	List(opts v1.ListOptions) (*v1alpha1.InstanceGroupList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.InstanceGroup, err error)
	InstanceGroupExpansion
}

// instanceGroups implements InstanceGroupInterface
type instanceGroups struct {
	client rest.Interface
	ns     string
}

// newInstanceGroups returns a InstanceGroups
func newInstanceGroups(c *KopsV1alpha1Client, namespace string) *instanceGroups {
	return &instanceGroups{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the instanceGroup, and returns the corresponding instanceGroup object, and an error if there is any.
func (c *instanceGroups) Get(name string, options v1.GetOptions) (result *v1alpha1.InstanceGroup, err error) {
	result = &v1alpha1.InstanceGroup{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("instancegroups").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of InstanceGroups that match those selectors.
func (c *instanceGroups) List(opts v1.ListOptions) (result *v1alpha1.InstanceGroupList, err error) {
	result = &v1alpha1.InstanceGroupList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("instancegroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested instanceGroups.
func (c *instanceGroups) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("instancegroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a instanceGroup and creates it.  Returns the server's representation of the instanceGroup, and an error, if there is any.
func (c *instanceGroups) Create(instanceGroup *v1alpha1.InstanceGroup) (result *v1alpha1.InstanceGroup, err error) {
	result = &v1alpha1.InstanceGroup{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("instancegroups").
		Body(instanceGroup).
		Do().
		Into(result)
	return
}

// Update takes the representation of a instanceGroup and updates it. Returns the server's representation of the instanceGroup, and an error, if there is any.
func (c *instanceGroups) Update(instanceGroup *v1alpha1.InstanceGroup) (result *v1alpha1.InstanceGroup, err error) {
	result = &v1alpha1.InstanceGroup{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("instancegroups").
		Name(instanceGroup.Name).
		Body(instanceGroup).
		Do().
		Into(result)
	return
}

// Delete takes name of the instanceGroup and deletes it. Returns an error if one occurs.
func (c *instanceGroups) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("instancegroups").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *instanceGroups) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("instancegroups").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched instanceGroup.
func (c *instanceGroups) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.InstanceGroup, err error) {
	result = &v1alpha1.InstanceGroup{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("instancegroups").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
