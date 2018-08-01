package samrtc

import (
	"log"
	"testing"
	"time"
)

func TestSamRTCHost(t *testing.T) {
	var samForwarder *SamRTCHost
	var err error
	var base32 string
	if samForwarder, err = NewSamRTCHostFromOptions(
		SetHostSamHost("127.0.0.1"),
		SetHostSamPort("7656"),
		SetHostLocalHost("127.0.0.1"),
		SetHostLocalPort("7681"),
		SetHostSamTunName("testSamTun"),
		SetHostSamVerbose(true),
		SetHostSamIniFile("../etc/samrtc/samrtc.conf"),
	); err != nil {
		t.Fatal(err.Error())
	}
	go samForwarder.Serve()
	log.Println("Waiting 1 second")
	time.Sleep(1 * time.Second)
	if client, err := NewSamRTCClient(); err != nil {
		t.Fatal(err.Error())
	} else {
		if base32, err = client.GetBase32(); err != nil {
			t.Fatal(err.Error())
		}
	}
	log.Println(base32)
}
