/*
Copyright 2020 The Flux authors

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

package fake

import (
	"context"

	v1beta1 "github.com/fluxcd/flagger/pkg/apis/appmesh/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeVirtualServices implements VirtualServiceInterface
type FakeVirtualServices struct {
	Fake *FakeAppmeshV1beta1
	ns   string
}

var virtualservicesResource = schema.GroupVersionResource{Group: "appmesh.k8s.aws", Version: "v1beta1", Resource: "virtualservices"}

var virtualservicesKind = schema.GroupVersionKind{Group: "appmesh.k8s.aws", Version: "v1beta1", Kind: "VirtualService"}

// Get takes name of the virtualService, and returns the corresponding virtualService object, and an error if there is any.
func (c *FakeVirtualServices) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.VirtualService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(virtualservicesResource, c.ns, name), &v1beta1.VirtualService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.VirtualService), err
}

// List takes label and field selectors, and returns the list of VirtualServices that match those selectors.
func (c *FakeVirtualServices) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.VirtualServiceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(virtualservicesResource, virtualservicesKind, c.ns, opts), &v1beta1.VirtualServiceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.VirtualServiceList{ListMeta: obj.(*v1beta1.VirtualServiceList).ListMeta}
	for _, item := range obj.(*v1beta1.VirtualServiceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested virtualServices.
func (c *FakeVirtualServices) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(virtualservicesResource, c.ns, opts))

}

// Create takes the representation of a virtualService and creates it.  Returns the server's representation of the virtualService, and an error, if there is any.
func (c *FakeVirtualServices) Create(ctx context.Context, virtualService *v1beta1.VirtualService, opts v1.CreateOptions) (result *v1beta1.VirtualService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(virtualservicesResource, c.ns, virtualService), &v1beta1.VirtualService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.VirtualService), err
}

// Update takes the representation of a virtualService and updates it. Returns the server's representation of the virtualService, and an error, if there is any.
func (c *FakeVirtualServices) Update(ctx context.Context, virtualService *v1beta1.VirtualService, opts v1.UpdateOptions) (result *v1beta1.VirtualService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(virtualservicesResource, c.ns, virtualService), &v1beta1.VirtualService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.VirtualService), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeVirtualServices) UpdateStatus(ctx context.Context, virtualService *v1beta1.VirtualService, opts v1.UpdateOptions) (*v1beta1.VirtualService, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(virtualservicesResource, "status", c.ns, virtualService), &v1beta1.VirtualService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.VirtualService), err
}

// Delete takes name of the virtualService and deletes it. Returns an error if one occurs.
func (c *FakeVirtualServices) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(virtualservicesResource, c.ns, name), &v1beta1.VirtualService{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeVirtualServices) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(virtualservicesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.VirtualServiceList{})
	return err
}

// Patch applies the patch and returns the patched virtualService.
func (c *FakeVirtualServices) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.VirtualService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(virtualservicesResource, c.ns, name, pt, data, subresources...), &v1beta1.VirtualService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.VirtualService), err
}
