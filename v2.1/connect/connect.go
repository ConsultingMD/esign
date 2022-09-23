// Copyright 2019 James Cote
// All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by gen-esign; DO NOT EDIT.

// Package connect implements the DocuSign SDK
// category Connect.
//
// The Connect service enables your application to be called via
// HTTPS when an event of interest occurs.
//
// Use the Connect service to "end the polling madness." With
// Connect, there is no need for your application to poll DocuSign
// every 15 minutes to learn the latest about your envelopes.
//
// Instead, you register your interest in one or more types of
// envelope or recipient events. Then, when an interesting event
// occurs, the DocuSign platform will contact your application with
// the event's details and data. You can register interest in
// envelopes sent by particular users in your account, or for
// envelopes sent by any user.
//
// Connect can empower your organization to manage document actions
// as they occur, and allows you to track their changes within your
// own systems. Upon completion, envelope information, including
// document content, can be stored in your own databases or CMS
// systems, and these events can be triggered via webhooks
// delivering messages to your application.
//
// **Note:** To make API calls to any of the Connect endpoints, you must be an account administrator.
//
// ## Incoming Connect Calls
//
// To use the Connect service, your application needs to provide an
// HTTPS URL that can be called from the public Internet. If your
// application runs on a server behind your organization's firewall,
// then you will need to create a "pinhole" in the firewall to allow
// the incoming Connect calls from DocuSign to reach your
// application. You can also use other techniques such as proxy
// servers and DMZ networking for receiving the incoming calls.
//
// Connect delivers events over HTTP requests in JSON or XML.
// See [DocuSign Connect overview](/platform/webhooks/connect/).
//
// If your application is not configured to accept post messages,
// DocuSign will NOT return an additional post error response to
// your listener application. If you've enabled logging on your
// configuration, it will be logged in Admin under the configuration
// failure log.
//
// ## Per-envelope Connect Configuration
//
// Instead of registering a general Connect configuration and
// listener, an individual envelope can have its own Connect
// configuration. See the
// [`eventNotification`](/docs/esign-rest-api/reference/envelopes/envelopes/create/#schema__envelopedefinition_eventnotification)
// property for envelopes.
//
// ## Historical Publish Endpoint
//
// To submit existing envelopes to an endpoint, use the [EnvelopePublish](/docs/esign-rest-api/reference/envelopes/envelopepublish/) resource.
//
// Service Api documentation may be found at:
// https://developers.docusign.com/docs/esign-rest-api/reference/Connect
// Usage example:
//
//   import (
//       "github.com/jfcote87/esign"
//       "github.com/jfcote87/esign/v2.1/connect"
//       "github.com/jfcote87/esign/v2.1/model"
//   )
//   ...
//   connectService := connect.New(esignCredential)
package connect // import "github.com/jfcote87/esign/v2.1//connect"

import (
	"context"
	"fmt"
	"io"
	"net/url"
	"strings"
	"time"

	"github.com/jfcote87/esign"
	"github.com/jfcote87/esign/v2.1/model"
)

// Service implements DocuSign Connect Category API operations
type Service struct {
	credential esign.Credential
}

// New initializes a connect service using cred to authorize ops.
func New(cred esign.Credential) *Service {
	return &Service{credential: cred}
}

// ConfigurationsCreate creates a Connect configuration.
//
// https://developers.docusign.com/docs/esign-rest-api/reference/connect/connectconfigurations/create
//
// SDK Method Connect::createConfiguration
func (s *Service) ConfigurationsCreate(connectCustomConfiguration *model.ConnectCustomConfiguration) *ConfigurationsCreateOp {
	return &ConfigurationsCreateOp{
		Credential: s.credential,
		Method:     "POST",
		Path:       "connect",
		Payload:    connectCustomConfiguration,
		QueryOpts:  make(url.Values),
		Version:    esign.APIv21,
	}
}

// ConfigurationsCreateOp implements DocuSign API SDK Connect::createConfiguration
type ConfigurationsCreateOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *ConfigurationsCreateOp) Do(ctx context.Context) (*model.ConnectCustomConfiguration, error) {
	var res *model.ConnectCustomConfiguration
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// ConfigurationsDelete deletes the specified Connect configuration.
//
// https://developers.docusign.com/docs/esign-rest-api/reference/connect/connectconfigurations/delete
//
// SDK Method Connect::deleteConfiguration
func (s *Service) ConfigurationsDelete(connectID string) *ConfigurationsDeleteOp {
	return &ConfigurationsDeleteOp{
		Credential: s.credential,
		Method:     "DELETE",
		Path:       strings.Join([]string{"connect", connectID}, "/"),
		QueryOpts:  make(url.Values),
		Version:    esign.APIv21,
	}
}

// ConfigurationsDeleteOp implements DocuSign API SDK Connect::deleteConfiguration
type ConfigurationsDeleteOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *ConfigurationsDeleteOp) Do(ctx context.Context) error {
	return ((*esign.Op)(op)).Do(ctx, nil)
}

// ConfigurationsGet gets the details about a Connect configuration.
//
// https://developers.docusign.com/docs/esign-rest-api/reference/connect/connectconfigurations/get
//
// SDK Method Connect::getConfiguration
func (s *Service) ConfigurationsGet(connectID string) *ConfigurationsGetOp {
	return &ConfigurationsGetOp{
		Credential: s.credential,
		Method:     "GET",
		Path:       strings.Join([]string{"connect", connectID}, "/"),
		Accept:     "application/json",
		QueryOpts:  make(url.Values),
		Version:    esign.APIv21,
	}
}

// ConfigurationsGetOp implements DocuSign API SDK Connect::getConfiguration
type ConfigurationsGetOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *ConfigurationsGetOp) Do(ctx context.Context) (*model.ConnectConfigResults, error) {
	var res *model.ConnectConfigResults
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// ConfigurationsList get Connect configuration information.
//
// https://developers.docusign.com/docs/esign-rest-api/reference/connect/connectconfigurations/list
//
// SDK Method Connect::listConfigurations
func (s *Service) ConfigurationsList() *ConfigurationsListOp {
	return &ConfigurationsListOp{
		Credential: s.credential,
		Method:     "GET",
		Path:       "connect",
		Accept:     "application/json",
		QueryOpts:  make(url.Values),
		Version:    esign.APIv21,
	}
}

// ConfigurationsListOp implements DocuSign API SDK Connect::listConfigurations
type ConfigurationsListOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *ConfigurationsListOp) Do(ctx context.Context) (*model.ConnectConfigResults, error) {
	var res *model.ConnectConfigResults
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// ConfigurationsListUsers returns users from the configured Connect service.
//
// https://developers.docusign.com/docs/esign-rest-api/reference/connect/connectconfigurations/listusers
//
// SDK Method Connect::connectUsers
func (s *Service) ConfigurationsListUsers(connectID string) *ConfigurationsListUsersOp {
	return &ConfigurationsListUsersOp{
		Credential: s.credential,
		Method:     "GET",
		Path:       strings.Join([]string{"connect", connectID, "users"}, "/"),
		Accept:     "application/json",
		QueryOpts:  make(url.Values),
		Version:    esign.APIv21,
	}
}

// ConfigurationsListUsersOp implements DocuSign API SDK Connect::connectUsers
type ConfigurationsListUsersOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *ConfigurationsListUsersOp) Do(ctx context.Context) (*model.IntegratedUserInfoList, error) {
	var res *model.IntegratedUserInfoList
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// Count is the maximum number of results to return.
//
// Use `start_position` to specify the number of results to skip.
func (op *ConfigurationsListUsersOp) Count(val int) *ConfigurationsListUsersOp {
	if op != nil {
		op.QueryOpts.Set("count", fmt.Sprintf("%d", val))
	}
	return op
}

// EmailSubstring filters returned user records by full email address or a substring of email address.
func (op *ConfigurationsListUsersOp) EmailSubstring(val string) *ConfigurationsListUsersOp {
	if op != nil {
		op.QueryOpts.Set("email_substring", val)
	}
	return op
}

// ListIncludedUsers set the call query parameter list_included_users
func (op *ConfigurationsListUsersOp) ListIncludedUsers() *ConfigurationsListUsersOp {
	if op != nil {
		op.QueryOpts.Set("list_included_users", "true")
	}
	return op
}

// StartPosition is the zero-based index of the
// result from which to start returning results.
//
// Use with `count` to limit the number
// of results.
//
// The default value is `0`.
func (op *ConfigurationsListUsersOp) StartPosition(val int) *ConfigurationsListUsersOp {
	if op != nil {
		op.QueryOpts.Set("start_position", fmt.Sprintf("%d", val))
	}
	return op
}

// Status filters the results by user status.
// You can specify a comma-separated
// list of the following statuses:
//
// * ActivationRequired
// * ActivationSent
// * Active
// * Closed
// * Disabled
func (op *ConfigurationsListUsersOp) Status(val ...string) *ConfigurationsListUsersOp {
	if op != nil {
		op.QueryOpts.Set("status", strings.Join(val, ","))
	}
	return op
}

// UserNameSubstring filters results based on a full or partial user name.
//
// **Note:** When you enter a partial user name, you do not use a wildcard character.
func (op *ConfigurationsListUsersOp) UserNameSubstring(val string) *ConfigurationsListUsersOp {
	if op != nil {
		op.QueryOpts.Set("user_name_substring", val)
	}
	return op
}

// ConfigurationsUpdate updates a specified Connect configuration.
//
// https://developers.docusign.com/docs/esign-rest-api/reference/connect/connectconfigurations/update
//
// SDK Method Connect::updateConfiguration
func (s *Service) ConfigurationsUpdate(connectCustomConfiguration *model.ConnectCustomConfiguration) *ConfigurationsUpdateOp {
	return &ConfigurationsUpdateOp{
		Credential: s.credential,
		Method:     "PUT",
		Path:       "connect",
		Payload:    connectCustomConfiguration,
		Accept:     "application/json",
		QueryOpts:  make(url.Values),
		Version:    esign.APIv21,
	}
}

// ConfigurationsUpdateOp implements DocuSign API SDK Connect::updateConfiguration
type ConfigurationsUpdateOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *ConfigurationsUpdateOp) Do(ctx context.Context) (*model.ConnectCustomConfiguration, error) {
	var res *model.ConnectCustomConfiguration
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// EventsDelete deletes a specified Connect log entry.
//
// https://developers.docusign.com/docs/esign-rest-api/reference/connect/connectevents/delete
//
// SDK Method Connect::deleteEventLog
func (s *Service) EventsDelete(logID string) *EventsDeleteOp {
	return &EventsDeleteOp{
		Credential: s.credential,
		Method:     "DELETE",
		Path:       strings.Join([]string{"connect", "logs", logID}, "/"),
		QueryOpts:  make(url.Values),
		Version:    esign.APIv21,
	}
}

// EventsDeleteOp implements DocuSign API SDK Connect::deleteEventLog
type EventsDeleteOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *EventsDeleteOp) Do(ctx context.Context) error {
	return ((*esign.Op)(op)).Do(ctx, nil)
}

// EventsDeleteFailure deletes a Connect failure log entry.
//
// https://developers.docusign.com/docs/esign-rest-api/reference/connect/connectevents/deletefailure
//
// SDK Method Connect::deleteEventFailureLog
func (s *Service) EventsDeleteFailure(failureID string) *EventsDeleteFailureOp {
	return &EventsDeleteFailureOp{
		Credential: s.credential,
		Method:     "DELETE",
		Path:       strings.Join([]string{"connect", "failures", failureID}, "/"),
		Accept:     "application/json",
		QueryOpts:  make(url.Values),
		Version:    esign.APIv21,
	}
}

// EventsDeleteFailureOp implements DocuSign API SDK Connect::deleteEventFailureLog
type EventsDeleteFailureOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *EventsDeleteFailureOp) Do(ctx context.Context) (*model.ConnectDeleteFailureResult, error) {
	var res *model.ConnectDeleteFailureResult
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// EventsDeleteList deletes a list of Connect log entries.
//
// https://developers.docusign.com/docs/esign-rest-api/reference/connect/connectevents/deletelist
//
// SDK Method Connect::deleteEventLogs
func (s *Service) EventsDeleteList() *EventsDeleteListOp {
	return &EventsDeleteListOp{
		Credential: s.credential,
		Method:     "DELETE",
		Path:       "connect/logs",
		QueryOpts:  make(url.Values),
		Version:    esign.APIv21,
	}
}

// EventsDeleteListOp implements DocuSign API SDK Connect::deleteEventLogs
type EventsDeleteListOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *EventsDeleteListOp) Do(ctx context.Context) error {
	return ((*esign.Op)(op)).Do(ctx, nil)
}

// EventsGet gets a Connect log entry.
//
// https://developers.docusign.com/docs/esign-rest-api/reference/connect/connectevents/get
//
// SDK Method Connect::getEventLog
func (s *Service) EventsGet(logID string) *EventsGetOp {
	return &EventsGetOp{
		Credential: s.credential,
		Method:     "GET",
		Path:       strings.Join([]string{"connect", "logs", logID}, "/"),
		Accept:     "application/json",
		QueryOpts:  make(url.Values),
		Version:    esign.APIv21,
	}
}

// EventsGetOp implements DocuSign API SDK Connect::getEventLog
type EventsGetOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *EventsGetOp) Do(ctx context.Context) (*model.ConnectLog, error) {
	var res *model.ConnectLog
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// AdditionalInfo when **true,** the response includes the `connectDebugLog` information.
func (op *EventsGetOp) AdditionalInfo() *EventsGetOp {
	if op != nil {
		op.QueryOpts.Set("additional_info", "true")
	}
	return op
}

// EventsList gets the Connect log.
//
// https://developers.docusign.com/docs/esign-rest-api/reference/connect/connectevents/list
//
// SDK Method Connect::listEventLogs
func (s *Service) EventsList() *EventsListOp {
	return &EventsListOp{
		Credential: s.credential,
		Method:     "GET",
		Path:       "connect/logs",
		Accept:     "application/json",
		QueryOpts:  make(url.Values),
		Version:    esign.APIv21,
	}
}

// EventsListOp implements DocuSign API SDK Connect::listEventLogs
type EventsListOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *EventsListOp) Do(ctx context.Context) (*model.ConnectLogs, error) {
	var res *model.ConnectLogs
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// FromDate is the start date for a date range in UTC DateTime format.
//
// **Note:** If this property is null, no date filtering is applied.
func (op *EventsListOp) FromDate(val time.Time) *EventsListOp {
	if op != nil {
		op.QueryOpts.Set("from_date", val.Format(time.RFC3339))
	}
	return op
}

// ToDate is the end of a search date range in UTC DateTime format. When you use this parameter, only templates created up to this date and time are returned.
//
// **Note:** If this property is null, the value defaults to the current date.
func (op *EventsListOp) ToDate(val time.Time) *EventsListOp {
	if op != nil {
		op.QueryOpts.Set("to_date", val.Format(time.RFC3339))
	}
	return op
}

// EventsListFailures gets the Connect failure log information.
//
// https://developers.docusign.com/docs/esign-rest-api/reference/connect/connectevents/listfailures
//
// SDK Method Connect::listEventFailureLogs
func (s *Service) EventsListFailures() *EventsListFailuresOp {
	return &EventsListFailuresOp{
		Credential: s.credential,
		Method:     "GET",
		Path:       "connect/failures",
		Accept:     "application/json",
		QueryOpts:  make(url.Values),
		Version:    esign.APIv21,
	}
}

// EventsListFailuresOp implements DocuSign API SDK Connect::listEventFailureLogs
type EventsListFailuresOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *EventsListFailuresOp) Do(ctx context.Context) (*model.ConnectLogs, error) {
	var res *model.ConnectLogs
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// FromDate is the start date for a date range in UTC DateTime format.
//
// **Note:** If this property is null, no date filtering is applied.
func (op *EventsListFailuresOp) FromDate(val time.Time) *EventsListFailuresOp {
	if op != nil {
		op.QueryOpts.Set("from_date", val.Format(time.RFC3339))
	}
	return op
}

// ToDate is the end of a search date range in UTC DateTime format. When you use this parameter, only templates created up to this date and time are returned.
//
// **Note:** If this property is null, the value defaults to the current date.
func (op *EventsListFailuresOp) ToDate(val time.Time) *EventsListFailuresOp {
	if op != nil {
		op.QueryOpts.Set("to_date", val.Format(time.RFC3339))
	}
	return op
}

// EventsRetryForEnvelope republishes Connect information for the specified envelope.
// If media is an io.ReadCloser, Do() will close media.
//
// https://developers.docusign.com/docs/esign-rest-api/reference/connect/connectevents/retryforenvelope
//
// SDK Method Connect::retryEventForEnvelope
func (s *Service) EventsRetryForEnvelope(envelopeID string, media io.Reader, mimeType string) *EventsRetryForEnvelopeOp {
	return &EventsRetryForEnvelopeOp{
		Credential: s.credential,
		Method:     "PUT",
		Path:       strings.Join([]string{"connect", "envelopes", envelopeID, "retry_queue"}, "/"),
		Payload:    &esign.UploadFile{Reader: media, ContentType: mimeType},
		Accept:     "application/json",
		QueryOpts:  make(url.Values),
		Version:    esign.APIv21,
	}
}

// EventsRetryForEnvelopeOp implements DocuSign API SDK Connect::retryEventForEnvelope
type EventsRetryForEnvelopeOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *EventsRetryForEnvelopeOp) Do(ctx context.Context) (*model.ConnectFailureResults, error) {
	var res *model.ConnectFailureResults
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// EventsRetryForEnvelopes republishes Connect information for multiple envelopes.
//
// https://developers.docusign.com/docs/esign-rest-api/reference/connect/connectevents/retryforenvelopes
//
// SDK Method Connect::retryEventForEnvelopes
func (s *Service) EventsRetryForEnvelopes(connectFailureFilter *model.ConnectFailureFilter) *EventsRetryForEnvelopesOp {
	return &EventsRetryForEnvelopesOp{
		Credential: s.credential,
		Method:     "PUT",
		Path:       "connect/envelopes/retry_queue",
		Payload:    connectFailureFilter,
		Accept:     "application/json",
		QueryOpts:  make(url.Values),
		Version:    esign.APIv21,
	}
}

// EventsRetryForEnvelopesOp implements DocuSign API SDK Connect::retryEventForEnvelopes
type EventsRetryForEnvelopesOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *EventsRetryForEnvelopesOp) Do(ctx context.Context) (*model.ConnectFailureResults, error) {
	var res *model.ConnectFailureResults
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// ConfigurationsCreateConnectOAuthConfig sets the Connect OAuth Config for the account.
//
// https://developers.docusign.com/docs/esign-rest-api/reference/connect/connectconfigurations/createconnectoauthconfig
//
// SDK Method Connect::createConnectOAuthConfig
func (s *Service) ConfigurationsCreateConnectOAuthConfig(connectOAuthConfig *model.ConnectOAuthConfig) *ConfigurationsCreateConnectOAuthConfigOp {
	return &ConfigurationsCreateConnectOAuthConfigOp{
		Credential: s.credential,
		Method:     "POST",
		Path:       "connect/oauth",
		Payload:    connectOAuthConfig,
		QueryOpts:  make(url.Values),
		Version:    esign.APIv21,
	}
}

// ConfigurationsCreateConnectOAuthConfigOp implements DocuSign API SDK Connect::createConnectOAuthConfig
type ConfigurationsCreateConnectOAuthConfigOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *ConfigurationsCreateConnectOAuthConfigOp) Do(ctx context.Context) (*model.ConnectOAuthConfig, error) {
	var res *model.ConnectOAuthConfig
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// ConfigurationsDeleteConnectOAuthConfig sets the Connect OAuth Config for the account.
//
// https://developers.docusign.com/docs/esign-rest-api/reference/connect/connectconfigurations/deleteconnectoauthconfig
//
// SDK Method Connect::deleteConnectOAuthConfig
func (s *Service) ConfigurationsDeleteConnectOAuthConfig() *ConfigurationsDeleteConnectOAuthConfigOp {
	return &ConfigurationsDeleteConnectOAuthConfigOp{
		Credential: s.credential,
		Method:     "DELETE",
		Path:       "connect/oauth",
		QueryOpts:  make(url.Values),
		Version:    esign.APIv21,
	}
}

// ConfigurationsDeleteConnectOAuthConfigOp implements DocuSign API SDK Connect::deleteConnectOAuthConfig
type ConfigurationsDeleteConnectOAuthConfigOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *ConfigurationsDeleteConnectOAuthConfigOp) Do(ctx context.Context) error {
	return ((*esign.Op)(op)).Do(ctx, nil)
}

// ConfigurationsGetConnectAllUsers returns all users from the configured Connect service.
//
// https://developers.docusign.com/docs/esign-rest-api/reference/connect/connectconfigurations/getconnectallusers
//
// SDK Method Connect::getConnectAllUsers
func (s *Service) ConfigurationsGetConnectAllUsers(connectID string) *ConfigurationsGetConnectAllUsersOp {
	return &ConfigurationsGetConnectAllUsersOp{
		Credential: s.credential,
		Method:     "GET",
		Path:       strings.Join([]string{"connect", connectID, "all", "users"}, "/"),
		Accept:     "application/json",
		QueryOpts:  make(url.Values),
		Version:    esign.APIv21,
	}
}

// ConfigurationsGetConnectAllUsersOp implements DocuSign API SDK Connect::getConnectAllUsers
type ConfigurationsGetConnectAllUsersOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *ConfigurationsGetConnectAllUsersOp) Do(ctx context.Context) (*model.IntegratedConnectUserInfoList, error) {
	var res *model.IntegratedConnectUserInfoList
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// Count is the maximum number of results to return.
func (op *ConfigurationsGetConnectAllUsersOp) Count(val string) *ConfigurationsGetConnectAllUsersOp {
	if op != nil {
		op.QueryOpts.Set("count", val)
	}
	return op
}

// DomainUsersOnly set the call query parameter domain_users_only
func (op *ConfigurationsGetConnectAllUsersOp) DomainUsersOnly(val string) *ConfigurationsGetConnectAllUsersOp {
	if op != nil {
		op.QueryOpts.Set("domain_users_only", val)
	}
	return op
}

// EmailSubstring filters returned user records by full email address or a substring of email address.
func (op *ConfigurationsGetConnectAllUsersOp) EmailSubstring(val string) *ConfigurationsGetConnectAllUsersOp {
	if op != nil {
		op.QueryOpts.Set("email_substring", val)
	}
	return op
}

// StartPosition is the position within the total result set from which to start returning values. The value **thumbnail** may be used to return the page image.
func (op *ConfigurationsGetConnectAllUsersOp) StartPosition(val string) *ConfigurationsGetConnectAllUsersOp {
	if op != nil {
		op.QueryOpts.Set("start_position", val)
	}
	return op
}

// Status is the status of the item.
func (op *ConfigurationsGetConnectAllUsersOp) Status(val string) *ConfigurationsGetConnectAllUsersOp {
	if op != nil {
		op.QueryOpts.Set("status", val)
	}
	return op
}

// UserNameSubstring filters results based on a full or partial user name.
//
// **Note:** When you enter a partial user name, you do not use a wildcard character.
func (op *ConfigurationsGetConnectAllUsersOp) UserNameSubstring(val string) *ConfigurationsGetConnectAllUsersOp {
	if op != nil {
		op.QueryOpts.Set("user_name_substring", val)
	}
	return op
}

// ConfigurationsGetConnectOAuthConfig sets the Connect OAuth Config for the account.
//
// https://developers.docusign.com/docs/esign-rest-api/reference/connect/connectconfigurations/getconnectoauthconfig
//
// SDK Method Connect::getConnectOAuthConfig
func (s *Service) ConfigurationsGetConnectOAuthConfig() *ConfigurationsGetConnectOAuthConfigOp {
	return &ConfigurationsGetConnectOAuthConfigOp{
		Credential: s.credential,
		Method:     "GET",
		Path:       "connect/oauth",
		Accept:     "application/json",
		QueryOpts:  make(url.Values),
		Version:    esign.APIv21,
	}
}

// ConfigurationsGetConnectOAuthConfigOp implements DocuSign API SDK Connect::getConnectOAuthConfig
type ConfigurationsGetConnectOAuthConfigOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *ConfigurationsGetConnectOAuthConfigOp) Do(ctx context.Context) (*model.ConnectOAuthConfig, error) {
	var res *model.ConnectOAuthConfig
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}
