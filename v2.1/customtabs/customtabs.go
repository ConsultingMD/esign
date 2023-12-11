// Copyright 2022 James Cote
// All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by gen-esign; DO NOT EDIT.

// Package customtabs implements the DocuSign SDK
// category CustomTabs.
//
// Custom Tabs enable accounts to have one or more pre-configured (custom) tabs. Custom tabs save time when users are tagging documents since the users don't have to manually set the tabs' parameters.
//
// This category enables custom tabs to be managed programmatically, including creation, deletion, etc.
//
// Service Api documentation may be found at:
// https://developers.docusign.com/docs/esign-rest-api/reference/CustomTabs
// Usage example:
//
//	import (
//	    "github.com/ConsultingMD/esign"
//	    "github.com/ConsultingMD/esign/v2.1/model"
//	)
//	...
//	customtabsService := customtabs.New(esignCredential)
package customtabs // import "github.com/ConsultingMD/esignv2.1/customtabs"

import (
	"context"
	"net/url"
	"strings"

	"github.com/ConsultingMD/esign"
	"github.com/ConsultingMD/esign/v2.1/model"
)

// Service implements DocuSign CustomTabs API operations
type Service struct {
	credential esign.Credential
}

// New initializes a customtabs service using cred to authorize ops.
func New(cred esign.Credential) *Service {
	return &Service{credential: cred}
}

// Create creates a custom tab.
//
// https://developers.docusign.com/docs/esign-rest-api/reference/customtabs/customtabs/create
//
// SDK Method CustomTabs::create
func (s *Service) Create(tabMetadata *model.TabMetadata) *CreateOp {
	return &CreateOp{
		Credential: s.credential,
		Method:     "POST",
		Path:       "tab_definitions",
		Payload:    tabMetadata,
		QueryOpts:  make(url.Values),
		Version:    esign.APIv21,
	}
}

// CreateOp implements DocuSign API SDK CustomTabs::create
type CreateOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *CreateOp) Do(ctx context.Context) (*model.TabMetadata, error) {
	var res *model.TabMetadata
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// Delete deletes custom tab information.
//
// https://developers.docusign.com/docs/esign-rest-api/reference/customtabs/customtabs/delete
//
// SDK Method CustomTabs::delete
func (s *Service) Delete(customTabID string) *DeleteOp {
	return &DeleteOp{
		Credential: s.credential,
		Method:     "DELETE",
		Path:       strings.Join([]string{"tab_definitions", customTabID}, "/"),
		QueryOpts:  make(url.Values),
		Version:    esign.APIv21,
	}
}

// DeleteOp implements DocuSign API SDK CustomTabs::delete
type DeleteOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *DeleteOp) Do(ctx context.Context) error {
	return ((*esign.Op)(op)).Do(ctx, nil)
}

// Get gets custom tab information.
//
// https://developers.docusign.com/docs/esign-rest-api/reference/customtabs/customtabs/get
//
// SDK Method CustomTabs::get
func (s *Service) Get(customTabID string) *GetOp {
	return &GetOp{
		Credential: s.credential,
		Method:     "GET",
		Path:       strings.Join([]string{"tab_definitions", customTabID}, "/"),
		Accept:     "application/json",
		QueryOpts:  make(url.Values),
		Version:    esign.APIv21,
	}
}

// GetOp implements DocuSign API SDK CustomTabs::get
type GetOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *GetOp) Do(ctx context.Context) (*model.TabMetadata, error) {
	var res *model.TabMetadata
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// List gets a list of all account tabs.
//
// https://developers.docusign.com/docs/esign-rest-api/reference/customtabs/customtabs/list
//
// SDK Method CustomTabs::list
func (s *Service) List() *ListOp {
	return &ListOp{
		Credential: s.credential,
		Method:     "GET",
		Path:       "tab_definitions",
		Accept:     "application/json",
		QueryOpts:  make(url.Values),
		Version:    esign.APIv21,
	}
}

// ListOp implements DocuSign API SDK CustomTabs::list
type ListOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *ListOp) Do(ctx context.Context) (*model.TabMetadataList, error) {
	var res *model.TabMetadataList
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// CustomTabOnly when **true,** only custom tabs are returned in the response.
func (op *ListOp) CustomTabOnly() *ListOp {
	if op != nil {
		op.QueryOpts.Set("custom_tab_only", "true")
	}
	return op
}

// Update updates custom tab information.
//
// https://developers.docusign.com/docs/esign-rest-api/reference/customtabs/customtabs/update
//
// SDK Method CustomTabs::update
func (s *Service) Update(customTabID string, tabMetadata *model.TabMetadata) *UpdateOp {
	return &UpdateOp{
		Credential: s.credential,
		Method:     "PUT",
		Path:       strings.Join([]string{"tab_definitions", customTabID}, "/"),
		Payload:    tabMetadata,
		Accept:     "application/json",
		QueryOpts:  make(url.Values),
		Version:    esign.APIv21,
	}
}

// UpdateOp implements DocuSign API SDK CustomTabs::update
type UpdateOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *UpdateOp) Do(ctx context.Context) (*model.TabMetadata, error) {
	var res *model.TabMetadata
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}
