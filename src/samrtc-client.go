package samrtc

import (
	//	"fmt"
	"io/ioutil"
	"net"
	"net/http"
)

import (
	"github.com/kpetku/sam3"
)

//SamRTCClient is supposed to be a convenient way to embed this into other applications.
type SamRTCClient struct {
	samHost string
	samPort string

	localHost string
	localPort string

	tunName string

	verbose bool

	samConn *sam3.SAM
	samKeys sam3.I2PKeys

	publishStream *sam3.StreamSession
	connection    net.Conn

	whitelist []string
}

func (s *SamRTCClient) samAddress() string {
	return s.samHost + ":" + s.samPort
}

func (s *SamRTCClient) localAddress() string {
	return s.localHost + ":" + s.localPort
}

//GetBase32 asks the localhost:localport for it's corresponding base32
func (s *SamRTCClient) GetBase32() (string, error) {
	resp, err := http.Get("http://" + s.localAddress() + "/base32")
	if err != nil {
		return err.Error(), err
	}
	defer resp.Body.Close()
	var b []byte
	if b, err = ioutil.ReadAll(resp.Body); err != nil {
		return err.Error(), err
	}
	return string(b), nil
}

//NewSamRTCClient generates a default client
func NewSamRTCClient() (*SamRTCClient, error) {
	return NewSamRTCClientFromOptions()
}

//NewSamRTCClientFromOptions generates a client from functional arguments
func NewSamRTCClientFromOptions(opts ...func(*SamRTCClient) error) (*SamRTCClient, error) {
	var s SamRTCClient
	s.samHost = "127.0.0.1"
	s.samPort = "7656"
	s.localHost = "127.0.0.1"
	s.localPort = "7681"
	s.verbose = false
	s.tunName = "clientTun"
	for _, o := range opts {
		if err := o(&s); err != nil {
			return &s, err
		}
	}
	return &s, nil
}
