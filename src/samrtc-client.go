package samrtc

import (
	//	"fmt"
	//	"log"
	"net"
	//	"strings"
)

import (
	"github.com/kpetku/sam3"
)

//SamRTCClient is supposed to be a convenient way to embed this into other applications.
type SamRTCClient struct {
	samHost string
	samPort string
	tunName string

	verbose bool

	samConn *sam3.SAM
	samKeys sam3.I2PKeys

	publishStream *sam3.StreamSession
	//publishListen *sam3.StreamListener
	connection net.Conn

	whitelist []string
}

func (s *SamRTCClient) samAddress() string {
	return s.samHost + ":" + s.samPort
}

/*
func (s *SamRTCClient) Get() {

}
*/

//NewSamRTCClient generates a default client
func NewSamRTCClient() (*SamRTCClient, error) {
	return NewSamRTCClientFromOptions()
}

//NewSamRTCClientFromOptions generates a client from functional arguments
func NewSamRTCClientFromOptions(opts ...func(*SamRTCClient) error) (*SamRTCClient, error) {
	var s SamRTCClient
	s.samHost = "127.0.0.1"
	s.samPort = "7656"
	s.verbose = false
	s.tunName = "clientTun"
	for _, o := range opts {
		if err := o(&s); err != nil {
			return &s, err
		}
	}
	return &s, nil
}
