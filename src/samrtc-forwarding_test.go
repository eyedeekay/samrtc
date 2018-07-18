package samrtc

import (
	"testing"
)

func TestCreateSamHTTPOptionsSetSamHost(t *testing.T) {
	h, e := NewSamRTCServerFromOptions(
		SetSamHost("127.0.0.1"),
	)
	if e != nil {
		t.Fatal("sam-http-options_test.go Host setting error")
	}
}

func TestCreateSamHTTPOptionsSetSamPort(t *testing.T) {
	h, e := NewSamRTCServerFromOptions(
		SetSamPort("7656"),
	)
	if e != nil {
		t.Fatal("sam-http-options_test.go Port setting error from String")
	}
}

func TestCreateSamHTTPOptionsSetSamPortInt(t *testing.T) {
	h, e := NewSamRTCServerFromOptions(
		SetSamPortInt(7656),
	)
	if e != nil {
		t.Fatal("sam-http-options_test.go Port setting error from Int")
	}
}

func TestCreateSamHTTPOptionsSetSamWhitelist(t *testing.T) {
	h, e := NewSamRTCServerFromOptions(
		SetSamWhitelist("THISISNOTAREALDESTINATIONBUTABASE64WOULDNOTMALLYGOHERE"),
	)
	if e != nil {
		t.Fatal("sam-http-options_test.go Port setting error from String")
	}
}

func TestNewSamRTC(t *testing.T) {
	samrtc, err := NewSamRTCServer()
	if err != nil {
		t.Fatal(err.Error())
	}
	samrtc.Serve()
}
