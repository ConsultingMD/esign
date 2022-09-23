// Copyright 2019 James Cote
// All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by gen-esign; DO NOT EDIT.

// Package users implements the DocuSign SDK
// category Users.
//
// Use the Users category to manage the users in your accounts.
//
// You can:
//
// * Set custom user settings.
// * Manage a users profile.
// * Add delete users.
// * Add and delete the intials and signature images for a user.
//
// Service Api documentation may be found at:
// https://developers.docusign.com/docs/esign-rest-api/v2/reference/Users
// Usage example:
//
//   import (
//       "github.com/jfcote87/esign"
//       "github.com/jfcote87/esign/v2/users"
//       "github.com/jfcote87/esign/v2/model"
//   )
//   ...
//   usersService := users.New(esignCredential)
package users // import "github.com/jfcote87/esign/v2//users"

import (
	"context"
	"fmt"
	"io"
	"net/url"
	"strings"

	"github.com/jfcote87/esign"
	"github.com/jfcote87/esign/v2/model"
)

// Service implements DocuSign Users Category API operations
type Service struct {
	credential esign.Credential
}

// New initializes a users service using cred to authorize ops.
func New(cred esign.Credential) *Service {
	return &Service{credential: cred}
}

// ContactsCreate imports multiple new contacts into the contacts collection from CSV, JSON, or XML (based on content type).
//
// https://developers.docusign.com/docs/esign-rest-api/v2/reference/users/contacts/create
//
// SDK Method Users::postContacts
func (s *Service) ContactsCreate(contactModRequest *model.ContactModRequest) *ContactsCreateOp {
	return &ContactsCreateOp{
		Credential: s.credential,
		Method:     "POST",
		Path:       "contacts",
		Payload:    contactModRequest,
		QueryOpts:  make(url.Values),
		Version:    esign.APIv2,
	}
}

// ContactsCreateOp implements DocuSign API SDK Users::postContacts
type ContactsCreateOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *ContactsCreateOp) Do(ctx context.Context) (*model.ContactUpdateResponse, error) {
	var res *model.ContactUpdateResponse
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// ContactsDelete replaces a particular contact associated with an account for the DocuSign service.
//
// https://developers.docusign.com/docs/esign-rest-api/v2/reference/users/contacts/delete
//
// SDK Method Users::deleteContactWithId
func (s *Service) ContactsDelete(contactID string) *ContactsDeleteOp {
	return &ContactsDeleteOp{
		Credential: s.credential,
		Method:     "DELETE",
		Path:       strings.Join([]string{"contacts", contactID}, "/"),
		Accept:     "application/json",
		QueryOpts:  make(url.Values),
		Version:    esign.APIv2,
	}
}

// ContactsDeleteOp implements DocuSign API SDK Users::deleteContactWithId
type ContactsDeleteOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *ContactsDeleteOp) Do(ctx context.Context) (*model.ContactUpdateResponse, error) {
	var res *model.ContactUpdateResponse
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// ContactsDeleteList delete contacts associated with an account for the DocuSign service.
//
// https://developers.docusign.com/docs/esign-rest-api/v2/reference/users/contacts/deletelist
//
// SDK Method Users::deleteContacts
func (s *Service) ContactsDeleteList(contactModRequest *model.ContactModRequest) *ContactsDeleteListOp {
	return &ContactsDeleteListOp{
		Credential: s.credential,
		Method:     "DELETE",
		Path:       "contacts",
		Payload:    contactModRequest,
		Accept:     "application/json",
		QueryOpts:  make(url.Values),
		Version:    esign.APIv2,
	}
}

// ContactsDeleteListOp implements DocuSign API SDK Users::deleteContacts
type ContactsDeleteListOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *ContactsDeleteListOp) Do(ctx context.Context) (*model.ContactUpdateResponse, error) {
	var res *model.ContactUpdateResponse
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// ContactsGet gets a particular contact associated with the user's account.
//
// https://developers.docusign.com/docs/esign-rest-api/v2/reference/users/contacts/get
//
// SDK Method Users::getContactById
func (s *Service) ContactsGet(contactID string) *ContactsGetOp {
	return &ContactsGetOp{
		Credential: s.credential,
		Method:     "GET",
		Path:       strings.Join([]string{"contacts", contactID}, "/"),
		Accept:     "application/json",
		QueryOpts:  make(url.Values),
		Version:    esign.APIv2,
	}
}

// ContactsGetOp implements DocuSign API SDK Users::getContactById
type ContactsGetOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *ContactsGetOp) Do(ctx context.Context) (*model.ContactGetResponse, error) {
	var res *model.ContactGetResponse
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// ContactsUpdate replaces contacts associated with an account for the DocuSign service.
//
// https://developers.docusign.com/docs/esign-rest-api/v2/reference/users/contacts/update
//
// SDK Method Users::putContacts
func (s *Service) ContactsUpdate(contactModRequest *model.ContactModRequest) *ContactsUpdateOp {
	return &ContactsUpdateOp{
		Credential: s.credential,
		Method:     "PUT",
		Path:       "contacts",
		Payload:    contactModRequest,
		Accept:     "application/json",
		QueryOpts:  make(url.Values),
		Version:    esign.APIv2,
	}
}

// ContactsUpdateOp implements DocuSign API SDK Users::putContacts
type ContactsUpdateOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *ContactsUpdateOp) Do(ctx context.Context) (*model.ContactUpdateResponse, error) {
	var res *model.ContactUpdateResponse
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// CustomSettingsDelete deletes custom user settings for a specified user.
//
// https://developers.docusign.com/docs/esign-rest-api/v2/reference/users/usercustomsettings/delete
//
// SDK Method Users::deleteCustomSettings
func (s *Service) CustomSettingsDelete(userID string, customSettingsInformation *model.CustomSettingsInformation) *CustomSettingsDeleteOp {
	return &CustomSettingsDeleteOp{
		Credential: s.credential,
		Method:     "DELETE",
		Path:       strings.Join([]string{"users", userID, "custom_settings"}, "/"),
		Payload:    customSettingsInformation,
		Accept:     "application/json",
		QueryOpts:  make(url.Values),
		Version:    esign.APIv2,
	}
}

// CustomSettingsDeleteOp implements DocuSign API SDK Users::deleteCustomSettings
type CustomSettingsDeleteOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *CustomSettingsDeleteOp) Do(ctx context.Context) (*model.CustomSettingsInformation, error) {
	var res *model.CustomSettingsInformation
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// CustomSettingsList retrieves the custom user settings for a specified user.
//
// https://developers.docusign.com/docs/esign-rest-api/v2/reference/users/usercustomsettings/list
//
// SDK Method Users::listCustomSettings
func (s *Service) CustomSettingsList(userID string) *CustomSettingsListOp {
	return &CustomSettingsListOp{
		Credential: s.credential,
		Method:     "GET",
		Path:       strings.Join([]string{"users", userID, "custom_settings"}, "/"),
		Accept:     "application/json",
		QueryOpts:  make(url.Values),
		Version:    esign.APIv2,
	}
}

// CustomSettingsListOp implements DocuSign API SDK Users::listCustomSettings
type CustomSettingsListOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *CustomSettingsListOp) Do(ctx context.Context) (*model.CustomSettingsInformation, error) {
	var res *model.CustomSettingsInformation
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// CustomSettingsUpdate adds or updates custom user settings for the specified user.
//
// https://developers.docusign.com/docs/esign-rest-api/v2/reference/users/usercustomsettings/update
//
// SDK Method Users::updateCustomSettings
func (s *Service) CustomSettingsUpdate(userID string, customSettingsInformation *model.CustomSettingsInformation) *CustomSettingsUpdateOp {
	return &CustomSettingsUpdateOp{
		Credential: s.credential,
		Method:     "PUT",
		Path:       strings.Join([]string{"users", userID, "custom_settings"}, "/"),
		Payload:    customSettingsInformation,
		Accept:     "application/json",
		QueryOpts:  make(url.Values),
		Version:    esign.APIv2,
	}
}

// CustomSettingsUpdateOp implements DocuSign API SDK Users::updateCustomSettings
type CustomSettingsUpdateOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *CustomSettingsUpdateOp) Do(ctx context.Context) (*model.CustomSettingsInformation, error) {
	var res *model.CustomSettingsInformation
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// ProfilesGet retrieves the user profile for a specified user.
//
// https://developers.docusign.com/docs/esign-rest-api/v2/reference/users/userprofiles/get
//
// SDK Method Users::getProfile
func (s *Service) ProfilesGet(userID string) *ProfilesGetOp {
	return &ProfilesGetOp{
		Credential: s.credential,
		Method:     "GET",
		Path:       strings.Join([]string{"users", userID, "profile"}, "/"),
		Accept:     "application/json",
		QueryOpts:  make(url.Values),
		Version:    esign.APIv2,
	}
}

// ProfilesGetOp implements DocuSign API SDK Users::getProfile
type ProfilesGetOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *ProfilesGetOp) Do(ctx context.Context) (*model.UserProfile, error) {
	var res *model.UserProfile
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// ProfilesUpdate updates the user profile information for the specified user.
//
// https://developers.docusign.com/docs/esign-rest-api/v2/reference/users/userprofiles/update
//
// SDK Method Users::updateProfile
func (s *Service) ProfilesUpdate(userID string, userProfile *model.UserProfile) *ProfilesUpdateOp {
	return &ProfilesUpdateOp{
		Credential: s.credential,
		Method:     "PUT",
		Path:       strings.Join([]string{"users", userID, "profile"}, "/"),
		Payload:    userProfile,
		QueryOpts:  make(url.Values),
		Version:    esign.APIv2,
	}
}

// ProfilesUpdateOp implements DocuSign API SDK Users::updateProfile
type ProfilesUpdateOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *ProfilesUpdateOp) Do(ctx context.Context) error {
	return ((*esign.Op)(op)).Do(ctx, nil)
}

// SignaturesCreate adds user Signature and initials images to a Signature.
// If any uploads[x].Reader is an io.ReadCloser(s), Do() will always close Reader.
//
// https://developers.docusign.com/docs/esign-rest-api/v2/reference/users/usersignatures/create
//
// SDK Method Users::createSignatures
func (s *Service) SignaturesCreate(userID string, userSignaturesInformation *model.UserSignaturesInformation, uploads ...*esign.UploadFile) *SignaturesCreateOp {
	return &SignaturesCreateOp{
		Credential: s.credential,
		Method:     "POST",
		Path:       strings.Join([]string{"users", userID, "signatures"}, "/"),
		Payload:    userSignaturesInformation,
		Files:      uploads,
		QueryOpts:  make(url.Values),
		Version:    esign.APIv2,
	}
}

// SignaturesCreateOp implements DocuSign API SDK Users::createSignatures
type SignaturesCreateOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *SignaturesCreateOp) Do(ctx context.Context) (*model.UserSignaturesInformation, error) {
	var res *model.UserSignaturesInformation
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// SignaturesDelete removes removes signature information for the specified user.
//
// https://developers.docusign.com/docs/esign-rest-api/v2/reference/users/usersignatures/delete
//
// SDK Method Users::deleteSignature
func (s *Service) SignaturesDelete(signatureID string, userID string) *SignaturesDeleteOp {
	return &SignaturesDeleteOp{
		Credential: s.credential,
		Method:     "DELETE",
		Path:       strings.Join([]string{"users", userID, "signatures", signatureID}, "/"),
		QueryOpts:  make(url.Values),
		Version:    esign.APIv2,
	}
}

// SignaturesDeleteOp implements DocuSign API SDK Users::deleteSignature
type SignaturesDeleteOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *SignaturesDeleteOp) Do(ctx context.Context) error {
	return ((*esign.Op)(op)).Do(ctx, nil)
}

// SignaturesDeleteImage deletes the user initials image or the  user signature image for the specified user.
//
// https://developers.docusign.com/docs/esign-rest-api/v2/reference/users/usersignatures/deleteimage
//
// SDK Method Users::deleteSignatureImage
func (s *Service) SignaturesDeleteImage(imageType string, signatureID string, userID string) *SignaturesDeleteImageOp {
	return &SignaturesDeleteImageOp{
		Credential: s.credential,
		Method:     "DELETE",
		Path:       strings.Join([]string{"users", userID, "signatures", signatureID, imageType}, "/"),
		Accept:     "application/json",
		QueryOpts:  make(url.Values),
		Version:    esign.APIv2,
	}
}

// SignaturesDeleteImageOp implements DocuSign API SDK Users::deleteSignatureImage
type SignaturesDeleteImageOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *SignaturesDeleteImageOp) Do(ctx context.Context) (*model.UserSignature, error) {
	var res *model.UserSignature
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// SignaturesGet gets the user signature information for the specified user.
//
// https://developers.docusign.com/docs/esign-rest-api/v2/reference/users/usersignatures/get
//
// SDK Method Users::getSignature
func (s *Service) SignaturesGet(signatureID string, userID string) *SignaturesGetOp {
	return &SignaturesGetOp{
		Credential: s.credential,
		Method:     "GET",
		Path:       strings.Join([]string{"users", userID, "signatures", signatureID}, "/"),
		Accept:     "application/json",
		QueryOpts:  make(url.Values),
		Version:    esign.APIv2,
	}
}

// SignaturesGetOp implements DocuSign API SDK Users::getSignature
type SignaturesGetOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *SignaturesGetOp) Do(ctx context.Context) (*model.UserSignature, error) {
	var res *model.UserSignature
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// SignaturesGetImage retrieves the user initials image or the  user signature image for the specified user.
//
// https://developers.docusign.com/docs/esign-rest-api/v2/reference/users/usersignatures/getimage
//
// SDK Method Users::getSignatureImage
func (s *Service) SignaturesGetImage(imageType string, signatureID string, userID string) *SignaturesGetImageOp {
	return &SignaturesGetImageOp{
		Credential: s.credential,
		Method:     "GET",
		Path:       strings.Join([]string{"users", userID, "signatures", signatureID, imageType}, "/"),
		QueryOpts:  make(url.Values),
		Version:    esign.APIv2,
	}
}

// SignaturesGetImageOp implements DocuSign API SDK Users::getSignatureImage
type SignaturesGetImageOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *SignaturesGetImageOp) Do(ctx context.Context) (*esign.Download, error) {
	var res *esign.Download
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// IncludeChrome set the call query parameter include_chrome
func (op *SignaturesGetImageOp) IncludeChrome() *SignaturesGetImageOp {
	if op != nil {
		op.QueryOpts.Set("include_chrome", "true")
	}
	return op
}

// SignaturesList retrieves a list of user signature definitions for a specified user.
//
// https://developers.docusign.com/docs/esign-rest-api/v2/reference/users/usersignatures/list
//
// SDK Method Users::listSignatures
func (s *Service) SignaturesList(userID string) *SignaturesListOp {
	return &SignaturesListOp{
		Credential: s.credential,
		Method:     "GET",
		Path:       strings.Join([]string{"users", userID, "signatures"}, "/"),
		Accept:     "application/json",
		QueryOpts:  make(url.Values),
		Version:    esign.APIv2,
	}
}

// SignaturesListOp implements DocuSign API SDK Users::listSignatures
type SignaturesListOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *SignaturesListOp) Do(ctx context.Context) (*model.UserSignaturesInformation, error) {
	var res *model.UserSignaturesInformation
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// StampType set the call query parameter stamp_type
func (op *SignaturesListOp) StampType(val string) *SignaturesListOp {
	if op != nil {
		op.QueryOpts.Set("stamp_type", val)
	}
	return op
}

// SignaturesUpdate updates the user signature for a specified user.
//
// https://developers.docusign.com/docs/esign-rest-api/v2/reference/users/usersignatures/update
//
// SDK Method Users::updateSignature
func (s *Service) SignaturesUpdate(signatureID string, userID string, userSignatureDefinition *model.UserSignatureDefinition) *SignaturesUpdateOp {
	return &SignaturesUpdateOp{
		Credential: s.credential,
		Method:     "PUT",
		Path:       strings.Join([]string{"users", userID, "signatures", signatureID}, "/"),
		Payload:    userSignatureDefinition,
		Accept:     "application/json",
		QueryOpts:  make(url.Values),
		Version:    esign.APIv2,
	}
}

// SignaturesUpdateOp implements DocuSign API SDK Users::updateSignature
type SignaturesUpdateOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *SignaturesUpdateOp) Do(ctx context.Context) (*model.UserSignature, error) {
	var res *model.UserSignature
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// CloseExistingSignature when set to **true**, closes the current signature.
func (op *SignaturesUpdateOp) CloseExistingSignature() *SignaturesUpdateOp {
	if op != nil {
		op.QueryOpts.Set("close_existing_signature", "true")
	}
	return op
}

// SignaturesUpdateImage updates the user signature image or user initials image for the specified user.
// If media is an io.ReadCloser, Do() will close media.
//
// https://developers.docusign.com/docs/esign-rest-api/v2/reference/users/usersignatures/updateimage
//
// SDK Method Users::updateSignatureImage
func (s *Service) SignaturesUpdateImage(imageType string, signatureID string, userID string, media io.Reader, mimeType string) *SignaturesUpdateImageOp {
	return &SignaturesUpdateImageOp{
		Credential: s.credential,
		Method:     "PUT",
		Path:       strings.Join([]string{"users", userID, "signatures", signatureID, imageType}, "/"),
		Payload:    &esign.UploadFile{Reader: media, ContentType: mimeType},
		Accept:     "application/json",
		QueryOpts:  make(url.Values),
		Version:    esign.APIv2,
	}
}

// SignaturesUpdateImageOp implements DocuSign API SDK Users::updateSignatureImage
type SignaturesUpdateImageOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *SignaturesUpdateImageOp) Do(ctx context.Context) (*model.UserSignature, error) {
	var res *model.UserSignature
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// SignaturesUpdateList adds/updates a user signature.
//
// https://developers.docusign.com/docs/esign-rest-api/v2/reference/users/usersignatures/updatelist
//
// SDK Method Users::updateSignatures
func (s *Service) SignaturesUpdateList(userID string, userSignaturesInformation *model.UserSignaturesInformation) *SignaturesUpdateListOp {
	return &SignaturesUpdateListOp{
		Credential: s.credential,
		Method:     "PUT",
		Path:       strings.Join([]string{"users", userID, "signatures"}, "/"),
		Payload:    userSignaturesInformation,
		Accept:     "application/json",
		QueryOpts:  make(url.Values),
		Version:    esign.APIv2,
	}
}

// SignaturesUpdateListOp implements DocuSign API SDK Users::updateSignatures
type SignaturesUpdateListOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *SignaturesUpdateListOp) Do(ctx context.Context) (*model.UserSignaturesInformation, error) {
	var res *model.UserSignaturesInformation
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// Create adds news user to the specified account.
//
// https://developers.docusign.com/docs/esign-rest-api/v2/reference/users/users/create
//
// SDK Method Users::create
func (s *Service) Create(newUsersDefinition *model.NewUsersDefinition) *CreateOp {
	return &CreateOp{
		Credential: s.credential,
		Method:     "POST",
		Path:       "users",
		Payload:    newUsersDefinition,
		QueryOpts:  make(url.Values),
		Version:    esign.APIv2,
	}
}

// CreateOp implements DocuSign API SDK Users::create
type CreateOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *CreateOp) Do(ctx context.Context) (*model.NewUsersSummary, error) {
	var res *model.NewUsersSummary
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// Delete removes users account privileges.
//
// https://developers.docusign.com/docs/esign-rest-api/v2/reference/users/users/delete
//
// SDK Method Users::delete
func (s *Service) Delete(userInfoList *model.UserInfoList) *DeleteOp {
	return &DeleteOp{
		Credential: s.credential,
		Method:     "DELETE",
		Path:       "users",
		Payload:    userInfoList,
		Accept:     "application/json",
		QueryOpts:  make(url.Values),
		Version:    esign.APIv2,
	}
}

// DeleteOp implements DocuSign API SDK Users::delete
type DeleteOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *DeleteOp) Do(ctx context.Context) (*model.UsersResponse, error) {
	var res *model.UsersResponse
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// Delete set the call query parameter delete
func (op *DeleteOp) Delete(val string) *DeleteOp {
	if op != nil {
		op.QueryOpts.Set("delete", val)
	}
	return op
}

// DeleteProfileImage deletes the user profile image for the specified user.
//
// https://developers.docusign.com/docs/esign-rest-api/v2/reference/users/users/deleteprofileimage
//
// SDK Method Users::deleteProfileImage
func (s *Service) DeleteProfileImage(userID string) *DeleteProfileImageOp {
	return &DeleteProfileImageOp{
		Credential: s.credential,
		Method:     "DELETE",
		Path:       strings.Join([]string{"users", userID, "profile", "image"}, "/"),
		QueryOpts:  make(url.Values),
		Version:    esign.APIv2,
	}
}

// DeleteProfileImageOp implements DocuSign API SDK Users::deleteProfileImage
type DeleteProfileImageOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *DeleteProfileImageOp) Do(ctx context.Context) error {
	return ((*esign.Op)(op)).Do(ctx, nil)
}

// Get gets the user information for a specified user.
//
// https://developers.docusign.com/docs/esign-rest-api/v2/reference/users/users/get
//
// SDK Method Users::getInformation
func (s *Service) Get(userID string) *GetOp {
	return &GetOp{
		Credential: s.credential,
		Method:     "GET",
		Path:       strings.Join([]string{"users", userID}, "/"),
		Accept:     "application/json",
		QueryOpts:  make(url.Values),
		Version:    esign.APIv2,
	}
}

// GetOp implements DocuSign API SDK Users::getInformation
type GetOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *GetOp) Do(ctx context.Context) (*model.UserInformation, error) {
	var res *model.UserInformation
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// AdditionalInfo when set to **true**, the full list of user information is returned for each user in the account.
func (op *GetOp) AdditionalInfo() *GetOp {
	if op != nil {
		op.QueryOpts.Set("additional_info", "true")
	}
	return op
}

// Email set the call query parameter email
func (op *GetOp) Email(val string) *GetOp {
	if op != nil {
		op.QueryOpts.Set("email", val)
	}
	return op
}

// GetProfileImage retrieves the user profile image for the specified user.
//
// https://developers.docusign.com/docs/esign-rest-api/v2/reference/users/users/getprofileimage
//
// SDK Method Users::getProfileImage
func (s *Service) GetProfileImage(userID string) *GetProfileImageOp {
	return &GetProfileImageOp{
		Credential: s.credential,
		Method:     "GET",
		Path:       strings.Join([]string{"users", userID, "profile", "image"}, "/"),
		QueryOpts:  make(url.Values),
		Version:    esign.APIv2,
	}
}

// GetProfileImageOp implements DocuSign API SDK Users::getProfileImage
type GetProfileImageOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *GetProfileImageOp) Do(ctx context.Context) (*esign.Download, error) {
	var res *esign.Download
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// Encoding set the call query parameter encoding
func (op *GetProfileImageOp) Encoding(val string) *GetProfileImageOp {
	if op != nil {
		op.QueryOpts.Set("encoding", val)
	}
	return op
}

// GetSettings gets the user account settings for a specified user.
//
// https://developers.docusign.com/docs/esign-rest-api/v2/reference/users/users/getsettings
//
// SDK Method Users::getSettings
func (s *Service) GetSettings(userID string) *GetSettingsOp {
	return &GetSettingsOp{
		Credential: s.credential,
		Method:     "GET",
		Path:       strings.Join([]string{"users", userID, "settings"}, "/"),
		Accept:     "application/json",
		QueryOpts:  make(url.Values),
		Version:    esign.APIv2,
	}
}

// GetSettingsOp implements DocuSign API SDK Users::getSettings
type GetSettingsOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *GetSettingsOp) Do(ctx context.Context) (*model.UserSettingsInformation, error) {
	var res *model.UserSettingsInformation
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// List retrieves the list of users for the specified account.
//
// https://developers.docusign.com/docs/esign-rest-api/v2/reference/users/users/list
//
// SDK Method Users::list
func (s *Service) List() *ListOp {
	return &ListOp{
		Credential: s.credential,
		Method:     "GET",
		Path:       "users",
		Accept:     "application/json",
		QueryOpts:  make(url.Values),
		Version:    esign.APIv2,
	}
}

// ListOp implements DocuSign API SDK Users::list
type ListOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *ListOp) Do(ctx context.Context) (*model.UserInformationList, error) {
	var res *model.UserInformationList
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// AdditionalInfo when set to **true**, the full list of user information is returned for each user in the account.
func (op *ListOp) AdditionalInfo() *ListOp {
	if op != nil {
		op.QueryOpts.Set("additional_info", "true")
	}
	return op
}

// Count number of records to return. The number must be greater than 0 and less than or equal to 100.
func (op *ListOp) Count(val int) *ListOp {
	if op != nil {
		op.QueryOpts.Set("count", fmt.Sprintf("%d", val))
	}
	return op
}

// Email set the call query parameter email
func (op *ListOp) Email(val string) *ListOp {
	if op != nil {
		op.QueryOpts.Set("email", val)
	}
	return op
}

// EmailSubstring filters the returned user records by the email address or a sub-string of email address.
func (op *ListOp) EmailSubstring(val string) *ListOp {
	if op != nil {
		op.QueryOpts.Set("email_substring", val)
	}
	return op
}

// GroupID filters user records returned by one or more group Id's.
func (op *ListOp) GroupID(val string) *ListOp {
	if op != nil {
		op.QueryOpts.Set("group_id", val)
	}
	return op
}

// LoginStatus set the call query parameter login_status
func (op *ListOp) LoginStatus(val string) *ListOp {
	if op != nil {
		op.QueryOpts.Set("login_status", val)
	}
	return op
}

// NotGroupID set the call query parameter not_group_id
func (op *ListOp) NotGroupID(val string) *ListOp {
	if op != nil {
		op.QueryOpts.Set("not_group_id", val)
	}
	return op
}

// StartPosition starting value for the list.
func (op *ListOp) StartPosition(val int) *ListOp {
	if op != nil {
		op.QueryOpts.Set("start_position", fmt.Sprintf("%d", val))
	}
	return op
}

// Status filters the user records by account status. One of:
//
// * ActivationRequired
// * ActivationSent
// * Active
// * Closed
// * Disabled
func (op *ListOp) Status(val ...string) *ListOp {
	if op != nil {
		op.QueryOpts.Set("status", strings.Join(val, ","))
	}
	return op
}

// UserNameSubstring filters the user records returned by the user name or a sub-string of user name.
func (op *ListOp) UserNameSubstring(val string) *ListOp {
	if op != nil {
		op.QueryOpts.Set("user_name_substring", val)
	}
	return op
}

// Update updates the specified user information.
//
// https://developers.docusign.com/docs/esign-rest-api/v2/reference/users/users/update
//
// SDK Method Users::updateUser
func (s *Service) Update(userID string, userInformation *model.UserInformation) *UpdateOp {
	return &UpdateOp{
		Credential: s.credential,
		Method:     "PUT",
		Path:       strings.Join([]string{"users", userID}, "/"),
		Payload:    userInformation,
		Accept:     "application/json",
		QueryOpts:  make(url.Values),
		Version:    esign.APIv2,
	}
}

// UpdateOp implements DocuSign API SDK Users::updateUser
type UpdateOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *UpdateOp) Do(ctx context.Context) (*model.UserInformation, error) {
	var res *model.UserInformation
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// UpdateList change one or more user in the specified account.
//
// https://developers.docusign.com/docs/esign-rest-api/v2/reference/users/users/updatelist
//
// SDK Method Users::updateUsers
func (s *Service) UpdateList(userInformationList *model.UserInformationList) *UpdateListOp {
	return &UpdateListOp{
		Credential: s.credential,
		Method:     "PUT",
		Path:       "users",
		Payload:    userInformationList,
		Accept:     "application/json",
		QueryOpts:  make(url.Values),
		Version:    esign.APIv2,
	}
}

// UpdateListOp implements DocuSign API SDK Users::updateUsers
type UpdateListOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *UpdateListOp) Do(ctx context.Context) (*model.UserInformationList, error) {
	var res *model.UserInformationList
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// UpdateProfileImage updates the user profile image for a specified user.
// If media is an io.ReadCloser, Do() will close media.
//
// https://developers.docusign.com/docs/esign-rest-api/v2/reference/users/users/updateprofileimage
//
// SDK Method Users::updateProfileImage
func (s *Service) UpdateProfileImage(userID string, media io.Reader, mimeType string) *UpdateProfileImageOp {
	return &UpdateProfileImageOp{
		Credential: s.credential,
		Method:     "PUT",
		Path:       strings.Join([]string{"users", userID, "profile", "image"}, "/"),
		Payload:    &esign.UploadFile{Reader: media, ContentType: mimeType},
		Accept:     "image/gif",
		QueryOpts:  make(url.Values),
		Version:    esign.APIv2,
	}
}

// UpdateProfileImageOp implements DocuSign API SDK Users::updateProfileImage
type UpdateProfileImageOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *UpdateProfileImageOp) Do(ctx context.Context) error {
	return ((*esign.Op)(op)).Do(ctx, nil)
}

// UpdateSettings updates the user account settings for a specified user.
//
// https://developers.docusign.com/docs/esign-rest-api/v2/reference/users/users/updatesettings
//
// SDK Method Users::updateSettings
func (s *Service) UpdateSettings(userID string, userSettingsInformation *model.UserSettingsInformation) *UpdateSettingsOp {
	return &UpdateSettingsOp{
		Credential: s.credential,
		Method:     "PUT",
		Path:       strings.Join([]string{"users", userID, "settings"}, "/"),
		Payload:    userSettingsInformation,
		QueryOpts:  make(url.Values),
		Version:    esign.APIv2,
	}
}

// UpdateSettingsOp implements DocuSign API SDK Users::updateSettings
type UpdateSettingsOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *UpdateSettingsOp) Do(ctx context.Context) error {
	return ((*esign.Op)(op)).Do(ctx, nil)
}
