// Copyright 2022 James Cote
// All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by gen-esign; DO NOT EDIT.

// Package tasklists implements the DocuSign SDK
// category TaskLists.
//
// Your administrator may also have created custom task list templates that can be added to rooms. If your administrator created room templates, those room templates may include task lists for you to use.
//
// Service Api documentation may be found at:
// https://developers.docusign.com/docs/rooms-api/reference/TaskLists
// Usage example:
//
//	import (
//	    "github.com/ConsultingMD/esign"
//	    "github.com/ConsultingMD/esign/rooms"
//	)
//	...
//	tasklistsService := tasklists.New(esignCredential)
package tasklists // import "github.com/ConsultingMD/esignrooms//tasklists"

import (
	"context"
	"fmt"
	"net/url"
	"strings"

	"github.com/ConsultingMD/esign"
	"github.com/ConsultingMD/esign/rooms"
)

// Service implements DocuSign TaskLists API operations
type Service struct {
	credential esign.Credential
}

// New initializes a tasklists service using cred to authorize ops.
func New(cred esign.Credential) *Service {
	return &Service{credential: cred}
}

// GetTaskListTemplates gets task list templates.
//
// https://developers.docusign.com/docs/rooms-api/reference/tasklists/tasklisttemplates/gettasklisttemplates
//
// SDK Method TaskLists::GetTaskListTemplates
func (s *Service) GetTaskListTemplates() *GetTaskListTemplatesOp {
	return &GetTaskListTemplatesOp{
		Credential: s.credential,
		Method:     "GET",
		Path:       "task_list_templates",
		Accept:     "application/json",
		QueryOpts:  make(url.Values),
		Version:    esign.RoomsV2,
	}
}

// GetTaskListTemplatesOp implements DocuSign API SDK TaskLists::GetTaskListTemplates
type GetTaskListTemplatesOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *GetTaskListTemplatesOp) Do(ctx context.Context) (*rooms.TaskListTemplateList, error) {
	var res *rooms.TaskListTemplateList
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// StartPosition (Optional) The starting zero-based index position from which to start returning values. The default is `0`.
func (op *GetTaskListTemplatesOp) StartPosition(val int) *GetTaskListTemplatesOp {
	if op != nil {
		op.QueryOpts.Set("startPosition", fmt.Sprintf("%d", val))
	}
	return op
}

// Count (Optional) The number of results to return. This value must be a number between `1` and `100` (default).
func (op *GetTaskListTemplatesOp) Count(val int) *GetTaskListTemplatesOp {
	if op != nil {
		op.QueryOpts.Set("count", fmt.Sprintf("%d", val))
	}
	return op
}

// CreateTaskList applies a task list to a room.
//
// https://developers.docusign.com/docs/rooms-api/reference/tasklists/tasklists/createtasklist
//
// SDK Method TaskLists::CreateTaskList
func (s *Service) CreateTaskList(roomID string, body *rooms.TaskListForCreate) *CreateTaskListOp {
	return &CreateTaskListOp{
		Credential: s.credential,
		Method:     "POST",
		Path:       strings.Join([]string{"rooms", roomID, "task_lists"}, "/"),
		Payload:    body,
		Accept:     "application/json-patch+json, application/json, text/json, application/*+json",
		QueryOpts:  make(url.Values),
		Version:    esign.RoomsV2,
	}
}

// CreateTaskListOp implements DocuSign API SDK TaskLists::CreateTaskList
type CreateTaskListOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *CreateTaskListOp) Do(ctx context.Context) (*rooms.TaskList, error) {
	var res *rooms.TaskList
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// DeleteTaskList deletes a task list from a room.
//
// https://developers.docusign.com/docs/rooms-api/reference/tasklists/tasklists/deletetasklist
//
// SDK Method TaskLists::DeleteTaskList
func (s *Service) DeleteTaskList(taskListID string) *DeleteTaskListOp {
	return &DeleteTaskListOp{
		Credential: s.credential,
		Method:     "DELETE",
		Path:       strings.Join([]string{"task_lists", taskListID}, "/"),
		QueryOpts:  make(url.Values),
		Version:    esign.RoomsV2,
	}
}

// DeleteTaskListOp implements DocuSign API SDK TaskLists::DeleteTaskList
type DeleteTaskListOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *DeleteTaskListOp) Do(ctx context.Context) error {
	return ((*esign.Op)(op)).Do(ctx, nil)
}

// GetTaskLists gets task lists for a room.
//
// https://developers.docusign.com/docs/rooms-api/reference/tasklists/tasklists/gettasklists
//
// SDK Method TaskLists::GetTaskLists
func (s *Service) GetTaskLists(roomID string) *GetTaskListsOp {
	return &GetTaskListsOp{
		Credential: s.credential,
		Method:     "GET",
		Path:       strings.Join([]string{"rooms", roomID, "task_lists"}, "/"),
		Accept:     "application/json",
		QueryOpts:  make(url.Values),
		Version:    esign.RoomsV2,
	}
}

// GetTaskListsOp implements DocuSign API SDK TaskLists::GetTaskLists
type GetTaskListsOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *GetTaskListsOp) Do(ctx context.Context) (*rooms.TaskListSummaryList, error) {
	var res *rooms.TaskListSummaryList
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}
