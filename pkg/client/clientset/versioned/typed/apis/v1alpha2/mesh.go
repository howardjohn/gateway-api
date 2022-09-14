/*
Copyright The Kubernetes Authors.

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

package v1alpha2

import (
	"context"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	v1alpha2 "sigs.k8s.io/gateway-api/apis/v1alpha2"
	scheme "sigs.k8s.io/gateway-api/pkg/client/clientset/versioned/scheme"
)

// MeshesGetter has a method to return a MeshInterface.
// A group's client should implement this interface.
type MeshesGetter interface {
	Meshes() MeshInterface
}

// MeshInterface has methods to work with Mesh resources.
type MeshInterface interface {
	Create(ctx context.Context, mesh *v1alpha2.Mesh, opts v1.CreateOptions) (*v1alpha2.Mesh, error)
	Update(ctx context.Context, mesh *v1alpha2.Mesh, opts v1.UpdateOptions) (*v1alpha2.Mesh, error)
	UpdateStatus(ctx context.Context, mesh *v1alpha2.Mesh, opts v1.UpdateOptions) (*v1alpha2.Mesh, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha2.Mesh, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha2.MeshList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha2.Mesh, err error)
	MeshExpansion
}

// meshes implements MeshInterface
type meshes struct {
	client rest.Interface
}

// newMeshes returns a Meshes
func newMeshes(c *GatewayV1alpha2Client) *meshes {
	return &meshes{
		client: c.RESTClient(),
	}
}

// Get takes name of the mesh, and returns the corresponding mesh object, and an error if there is any.
func (c *meshes) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha2.Mesh, err error) {
	result = &v1alpha2.Mesh{}
	err = c.client.Get().
		Resource("meshes").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Meshes that match those selectors.
func (c *meshes) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha2.MeshList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha2.MeshList{}
	err = c.client.Get().
		Resource("meshes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested meshes.
func (c *meshes) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("meshes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a mesh and creates it.  Returns the server's representation of the mesh, and an error, if there is any.
func (c *meshes) Create(ctx context.Context, mesh *v1alpha2.Mesh, opts v1.CreateOptions) (result *v1alpha2.Mesh, err error) {
	result = &v1alpha2.Mesh{}
	err = c.client.Post().
		Resource("meshes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(mesh).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a mesh and updates it. Returns the server's representation of the mesh, and an error, if there is any.
func (c *meshes) Update(ctx context.Context, mesh *v1alpha2.Mesh, opts v1.UpdateOptions) (result *v1alpha2.Mesh, err error) {
	result = &v1alpha2.Mesh{}
	err = c.client.Put().
		Resource("meshes").
		Name(mesh.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(mesh).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *meshes) UpdateStatus(ctx context.Context, mesh *v1alpha2.Mesh, opts v1.UpdateOptions) (result *v1alpha2.Mesh, err error) {
	result = &v1alpha2.Mesh{}
	err = c.client.Put().
		Resource("meshes").
		Name(mesh.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(mesh).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the mesh and deletes it. Returns an error if one occurs.
func (c *meshes) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("meshes").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *meshes) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("meshes").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched mesh.
func (c *meshes) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha2.Mesh, err error) {
	result = &v1alpha2.Mesh{}
	err = c.client.Patch(pt).
		Resource("meshes").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
