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

	v1alpha1 "kubeform.dev/provider-aws-api/apis/pinpoint/v1alpha1"
	scheme "kubeform.dev/provider-aws-api/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ApnsChannelsGetter has a method to return a ApnsChannelInterface.
// A group's client should implement this interface.
type ApnsChannelsGetter interface {
	ApnsChannels(namespace string) ApnsChannelInterface
}

// ApnsChannelInterface has methods to work with ApnsChannel resources.
type ApnsChannelInterface interface {
	Create(ctx context.Context, apnsChannel *v1alpha1.ApnsChannel, opts v1.CreateOptions) (*v1alpha1.ApnsChannel, error)
	Update(ctx context.Context, apnsChannel *v1alpha1.ApnsChannel, opts v1.UpdateOptions) (*v1alpha1.ApnsChannel, error)
	UpdateStatus(ctx context.Context, apnsChannel *v1alpha1.ApnsChannel, opts v1.UpdateOptions) (*v1alpha1.ApnsChannel, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.ApnsChannel, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.ApnsChannelList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ApnsChannel, err error)
	ApnsChannelExpansion
}

// apnsChannels implements ApnsChannelInterface
type apnsChannels struct {
	client rest.Interface
	ns     string
}

// newApnsChannels returns a ApnsChannels
func newApnsChannels(c *PinpointV1alpha1Client, namespace string) *apnsChannels {
	return &apnsChannels{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the apnsChannel, and returns the corresponding apnsChannel object, and an error if there is any.
func (c *apnsChannels) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ApnsChannel, err error) {
	result = &v1alpha1.ApnsChannel{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("apnschannels").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ApnsChannels that match those selectors.
func (c *apnsChannels) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ApnsChannelList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ApnsChannelList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("apnschannels").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested apnsChannels.
func (c *apnsChannels) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("apnschannels").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a apnsChannel and creates it.  Returns the server's representation of the apnsChannel, and an error, if there is any.
func (c *apnsChannels) Create(ctx context.Context, apnsChannel *v1alpha1.ApnsChannel, opts v1.CreateOptions) (result *v1alpha1.ApnsChannel, err error) {
	result = &v1alpha1.ApnsChannel{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("apnschannels").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(apnsChannel).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a apnsChannel and updates it. Returns the server's representation of the apnsChannel, and an error, if there is any.
func (c *apnsChannels) Update(ctx context.Context, apnsChannel *v1alpha1.ApnsChannel, opts v1.UpdateOptions) (result *v1alpha1.ApnsChannel, err error) {
	result = &v1alpha1.ApnsChannel{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("apnschannels").
		Name(apnsChannel.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(apnsChannel).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *apnsChannels) UpdateStatus(ctx context.Context, apnsChannel *v1alpha1.ApnsChannel, opts v1.UpdateOptions) (result *v1alpha1.ApnsChannel, err error) {
	result = &v1alpha1.ApnsChannel{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("apnschannels").
		Name(apnsChannel.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(apnsChannel).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the apnsChannel and deletes it. Returns an error if one occurs.
func (c *apnsChannels) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("apnschannels").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *apnsChannels) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("apnschannels").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched apnsChannel.
func (c *apnsChannels) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ApnsChannel, err error) {
	result = &v1alpha1.ApnsChannel{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("apnschannels").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}