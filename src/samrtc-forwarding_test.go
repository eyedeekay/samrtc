package samrtc

import (
	"testing"
)

func TestNewSamRTC(t *testing.T) {
	samrtc, err := NewSamRTCServer()
	if err != nil {
		t.Fatal(err.Error())
	}
	samrtc.Serve()
}
