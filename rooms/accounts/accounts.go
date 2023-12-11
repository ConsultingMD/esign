// Copyright 2022 James Cote
// All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by gen-esign; DO NOT EDIT.

// Package accounts implements the DocuSign SDK
// category Accounts.
//
// Information about accounts.
//
// Service Api documentation may be found at:
// https://developers.docusign.com/docs/rooms-api/reference/Accounts
// Usage example:
//
//	import (
//	    "github.com/ConsultingMD/esign"
//	    "github.com/ConsultingMD/esign/rooms"
//	)
//	...
//	accountsService := accounts.New(esignCredential)
package accounts // import "github.com/ConsultingMD/esignrooms//accounts"

import (
	"context"
	"net/url"

	"github.com/ConsultingMD/esign"
	"github.com/ConsultingMD/esign/rooms"
)

// Service implements DocuSign Accounts API operations
type Service struct {
	credential esign.Credential
}

// New initializes a accounts service using cred to authorize ops.
func New(cred esign.Credential) *Service {
	return &Service{credential: cred}
}

// GetAccountInformation gets account information.
//
// https://developers.docusign.com/docs/rooms-api/reference/accounts/accounts/getaccountinformation
//
// SDK Method Accounts::GetAccountInformation
func (s *Service) GetAccountInformation() *GetAccountInformationOp {
	return &GetAccountInformationOp{
		Credential: s.credential,
		Method:     "GET",
		Path:       "/v2/accounts/{accountId}",
		Accept:     "application/json",
		QueryOpts:  make(url.Values),
		Version:    esign.RoomsV2,
	}
}

// GetAccountInformationOp implements DocuSign API SDK Accounts::GetAccountInformation
type GetAccountInformationOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *GetAccountInformationOp) Do(ctx context.Context) (*rooms.AccountSummary, error) {
	var res *rooms.AccountSummary
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}
