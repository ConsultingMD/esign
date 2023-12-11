// Copyright 2019 James Cote
// All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package esign_test

import (
	"context"
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/jfcote87/oauth2"
	"github.com/jfcote87/testutils"

	"github.com/ConsultingMD/esign"
)

const tokenSuccessResponse = `{
	"access_token": "ISSUED_ACCESS_TOKEN",
	"token_type": "Bearer",
	"refresh_token": "ISSUED_REFRESH_TOKEN",
	"expires_in": 28800
  }`

const userInfoSuccessResponse = `{
	"sub": "50d89ab1-dad5-d00d-b410-92ee3110b970",
	"accounts": [
	  {
		"account_id": "fe0b61a3-3b9b-cafe-b7be-4592af32aa9b",
		"is_default": true,
		"account_name": "World Wide Co",
		"base_uri": "https://gotest.docusign.net"
	  },
	  {
		"account_id": "abcd61a3-3b9b-cafe-b7be-4592af32aa9b",
		"is_default": false,
		"account_name": "Account2",
		"base_uri": "https://gotest.docusign.net"
	  }
	],
	"name": "Susan Smart",
	"given_name": "Susan",
	"family_name": "Smart",
	"email": "susan.smart@example.com"
  }`

func getOAuth2ConfigTransport() (*esign.OAuth2Config, *testutils.Transport) {
	testTransport := &testutils.Transport{}
	clx := &http.Client{Transport: testTransport}

	cfg := &esign.OAuth2Config{
		IntegratorKey: "KEY",
		Secret:        "SECRET",
		RedirURL:      "https://www.example.com/token",
		IsDemo:        true,
		HTTPClientFunc: func(ctx context.Context) (*http.Client, error) {
			return clx, nil
		},
	}
	return cfg, testTransport
}

func TestOuauth2Config_AuthURL(t *testing.T) {
	cfg, _ := getOAuth2ConfigTransport()
	authURL := cfg.AuthURL("STATE")
	expectedURL := "https://account-d.docusign.com/oauth/auth?client_id=KEY&redirect_uri=https%3A%2F%2Fwww.example.com%2Ftoken&response_type=code&scope=signature&state=STATE"
	if authURL != expectedURL {
		t.Errorf("expected %s; got %s", expectedURL, authURL)
		return
	}

	// check for %20 replacement
	cfg.ExtendedLifetime = true
	cfg.Prompt = true
	cfg.UIlocales = []string{"en-us"}
	authURL = cfg.AuthURL("STATE")
	expectedURL = "https://account-d.docusign.com/oauth/auth?client_id=KEY&prompt=login&redirect_uri=https%3A%2F%2Fwww.example.com%2Ftoken&response_type=code&scope=signature%20extended&state=STATE&ui_locales=en-us"
	if authURL != expectedURL {
		t.Errorf("expected %s; got %s", expectedURL, authURL)
		return
	}

	cfg.UIlocales = nil
	cfg.Prompt = false
	authURL = cfg.AuthURL("STATE", "ASCOPE", "extended")
	expectedURL = "https://account-d.docusign.com/oauth/auth?client_id=KEY&redirect_uri=https%3A%2F%2Fwww.example.com%2Ftoken&response_type=code&scope=ASCOPE%20extended&state=STATE"
	if authURL != expectedURL {
		t.Errorf("expected %s; got %s", expectedURL, authURL)
		return
	}
	authURL = cfg.AuthURL("STATE", "ASCOPE")
	if authURL != expectedURL {
		t.Errorf("expected %s; got %s", expectedURL, authURL)
		return
	}
}

var exchangeResponseTest = &testutils.RequestTester{
	Host:   "account-d.docusign.com",
	Path:   "/oauth/token",
	Method: "POST",
	Auth:   "Basic S0VZOlNFQ1JFVA==",
	Payload: []byte(
		"code=CODE&grant_type=authorization_code&redirect_uri=https%3A%2F%2Fwww.example.com%2Ftoken",
	),
	ResponseFunc: func(r *http.Request) (*http.Response, error) {
		return testutils.MakeResponse(200, []byte(tokenSuccessResponse), nil), nil
	},
}
var userinfoResponseDemoTest = &testutils.RequestTester{
	Host:   "account-d.docusign.com",
	Path:   "/oauth/userinfo",
	Method: "GET",
	Auth:   "Bearer ISSUED_ACCESS_TOKEN",
	ResponseFunc: func(r *http.Request) (*http.Response, error) {
		return testutils.MakeResponse(200, []byte(userInfoSuccessResponse), nil), nil
	},
}

var userinfoResponseTest = &testutils.RequestTester{
	Host:   "account.docusign.com",
	Path:   "/oauth/userinfo",
	Method: "GET",
	Auth:   "Bearer ISSUED_ACCESS_TOKEN",
	ResponseFunc: func(r *http.Request) (*http.Response, error) {
		return testutils.MakeResponse(200, []byte(userInfoSuccessResponse), nil), nil
	},
}

var refreshResponseTest = &testutils.RequestTester{
	Path:    "/oauth/token",
	Payload: []byte("grant_type=refresh_token&refresh_token=refresh"),
	ResponseFunc: func(r *http.Request) (*http.Response, error) {
		return testutils.MakeResponse(200, []byte(tokenSuccessResponse), nil), nil
	},
}

func TestOAuth2Config_Exchange(t *testing.T) {
	// Test OAuth2Credential flow
	cfg, testTransport := getOAuth2ConfigTransport()

	testTransport.Add(exchangeResponseTest, userinfoResponseDemoTest)
	ctx := context.Background()

	var savedToken *oauth2.Token
	var savedUserInfo *esign.UserInfo

	cfg.CacheFunc = func(cx context.Context, tk oauth2.Token, ui esign.UserInfo) {
		savedToken = &tk
		savedUserInfo = &ui
	}

	ocr, err := cfg.Exchange(ctx, "CODE")
	if err != nil {
		t.Fatalf("expected successful code exchage; got %v", err)
	}
	u, err := ocr.UserInfo(ctx)
	if err != nil {
		t.Fatalf("expected userInfo for Susan Smart; got error %v", err)
	}
	if u.Name != "Susan Smart" {
		t.Fatalf("expected user name Susan Smart; got %s", u.Name)
	}
	if savedToken == nil || savedUserInfo == nil {
		t.Fatalf(
			"token and userinfo should be cached; got savedToken is nil %v and savedUserInfo is nil %v",
			(savedToken == nil),
			(savedUserInfo == nil),
		)
	}

	tk, err := ocr.Token(ctx)
	if err != nil {
		t.Fatalf("expected token; got %v", err)
	}
	cfg.AccountID = "INVALID ACCOUNT"
	if _, err = cfg.Credential(tk, u); err == nil ||
		err.Error() != "no account INVALID ACCOUNT for susan.smart@example.com" {
		t.Fatalf("expected no account INVALID ACCOUNT for susan.smart@example.com; got %v", err)
	}
	cfg.AccountID = "fe0b61a3-3b9b-cafe-b7be-4592af32aa9b"
	if _, err = cfg.Credential(tk, u); err != nil {
		t.Fatalf("expected successful credential; got %v", err)
	}
	if _, err = cfg.Credential(nil, nil); err == nil || err.Error() != "token may not be nil" {
		t.Fatalf("expected \"token may not be nil\"; got %v", err)
	}
}

var tverV2 = &testVersion{
	Host:   "gotest.docusign.net",
	Demo:   "gotest-d.docusign.net",
	Prefix: "/restapi",
	Ver:    "/v2",
}
var tverV21 = &testVersion{
	Host:   "gotest.docusign.net",
	Demo:   "gotest-d.docusign.net",
	Prefix: "/restapi",
	Ver:    "/v2.1",
}

func TestOAuth2Config_Refresh(t *testing.T) {

	cfg, testTransport := getOAuth2ConfigTransport()

	var savedToken *oauth2.Token
	var savedUserInfo *esign.UserInfo

	cfg.CacheFunc = func(cx context.Context, tk oauth2.Token, ui esign.UserInfo) {
		savedToken = &tk
		savedUserInfo = &ui
	}

	testTransport.Add(refreshResponseTest, userinfoResponseDemoTest)

	var tk *oauth2.Token
	ctx := context.Background()
	cfg.IsDemo = true
	ocra, err := cfg.Credential(&oauth2.Token{RefreshToken: "refresh"}, nil)
	if err != nil {
		t.Fatalf("expected successful credential create; got %v", err)
	}
	if tk, err = ocra.Token(ctx); err != nil {
		t.Fatalf("expected token; got %v", err)
	}
	if tk.AccessToken != "ISSUED_ACCESS_TOKEN" {
		t.Fatalf("expected token ISSUED_ACCESS_TOKEN; got %s", tk.AccessToken)
	}

	testTransport.Add(refreshResponseTest, userinfoResponseTest)
	cfg.IsDemo = false
	ocr, err := cfg.Credential(&oauth2.Token{RefreshToken: "refresh"}, nil)
	if err != nil {
		t.Fatalf("expected successful credential create; got %v", err)
	}
	u, err := ocr.UserInfo(ctx)
	if err != nil {
		t.Fatalf("expecte userinfo success; got %v", err)
	}
	if u.Email != "susan.smart@example.com" {
		t.Fatalf("expected email susan.smart@example.com; got %s", u.Email)
	}

	if savedToken == nil || savedUserInfo == nil {
		t.Fatalf(
			"token and userinfo should be cached; got savedToken is nil %v and savedUserInfo is nil %v",
			(savedToken == nil),
			(savedUserInfo == nil),
		)
	}

	testTransport.Add(&testutils.RequestTester{
		Path:   "/restapi/v2/accounts/" + u.Accounts[0].AccountID + "/abc/def",
		Header: http.Header{"Authorization": {"Bearer ISSUED_ACCESS_TOKEN"}},
		Host:   "gotest.docusign.net",
	}, &testutils.RequestTester{
		Path:   "/restapi/v2.1/accounts/" + u.Accounts[0].AccountID + "/abc/def",
		Header: http.Header{"Authorization": {"Bearer ISSUED_ACCESS_TOKEN"}},
		Host:   "gotest-d.docusign.net",
	})
	op := &esign.Op{
		Method:  "GET",
		Path:    "abc/def",
		Version: tverV2,
	}
	if res, err := ocr.AuthDo(ctx, op); err != nil {
		_ = res
		t.Errorf("authdo(GET, abc/def, nil) expected success; got %v", err)

	} else {
		res.Body.Close()
	}

	op.Version = tverV21

	if res, err := ocra.AuthDo(ctx, op); err != nil {
		t.Errorf("authdo(GET, abc/def, VersionV21) expected success; got %v", err)
	} else {
		res.Body.Close()
	}
}

func TestJWTConfig(t *testing.T) {
	var testPK = `test`
	testTransport := &testutils.Transport{}
	clx := &http.Client{Transport: testTransport}

	cfg := esign.JWTConfig{
		IntegratorKey: "KEY",
		PrivateKey:    testPK,
		KeyPairID:     "1234567890123",
		IsDemo:        true,
		HTTPClientFunc: func(ctx context.Context) (*http.Client, error) {
			return clx, nil
		},
	}
	var expectedConsentURL = "https://account-d.docusign.com/oauth/auth?client_id=KEY&redirect_uri=https%3A%2F%2Fwww.docusign.com&response_type=code&scope=signature%20impersonation"
	if userConsentURL := cfg.UserConsentURL("https://www.docusign.com"); userConsentURL != expectedConsentURL {
		t.Fatalf("expected %s; got %s", expectedConsentURL, userConsentURL)
	}
	ocr, _ := cfg.Credential("50d89ab1-dad5-d00d-b410-92ee3110b970", nil, nil)

	var exchangeResponseTest = &testutils.RequestTester{
		Host:   "account-d.docusign.com",
		Path:   "/oauth/token",
		Method: "POST",
		ResponseFunc: func(r *http.Request) (*http.Response, error) {
			return testutils.MakeResponse(200, []byte(tokenSuccessResponse), nil), nil
		},
	}
	var userinfoResponseTest = &testutils.RequestTester{
		Host:   "account-d.docusign.com",
		Path:   "/oauth/userinfo",
		Method: "GET",
		Auth:   "Bearer ISSUED_ACCESS_TOKEN",
		ResponseFunc: func(r *http.Request) (*http.Response, error) {
			return testutils.MakeResponse(200, []byte(userInfoSuccessResponse), nil), nil
		},
	}
	testTransport.Add(exchangeResponseTest, userinfoResponseTest)
	ctx := context.Background()
	tk, err := ocr.Token(ctx)
	if err != nil {
		t.Errorf("expected token; got error %v", err)
	}

	ocr, _ = cfg.Credential("50d89ab1-dad5-d00d-b410-92ee3110b970", tk, nil)

	testTransport.Add(userinfoResponseTest)

	u, err := ocr.UserInfo(ctx)
	if err != nil {
		t.Errorf("userinf error: %v", err)
	}

	testTransport.Add(&testutils.RequestTester{
		Path:   "/restapi/v2/accounts/" + u.Accounts[0].AccountID + "/abc/def",
		Header: http.Header{"Authorization": {"Bearer ISSUED_ACCESS_TOKEN"}},
		Host:   "gotest-d.docusign.net",
	}, &testutils.RequestTester{
		Path:   "/restapi/v2.1/accounts/" + u.Accounts[0].AccountID + "/abc/def",
		Header: http.Header{"Authorization": {"Bearer ISSUED_ACCESS_TOKEN"}},
		Host:   "gotest-d.docusign.net",
	})
	op := &esign.Op{
		Method:  "GET",
		Path:    "abc/def",
		Version: tverV2,
	}

	if res, err := ocr.AuthDo(ctx, op); err != nil {
		_ = res
		t.Errorf("%v", err)
	} else {
		res.Body.Close()
	}
	op.Version = tverV21
	if res, err := ocr.AuthDo(ctx, op); err != nil {
		t.Errorf("%v", err)
	} else {
		res.Body.Close()
	}
}

func TestTokenCredential(t *testing.T) {
	ctx := context.Background()
	testTransport := &testutils.Transport{}
	cred := esign.TokenCredential("ABCDEF", true).
		SetClientFunc(func(ctx context.Context) (*http.Client, error) {
			return &http.Client{Transport: testTransport}, nil
		})
	testOp := &esign.Op{
		Credential: cred,
		Method:     "GET",
		Path:       "testcmd",
		Version:    esign.APIv2,
	}
	_ = testOp
	expectedAuthHeader := http.Header{
		"Authorization": []string{"Bearer ABCDEF"},
	}
	testTransport.Add(
		&testutils.RequestTester{
			Path:     "/oauth/userinfo",
			Header:   expectedAuthHeader,
			Response: testutils.MakeResponse(400, []byte("invalid token"), nil),
		},
		&testutils.RequestTester{
			Path:     "/oauth/userinfo",
			Header:   expectedAuthHeader,
			Response: testutils.MakeResponse(200, []byte(userInfoSuccessResponse), nil),
		},
		&testutils.RequestTester{
			Path:     "/restapi/v2/accounts/fe0b61a3-3b9b-cafe-b7be-4592af32aa9b/testcmd",
			Header:   expectedAuthHeader,
			Response: testutils.MakeResponse(200, []byte("{}"), nil),
		},
		&testutils.RequestTester{
			Path:     "/restapi/v2/accounts/abcd61a3-3b9b-cafe-b7be-4592af32aa9b/testcmd",
			Header:   expectedAuthHeader,
			Response: testutils.MakeResponse(200, []byte("{}"), nil),
		},
	)
	// check for userinfo fail
	switch err := testOp.Do(ctx, nil).(type) {
	case nil:
		t.Errorf("invalid token expected 400 status; got success")
		return
	case *esign.ResponseError:
	default:
		t.Errorf("%v", err)
		return
	}

	if err := testOp.Do(ctx, nil); err != nil {
		t.Errorf("%v", err)
	}
	testOp.Credential = cred.WithAccountID("BAD_ACCT")
	if err := testOp.Do(ctx, nil); err == nil ||
		err.Error() != "no account BAD_ACCT for susan.smart@example.com" {
		t.Errorf("expected no account BAD_ACCT; got %v", err)
		return
	}
	testOp.Credential = cred.WithAccountID("abcd61a3-3b9b-cafe-b7be-4592af32aa9b")
	if err := testOp.Do(ctx, nil); err != nil {
		t.Errorf("expected success; got %v", err)
	}
}

func TestJWTExternalAdminConsentURL(t *testing.T) {
	jwtCfg := esign.JWTConfig{
		IntegratorKey: "INT_KEY",
		IsDemo:        false,
	}
	// invalid authType
	_, err := jwtCfg.ExternalAdminConsentURL("https://www.example.com", "a", "", false)
	if err == nil {
		t.Errorf("expected error; got success")
	}
	// scopes empty
	_, err = jwtCfg.ExternalAdminConsentURL("https://www.example.com", "code", "STATE", false)
	if err == nil {
		t.Errorf("expected error; got success")
	}

	authURL, _ := jwtCfg.ExternalAdminConsentURL(
		"https://www.example.com",
		"code",
		"STATE",
		false,
		"signature",
		"impersonation",
	)
	expectedURL := "https://account.docusign.com/oauth/auth?admin_consent_scope=signature%20impersonation&client_id=INT_KEY&redirect_uri=https%3A%2F%2Fwww.example.com&response_type=code&scope=openid&state=STATE"
	if authURL != expectedURL {
		t.Errorf("expected %s; got %s", expectedURL, authURL)
		return
	}
	authURL, _ = jwtCfg.ExternalAdminConsentURL(
		"https://www.example.com",
		"token",
		"STATE",
		true,
		"signature",
		"impersonation",
	)
	expectedURL = "https://account.docusign.com/oauth/auth?admin_consent_scope=signature%20impersonation&client_id=INT_KEY&prompt=login&redirect_uri=https%3A%2F%2Fwww.example.com&response_type=token&scope=openid&state=STATE"
	if authURL != expectedURL {
		t.Errorf("expected %s; got %s", expectedURL, authURL)
		return
	}
}

type testVersion struct {
	Host   string
	Demo   string
	Prefix string
	Ver    string
}

func (tv *testVersion) Name() string {
	return tv.Ver
}

func (tv *testVersion) ResolveDSURL(
	u *url.URL,
	host string,
	accountID string,
	isDemo bool,
) *url.URL {
	if tv == nil {
		return u
	}
	newURL := *u
	newURL.Scheme = "https"
	newURL.Host = tv.Host
	if isDemo {
		newURL.Host = tv.Demo
	}

	if !strings.HasPrefix(u.Path, "/") {
		newURL.Path = tv.Prefix + tv.Ver + "/accounts/" + accountID + "/" + u.Path
		return &newURL
	}
	newURL.Path = tv.Prefix + u.Path
	return &newURL
}
