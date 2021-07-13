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

	v1alpha1 "kubeform.dev/provider-aws-api/apis/cognito/v1alpha1"
	scheme "kubeform.dev/provider-aws-api/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// UserPoolClientsGetter has a method to return a UserPoolClientInterface.
// A group's client should implement this interface.
type UserPoolClientsGetter interface {
	UserPoolClients(namespace string) UserPoolClientInterface
}

// UserPoolClientInterface has methods to work with UserPoolClient resources.
type UserPoolClientInterface interface {
	Create(ctx context.Context, userPoolClient *v1alpha1.UserPoolClient, opts v1.CreateOptions) (*v1alpha1.UserPoolClient, error)
	Update(ctx context.Context, userPoolClient *v1alpha1.UserPoolClient, opts v1.UpdateOptions) (*v1alpha1.UserPoolClient, error)
	UpdateStatus(ctx context.Context, userPoolClient *v1alpha1.UserPoolClient, opts v1.UpdateOptions) (*v1alpha1.UserPoolClient, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.UserPoolClient, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.UserPoolClientList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.UserPoolClient, err error)
	UserPoolClientExpansion
}

// userPoolClients implements UserPoolClientInterface
type userPoolClients struct {
	client rest.Interface
	ns     string
}

// newUserPoolClients returns a UserPoolClients
func newUserPoolClients(c *CognitoV1alpha1Client, namespace string) *userPoolClients {
	return &userPoolClients{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the userPoolClient, and returns the corresponding userPoolClient object, and an error if there is any.
func (c *userPoolClients) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.UserPoolClient, err error) {
	result = &v1alpha1.UserPoolClient{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("userpoolclients").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of UserPoolClients that match those selectors.
func (c *userPoolClients) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.UserPoolClientList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.UserPoolClientList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("userpoolclients").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested userPoolClients.
func (c *userPoolClients) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("userpoolclients").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a userPoolClient and creates it.  Returns the server's representation of the userPoolClient, and an error, if there is any.
func (c *userPoolClients) Create(ctx context.Context, userPoolClient *v1alpha1.UserPoolClient, opts v1.CreateOptions) (result *v1alpha1.UserPoolClient, err error) {
	result = &v1alpha1.UserPoolClient{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("userpoolclients").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(userPoolClient).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a userPoolClient and updates it. Returns the server's representation of the userPoolClient, and an error, if there is any.
func (c *userPoolClients) Update(ctx context.Context, userPoolClient *v1alpha1.UserPoolClient, opts v1.UpdateOptions) (result *v1alpha1.UserPoolClient, err error) {
	result = &v1alpha1.UserPoolClient{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("userpoolclients").
		Name(userPoolClient.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(userPoolClient).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *userPoolClients) UpdateStatus(ctx context.Context, userPoolClient *v1alpha1.UserPoolClient, opts v1.UpdateOptions) (result *v1alpha1.UserPoolClient, err error) {
	result = &v1alpha1.UserPoolClient{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("userpoolclients").
		Name(userPoolClient.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(userPoolClient).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the userPoolClient and deletes it. Returns an error if one occurs.
func (c *userPoolClients) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("userpoolclients").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *userPoolClients) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("userpoolclients").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched userPoolClient.
func (c *userPoolClients) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.UserPoolClient, err error) {
	result = &v1alpha1.UserPoolClient{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("userpoolclients").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
