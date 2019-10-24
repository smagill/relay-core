/*
Copyright (c) Puppet, Inc.
*/

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	nebulapuppetcomv1 "github.com/puppetlabs/nebula-tasks/pkg/apis/nebula.puppet.com/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeWorkflowRuns implements WorkflowRunInterface
type FakeWorkflowRuns struct {
	Fake *FakeNebulaV1
	ns   string
}

var workflowrunsResource = schema.GroupVersionResource{Group: "nebula.puppet.com", Version: "v1", Resource: "workflowruns"}

var workflowrunsKind = schema.GroupVersionKind{Group: "nebula.puppet.com", Version: "v1", Kind: "WorkflowRun"}

// Get takes name of the workflowRun, and returns the corresponding workflowRun object, and an error if there is any.
func (c *FakeWorkflowRuns) Get(name string, options v1.GetOptions) (result *nebulapuppetcomv1.WorkflowRun, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(workflowrunsResource, c.ns, name), &nebulapuppetcomv1.WorkflowRun{})

	if obj == nil {
		return nil, err
	}
	return obj.(*nebulapuppetcomv1.WorkflowRun), err
}

// List takes label and field selectors, and returns the list of WorkflowRuns that match those selectors.
func (c *FakeWorkflowRuns) List(opts v1.ListOptions) (result *nebulapuppetcomv1.WorkflowRunList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(workflowrunsResource, workflowrunsKind, c.ns, opts), &nebulapuppetcomv1.WorkflowRunList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &nebulapuppetcomv1.WorkflowRunList{ListMeta: obj.(*nebulapuppetcomv1.WorkflowRunList).ListMeta}
	for _, item := range obj.(*nebulapuppetcomv1.WorkflowRunList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested workflowRuns.
func (c *FakeWorkflowRuns) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(workflowrunsResource, c.ns, opts))

}

// Create takes the representation of a workflowRun and creates it.  Returns the server's representation of the workflowRun, and an error, if there is any.
func (c *FakeWorkflowRuns) Create(workflowRun *nebulapuppetcomv1.WorkflowRun) (result *nebulapuppetcomv1.WorkflowRun, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(workflowrunsResource, c.ns, workflowRun), &nebulapuppetcomv1.WorkflowRun{})

	if obj == nil {
		return nil, err
	}
	return obj.(*nebulapuppetcomv1.WorkflowRun), err
}

// Update takes the representation of a workflowRun and updates it. Returns the server's representation of the workflowRun, and an error, if there is any.
func (c *FakeWorkflowRuns) Update(workflowRun *nebulapuppetcomv1.WorkflowRun) (result *nebulapuppetcomv1.WorkflowRun, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(workflowrunsResource, c.ns, workflowRun), &nebulapuppetcomv1.WorkflowRun{})

	if obj == nil {
		return nil, err
	}
	return obj.(*nebulapuppetcomv1.WorkflowRun), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeWorkflowRuns) UpdateStatus(workflowRun *nebulapuppetcomv1.WorkflowRun) (*nebulapuppetcomv1.WorkflowRun, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(workflowrunsResource, "status", c.ns, workflowRun), &nebulapuppetcomv1.WorkflowRun{})

	if obj == nil {
		return nil, err
	}
	return obj.(*nebulapuppetcomv1.WorkflowRun), err
}

// Delete takes name of the workflowRun and deletes it. Returns an error if one occurs.
func (c *FakeWorkflowRuns) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(workflowrunsResource, c.ns, name), &nebulapuppetcomv1.WorkflowRun{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeWorkflowRuns) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(workflowrunsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &nebulapuppetcomv1.WorkflowRunList{})
	return err
}

// Patch applies the patch and returns the patched workflowRun.
func (c *FakeWorkflowRuns) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *nebulapuppetcomv1.WorkflowRun, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(workflowrunsResource, c.ns, name, pt, data, subresources...), &nebulapuppetcomv1.WorkflowRun{})

	if obj == nil {
		return nil, err
	}
	return obj.(*nebulapuppetcomv1.WorkflowRun), err
}