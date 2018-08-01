package samrtc

import (
	//	"fmt"
	"io/ioutil"
	"log"
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

	localSamHost string
	localSamPort string

	tunName string

	verbose bool

	samConn       *sam3.SAM
	samKeys       sam3.I2PKeys
	publishStream *sam3.StreamSession
	connection    net.Conn
	Conn          *sam3.SAMConn
	remoteSamAddr sam3.I2PAddr

	whitelist []string
}

func (s *SamRTCClient) samAddress() string {
	return s.samHost + ":" + s.samPort
}

func (s *SamRTCClient) localAddress() string {
	return s.localHost + ":" + s.localPort
}

func (s *SamRTCClient) localSamAddress() string {
	return s.localSamHost + ":" + s.localSamPort
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

func (s *SamRTCClient) Connect() (*sam3.SAMConn, error) {
	var err error
	var b32 string
	if s.samConn, err = sam3.NewSAM(s.localSamAddress()); err != nil {
		return nil, err
	}
	log.Println("SAM Bridge connection established.")
	if s.samKeys, err = s.samConn.NewKeys(); err != nil {
		return nil, err
	}
	log.Println("Destination keys generated, tunnel name:", s.tunName, ".")
	if s.publishStream, err = s.samConn.NewStreamSession(s.tunName, s.samKeys, s.rtcOptions()); err != nil {
		log.Println("Stream Creation error:", err.Error())
		return nil, err
	}
	log.Println("SAM stream session established.")
	if b32, err = s.GetBase32(); err != nil {
		log.Println("")
		return nil, err
	}
	if s.remoteSamAddr, err = s.samConn.Lookup(b32); err != nil {
		log.Println("")
		return nil, err
	}
	if s.Conn, err = s.publishStream.DialI2P(s.remoteSamAddr); err != nil {
		log.Println("")
		return nil, err
	}
	return s.Conn, nil
}

func (s *SamRTCClient) rtcOptions() []string {
	rtcOptions := []string{
		"inbound.length=0", "outbound.length=0",
		"inbound.allowZeroHop=true", "outbound.allowZeroHop=true",
		"inbound.lengthVariance=0", "outbound.lengthVariance=0",
		"inbound.backupQuantity=4", "outbound.backupQuantity=4",
		"inbound.quantity=4", "outbound.quantity=4",
		"i2cp.reduceIdleTime=300000", "i2cp.reduceOnIdle=true", "i2cp.reduceQuantity=2",
		"i2cp.closeIdleTime=1200000", "i2cp.closeOnIdle=true",
		"i2cp.dontPublishLeaseSet=true", "i2cp.encryptLeaseSet=true",
	}
	return rtcOptions
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
	s.localSamHost = "127.0.0.1"
	s.localSamPort = "7686"
	s.verbose = false
	s.tunName = "clientTun"
	for _, o := range opts {
		if err := o(&s); err != nil {
			return &s, err
		}
	}
	return &s, nil
}
