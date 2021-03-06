package client

import (
	"github.com/rancher/norman/types"
)

const (
	NamespacedSSHAuthType                 = "namespacedSshAuth"
	NamespacedSSHAuthFieldAnnotations     = "annotations"
	NamespacedSSHAuthFieldCreated         = "created"
	NamespacedSSHAuthFieldCreatorID       = "creatorId"
	NamespacedSSHAuthFieldFinalizers      = "finalizers"
	NamespacedSSHAuthFieldFingerprint     = "certFingerprint"
	NamespacedSSHAuthFieldLabels          = "labels"
	NamespacedSSHAuthFieldName            = "name"
	NamespacedSSHAuthFieldNamespaceId     = "namespaceId"
	NamespacedSSHAuthFieldOwnerReferences = "ownerReferences"
	NamespacedSSHAuthFieldPrivateKey      = "privateKey"
	NamespacedSSHAuthFieldProjectID       = "projectId"
	NamespacedSSHAuthFieldRemoved         = "removed"
	NamespacedSSHAuthFieldUuid            = "uuid"
)

type NamespacedSSHAuth struct {
	types.Resource
	Annotations     map[string]string `json:"annotations,omitempty"`
	Created         string            `json:"created,omitempty"`
	CreatorID       string            `json:"creatorId,omitempty"`
	Finalizers      []string          `json:"finalizers,omitempty"`
	Fingerprint     string            `json:"certFingerprint,omitempty"`
	Labels          map[string]string `json:"labels,omitempty"`
	Name            string            `json:"name,omitempty"`
	NamespaceId     string            `json:"namespaceId,omitempty"`
	OwnerReferences []OwnerReference  `json:"ownerReferences,omitempty"`
	PrivateKey      string            `json:"privateKey,omitempty"`
	ProjectID       string            `json:"projectId,omitempty"`
	Removed         string            `json:"removed,omitempty"`
	Uuid            string            `json:"uuid,omitempty"`
}
type NamespacedSSHAuthCollection struct {
	types.Collection
	Data   []NamespacedSSHAuth `json:"data,omitempty"`
	client *NamespacedSSHAuthClient
}

type NamespacedSSHAuthClient struct {
	apiClient *Client
}

type NamespacedSSHAuthOperations interface {
	List(opts *types.ListOpts) (*NamespacedSSHAuthCollection, error)
	Create(opts *NamespacedSSHAuth) (*NamespacedSSHAuth, error)
	Update(existing *NamespacedSSHAuth, updates interface{}) (*NamespacedSSHAuth, error)
	ByID(id string) (*NamespacedSSHAuth, error)
	Delete(container *NamespacedSSHAuth) error
}

func newNamespacedSSHAuthClient(apiClient *Client) *NamespacedSSHAuthClient {
	return &NamespacedSSHAuthClient{
		apiClient: apiClient,
	}
}

func (c *NamespacedSSHAuthClient) Create(container *NamespacedSSHAuth) (*NamespacedSSHAuth, error) {
	resp := &NamespacedSSHAuth{}
	err := c.apiClient.Ops.DoCreate(NamespacedSSHAuthType, container, resp)
	return resp, err
}

func (c *NamespacedSSHAuthClient) Update(existing *NamespacedSSHAuth, updates interface{}) (*NamespacedSSHAuth, error) {
	resp := &NamespacedSSHAuth{}
	err := c.apiClient.Ops.DoUpdate(NamespacedSSHAuthType, &existing.Resource, updates, resp)
	return resp, err
}

func (c *NamespacedSSHAuthClient) List(opts *types.ListOpts) (*NamespacedSSHAuthCollection, error) {
	resp := &NamespacedSSHAuthCollection{}
	err := c.apiClient.Ops.DoList(NamespacedSSHAuthType, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *NamespacedSSHAuthCollection) Next() (*NamespacedSSHAuthCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &NamespacedSSHAuthCollection{}
		err := cc.client.apiClient.Ops.DoNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *NamespacedSSHAuthClient) ByID(id string) (*NamespacedSSHAuth, error) {
	resp := &NamespacedSSHAuth{}
	err := c.apiClient.Ops.DoByID(NamespacedSSHAuthType, id, resp)
	return resp, err
}

func (c *NamespacedSSHAuthClient) Delete(container *NamespacedSSHAuth) error {
	return c.apiClient.Ops.DoResourceDelete(NamespacedSSHAuthType, &container.Resource)
}
