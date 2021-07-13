/*
Copyright AppsCode Inc. and Contributors

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
	"context"
	"time"

	v1alpha1 "kubeform.dev/provider-aws-api/apis/emr/v1alpha1"
	scheme "kubeform.dev/provider-aws-api/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ManagedScalingPoliciesGetter has a method to return a ManagedScalingPolicyInterface.
// A group's client should implement this interface.
type ManagedScalingPoliciesGetter interface {
	ManagedScalingPolicies(namespace string) ManagedScalingPolicyInterface
}

// ManagedScalingPolicyInterface has methods to work with ManagedScalingPolicy resources.
type ManagedScalingPolicyInterface interface {
	Create(ctx context.Context, managedScalingPolicy *v1alpha1.ManagedScalingPolicy, opts v1.CreateOptions) (*v1alpha1.ManagedScalingPolicy, error)
	Update(ctx context.Context, managedScalingPolicy *v1alpha1.ManagedScalingPolicy, opts v1.UpdateOptions) (*v1alpha1.ManagedScalingPolicy, error)
	UpdateStatus(ctx context.Context, managedScalingPolicy *v1alpha1.ManagedScalingPolicy, opts v1.UpdateOptions) (*v1alpha1.ManagedScalingPolicy, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.ManagedScalingPolicy, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.ManagedScalingPolicyList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ManagedScalingPolicy, err error)
	ManagedScalingPolicyExpansion
}

// managedScalingPolicies implements ManagedScalingPolicyInterface
type managedScalingPolicies struct {
	client rest.Interface
	ns     string
}

// newManagedScalingPolicies returns a ManagedScalingPolicies
func newManagedScalingPolicies(c *EmrV1alpha1Client, namespace string) *managedScalingPolicies {
	return &managedScalingPolicies{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the managedScalingPolicy, and returns the corresponding managedScalingPolicy object, and an error if there is any.
func (c *managedScalingPolicies) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ManagedScalingPolicy, err error) {
	result = &v1alpha1.ManagedScalingPolicy{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("managedscalingpolicies").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ManagedScalingPolicies that match those selectors.
func (c *managedScalingPolicies) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ManagedScalingPolicyList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ManagedScalingPolicyList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("managedscalingpolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested managedScalingPolicies.
func (c *managedScalingPolicies) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("managedscalingpolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a managedScalingPolicy and creates it.  Returns the server's representation of the managedScalingPolicy, and an error, if there is any.
func (c *managedScalingPolicies) Create(ctx context.Context, managedScalingPolicy *v1alpha1.ManagedScalingPolicy, opts v1.CreateOptions) (result *v1alpha1.ManagedScalingPolicy, err error) {
	result = &v1alpha1.ManagedScalingPolicy{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("managedscalingpolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(managedScalingPolicy).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a managedScalingPolicy and updates it. Returns the server's representation of the managedScalingPolicy, and an error, if there is any.
func (c *managedScalingPolicies) Update(ctx context.Context, managedScalingPolicy *v1alpha1.ManagedScalingPolicy, opts v1.UpdateOptions) (result *v1alpha1.ManagedScalingPolicy, err error) {
	result = &v1alpha1.ManagedScalingPolicy{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("managedscalingpolicies").
		Name(managedScalingPolicy.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(managedScalingPolicy).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *managedScalingPolicies) UpdateStatus(ctx context.Context, managedScalingPolicy *v1alpha1.ManagedScalingPolicy, opts v1.UpdateOptions) (result *v1alpha1.ManagedScalingPolicy, err error) {
	result = &v1alpha1.ManagedScalingPolicy{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("managedscalingpolicies").
		Name(managedScalingPolicy.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(managedScalingPolicy).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the managedScalingPolicy and deletes it. Returns an error if one occurs.
func (c *managedScalingPolicies) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("managedscalingpolicies").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *managedScalingPolicies) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("managedscalingpolicies").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched managedScalingPolicy.
func (c *managedScalingPolicies) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ManagedScalingPolicy, err error) {
	result = &v1alpha1.ManagedScalingPolicy{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("managedscalingpolicies").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}