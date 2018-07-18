package samrtc

import (
	"testing"
)

func TestCreateSamHTTPOptionsSetSamHost(t *testing.T) {
	_, e := NewSamRTCServerFromOptions(
		SetSamHost("127.0.0.1"),
		SetSamTunName("SamHostTest"),
		SetSamVerbose(true),
	)
	if e != nil {
		t.Fatal("sam-http-options_test.go Host setting error", e.Error())
	}
}

func TestCreateSamHTTPOptionsSetSamPort(t *testing.T) {
	_, e := NewSamRTCServerFromOptions(
		SetSamPort("7656"),
		SetSamTunName("SamPortTest"),
		SetSamVerbose(true),
	)
	if e != nil {
		t.Fatal("sam-http-options_test.go Port setting error from String", e.Error())
	}
}

func TestCreateSamHTTPOptionsSetSamPortInt(t *testing.T) {
	_, e := NewSamRTCServerFromOptions(
		SetSamPortInt(7656),
		SetSamTunName("SamPortIntTest"),
		SetSamVerbose(true),
	)
	if e != nil {
		t.Fatal("sam-http-options_test.go Port setting error from Int", e.Error())
	}
}

func TestCreateSamHTTPOptionsSetSamWhitelist(t *testing.T) {
	_, e := NewSamRTCServerFromOptions(
		SetSamWhitelist("THISISNOTAREALDESTINATIONBUTABASE64WOULDNORMALLYGOHERE"),
		SetSamTunName("SamWhitelistTest"),
		SetSamVerbose(true),
	)
	if e != nil {
		t.Fatal("sam-http-options_test.go Port setting error from String", e.Error())
	}
}

func TestNewSamRTC(t *testing.T) {
	samrtc, err := NewSamRTCServerFromOptions(
		SetSamVerbose(true),
		SetSamWhitelist("THISISNOTAREALDESTINATIONBUTABASE64WOULDNORMALLYGOHERE"),
	)
	if err != nil {
		t.Fatal(err.Error())
	}
	//base64, base32 := samrtc.GetServerAddresses()
	go samrtc.Serve()
}
