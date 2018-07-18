package samrtc

import (
	"net"
	"strings"
)

import (
	"github.com/kpetku/sam3"
)

//SamRTCServer is an object representing a server who's sole purpose is to make
//the SAM bridge accessible in the browser.
type SamRTCServer struct {
	samHost string
	samPort string

	samConn *sam3.SAM
	samKeys sam3.I2PKeys

	publishStream *sam3.StreamSession
	publishListen *sam3.StreamListener
	connection    net.Conn

	whitelist []string
}

//Serve
func (s *SamRTCServer) Serve() error {
	var err error
	s.connection, err = s.publishListen.Accept()
	if err != nil {
		return err
	}
	return err
}

func (s *SamRTCServer) getWhitelist() string {
	list := "i2cp.accessList="
	for _, s := range s.whitelist {
		list += s + ","
	}
	return strings.TrimSuffix(list, ",")
}

func (s *SamRTCServer) rtc_options() []string {
	rtc_options := []string{
		"inbound.length=0", "outbound.length=0",
		"inbound.lengthVariance=0", "outbound.lengthVariance=0",
		"inbound.backupQuantity=4", "outbound.backupQuantity=4",
		"inbound.quantity=15", "outbound.quantity=15",
		"i2cp.enableAccessList=true", s.getWhitelist(),
	}
	return rtc_options
}

func (s *SamRTCServer) samAddress() string {
	return s.samHost + ":" + s.samPort
}

//NewSamRTCServer generates a new SamRTCServer
func NewSamRTCServer() (*SamRTCServer, error) {
	return NewSamRTCServerFromOptions()
}

//NewSamRTCServerFromOptions generates a new SamRTCServer using functional
//arguments
func NewSamRTCServerFromOptions(opts ...func(*SamRTCServer) error) (*SamRTCServer, error) {
	var err error
	var s SamRTCServer
	s.samHost = "127.0.0.1"
	s.samPort = "7656"
	for _, o := range opts {
		if err := o(&s); err != nil {
			return &s, err
		}
	}
	if s.samConn, err = sam3.NewSAM(s.samAddress()); err != nil {
		return nil, err
	}
	if s.samKeys, err = s.samConn.NewKeys(); err != nil {
		return nil, err
	}
	if s.publishStream, err = s.samConn.NewStreamSession("serverTun", s.samKeys, s.rtc_options()); err != nil {
		return nil, err
	}
	if s.publishListen, err = s.publishStream.Listen(); err != nil {
		return nil, err
	}
	return &s, nil
}
