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

	v1alpha1 "kubeform.dev/provider-aws-api/apis/kinesisanalyticsv2/v1alpha1"
	scheme "kubeform.dev/provider-aws-api/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ApplicationSnapshotsGetter has a method to return a ApplicationSnapshotInterface.
// A group's client should implement this interface.
type ApplicationSnapshotsGetter interface {
	ApplicationSnapshots(namespace string) ApplicationSnapshotInterface
}

// ApplicationSnapshotInterface has methods to work with ApplicationSnapshot resources.
type ApplicationSnapshotInterface interface {
	Create(ctx context.Context, applicationSnapshot *v1alpha1.ApplicationSnapshot, opts v1.CreateOptions) (*v1alpha1.ApplicationSnapshot, error)
	Update(ctx context.Context, applicationSnapshot *v1alpha1.ApplicationSnapshot, opts v1.UpdateOptions) (*v1alpha1.ApplicationSnapshot, error)
	UpdateStatus(ctx context.Context, applicationSnapshot *v1alpha1.ApplicationSnapshot, opts v1.UpdateOptions) (*v1alpha1.ApplicationSnapshot, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.ApplicationSnapshot, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.ApplicationSnapshotList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ApplicationSnapshot, err error)
	ApplicationSnapshotExpansion
}

// applicationSnapshots implements ApplicationSnapshotInterface
type applicationSnapshots struct {
	client rest.Interface
	ns     string
}

// newApplicationSnapshots returns a ApplicationSnapshots
func newApplicationSnapshots(c *Kinesisanalyticsv2V1alpha1Client, namespace string) *applicationSnapshots {
	return &applicationSnapshots{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the applicationSnapshot, and returns the corresponding applicationSnapshot object, and an error if there is any.
func (c *applicationSnapshots) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ApplicationSnapshot, err error) {
	result = &v1alpha1.ApplicationSnapshot{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("applicationsnapshots").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ApplicationSnapshots that match those selectors.
func (c *applicationSnapshots) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ApplicationSnapshotList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ApplicationSnapshotList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("applicationsnapshots").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested applicationSnapshots.
func (c *applicationSnapshots) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("applicationsnapshots").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a applicationSnapshot and creates it.  Returns the server's representation of the applicationSnapshot, and an error, if there is any.
func (c *applicationSnapshots) Create(ctx context.Context, applicationSnapshot *v1alpha1.ApplicationSnapshot, opts v1.CreateOptions) (result *v1alpha1.ApplicationSnapshot, err error) {
	result = &v1alpha1.ApplicationSnapshot{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("applicationsnapshots").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(applicationSnapshot).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a applicationSnapshot and updates it. Returns the server's representation of the applicationSnapshot, and an error, if there is any.
func (c *applicationSnapshots) Update(ctx context.Context, applicationSnapshot *v1alpha1.ApplicationSnapshot, opts v1.UpdateOptions) (result *v1alpha1.ApplicationSnapshot, err error) {
	result = &v1alpha1.ApplicationSnapshot{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("applicationsnapshots").
		Name(applicationSnapshot.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(applicationSnapshot).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *applicationSnapshots) UpdateStatus(ctx context.Context, applicationSnapshot *v1alpha1.ApplicationSnapshot, opts v1.UpdateOptions) (result *v1alpha1.ApplicationSnapshot, err error) {
	result = &v1alpha1.ApplicationSnapshot{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("applicationsnapshots").
		Name(applicationSnapshot.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(applicationSnapshot).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the applicationSnapshot and deletes it. Returns an error if one occurs.
func (c *applicationSnapshots) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("applicationsnapshots").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *applicationSnapshots) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("applicationsnapshots").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched applicationSnapshot.
func (c *applicationSnapshots) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ApplicationSnapshot, err error) {
	result = &v1alpha1.ApplicationSnapshot{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("applicationsnapshots").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}