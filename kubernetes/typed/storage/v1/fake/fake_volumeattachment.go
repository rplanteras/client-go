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

package fake

import (
	"context"

	storagev1 "k8s.io/api/storage/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeVolumeAttachments implements VolumeAttachmentInterface
type FakeVolumeAttachments struct {
	Fake *FakeStorageV1
}

var volumeattachmentsResource = schema.GroupVersionResource{Group: "storage.k8s.io", Version: "v1", Resource: "volumeattachments"}

var volumeattachmentsKind = schema.GroupVersionKind{Group: "storage.k8s.io", Version: "v1", Kind: "VolumeAttachment"}

// Get takes name of the volumeAttachment, and returns the corresponding volumeAttachment object, and an error if there is any.
func (c *FakeVolumeAttachments) Get(ctx context.Context, name string, options v1.GetOptions) (result *storagev1.VolumeAttachment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(volumeattachmentsResource, name), &storagev1.VolumeAttachment{})
	if obj == nil {
		return nil, err
	}
	return obj.(*storagev1.VolumeAttachment), err
}

// List takes label and field selectors, and returns the list of VolumeAttachments that match those selectors.
func (c *FakeVolumeAttachments) List(ctx context.Context, opts v1.ListOptions) (result *storagev1.VolumeAttachmentList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(volumeattachmentsResource, volumeattachmentsKind, opts), &storagev1.VolumeAttachmentList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &storagev1.VolumeAttachmentList{ListMeta: obj.(*storagev1.VolumeAttachmentList).ListMeta}
	for _, item := range obj.(*storagev1.VolumeAttachmentList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested volumeAttachments.
func (c *FakeVolumeAttachments) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(volumeattachmentsResource, opts))
}

// Create takes the representation of a volumeAttachment and creates it.  Returns the server's representation of the volumeAttachment, and an error, if there is any.
func (c *FakeVolumeAttachments) Create(ctx context.Context, volumeAttachment *storagev1.VolumeAttachment, opts v1.CreateOptions) (result *storagev1.VolumeAttachment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(volumeattachmentsResource, volumeAttachment), &storagev1.VolumeAttachment{})
	if obj == nil {
		return nil, err
	}
	return obj.(*storagev1.VolumeAttachment), err
}

// Update takes the representation of a volumeAttachment and updates it. Returns the server's representation of the volumeAttachment, and an error, if there is any.
func (c *FakeVolumeAttachments) Update(ctx context.Context, volumeAttachment *storagev1.VolumeAttachment, opts v1.UpdateOptions) (result *storagev1.VolumeAttachment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(volumeattachmentsResource, volumeAttachment), &storagev1.VolumeAttachment{})
	if obj == nil {
		return nil, err
	}
	return obj.(*storagev1.VolumeAttachment), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeVolumeAttachments) UpdateStatus(ctx context.Context, volumeAttachment *storagev1.VolumeAttachment, opts v1.UpdateOptions) (*storagev1.VolumeAttachment, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(volumeattachmentsResource, "status", volumeAttachment), &storagev1.VolumeAttachment{})
	if obj == nil {
		return nil, err
	}
	return obj.(*storagev1.VolumeAttachment), err
}

// Delete takes name of the volumeAttachment and deletes it. Returns an error if one occurs.
func (c *FakeVolumeAttachments) Delete(ctx context.Context, name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(volumeattachmentsResource, name), &storagev1.VolumeAttachment{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeVolumeAttachments) DeleteCollection(ctx context.Context, options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(volumeattachmentsResource, listOptions)

	_, err := c.Fake.Invokes(action, &storagev1.VolumeAttachmentList{})
	return err
}

// Patch applies the patch and returns the patched volumeAttachment.
func (c *FakeVolumeAttachments) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *storagev1.VolumeAttachment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(volumeattachmentsResource, name, pt, data, subresources...), &storagev1.VolumeAttachment{})
	if obj == nil {
		return nil, err
	}
	return obj.(*storagev1.VolumeAttachment), err
}
