package samrtc

import (
	"fmt"
	"io"
	"log"
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
	tunName string

	verbose bool

	samConn *sam3.SAM
	samKeys sam3.I2PKeys

	publishStream     *sam3.StreamSession
	publishListen     *sam3.StreamListener
	publishConnection net.Conn

	localConnection net.Conn

	whitelist []string
}

func (s *SamRTCServer) forward() error {
	var err error
	if s.localConnection, err = net.Dial("tcp", s.samAddress()); err != nil {
		return err
	}
	s.Log("Dialed local SAM bridge for forwarding")
	go func() {
		defer s.localConnection.Close()
		defer s.publishConnection.Close()
		io.Copy(s.localConnection, s.publishConnection)
	}()
	go func() {
		defer s.localConnection.Close()
		defer s.publishConnection.Close()
		io.Copy(s.publishConnection, s.localConnection)
	}()
	return nil
}

//Serve a the specified SAM port on an i2p destination
func (s *SamRTCServer) Serve(whitelist ...string) error {
	if len(whitelist) > 0 {
		for _, x := range whitelist {
			s.AddWhitelistDestination(x)
		}
	}
	var err error
	if err = s.checkWhitelist(); err != nil {
		log.Fatal(err.Error())
	}
	_, base32 := s.GetServerAddresses()
	s.Log("Starting server:", s.tunName+":", base32)
	s.publishConnection, err = s.publishListen.Accept()
	if err != nil {
		return err
	}
	s.Log("Server started.")
	for {
		go s.forward()
	}
}

//GetServerAddresses returns the base64 and base32 addresses of the server
func (s *SamRTCServer) GetServerAddresses() (string, string) {
	return s.samKeys.Addr().Base64(), s.samKeys.Addr().Base32()
}

//AddWhitelistDestination adds a client destination to the server whitelist
func (s *SamRTCServer) AddWhitelistDestination(dest string) error {
    var err error
	for _, w := range s.whitelist {
		if w == dest {
			return fmt.Errorf("Destination already exists on whitelist: %s", dest)
		}
	}
    s.Log("Re-initializing Stream Session")
    s.Log("Closing listener")
    if err = s.publishListen.Close(); err != nil {
        return err
    }
    s.Log("Listener closed")
    s.Log("Closing streamsession")
    if err = s.publishStream.Close(); err != nil {
        return err
    }
    s.Log("StreamSession closed")
    s.Log("Re-Opening with new whitelist")
	s.whitelist = append(s.whitelist, dest)
    if s.publishStream, err = s.samConn.NewStreamSession(s.tunName, s.samKeys, s.rtcOptions()); err != nil {
		return err
	}
	s.Log("SAM stream session established")
	if s.publishListen, err = s.publishStream.Listen(); err != nil {
		return err
	}
	s.Log("SAM Listener created")
	return nil
}

func (s *SamRTCServer) getWhitelist() string {
	list := "i2cp.accessList="
	for _, w := range s.whitelist {
		list += w + ","
	}
	return strings.TrimSuffix(list, ",")
}

func (s *SamRTCServer) checkWhitelist() error {
	if len(s.whitelist) < 1 {
		return fmt.Errorf("Never run without a whitelist member %x", len(s.whitelist))
	}
	for i, s := range s.whitelist {
		if len(s) <= 1 {
			if i <= 0 {
				return fmt.Errorf("Whitelist member failed trivial validation, string of length %x detected at index %x", len(s), i)
			}
			break
		}
	}
	return nil
}

func (s *SamRTCServer) rtcOptions() []string {
	rtcOptions := []string{
		"inbound.length=0", "outbound.length=0",
		"inbound.allowZeroHop=true", "outbound.allowZeroHop=true",
		"inbound.lengthVariance=0", "outbound.lengthVariance=0",
		"inbound.backupQuantity=4", "outbound.backupQuantity=4",
		"inbound.quantity=4", "outbound.quantity=4",
		"i2cp.reduceIdleTime=300000", "i2cp.reduceOnIdle=true", "i2cp.reduceQuantity=2",
		"i2cp.closeIdleTime=1200000", "i2cp.closeOnIdle=true",
		"i2cp.encryptLeaseSet=true", "i2cp.enableAccessList=true", s.getWhitelist(),
	}
	return rtcOptions
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
	s.tunName = "serverTun"
	s.verbose = false
	for _, o := range opts {
		if err := o(&s); err != nil {
			return &s, err
		}
	}
	if s.samConn, err = sam3.NewSAM(s.samAddress()); err != nil {
		return nil, err
	}
	s.Log("SAM Bridge connection established")
	if s.samKeys, err = s.samConn.NewKeys(); err != nil {
		return nil, err
	}
	s.Log("Destination keys generated, tunnel name:", s.tunName)
	if s.publishStream, err = s.samConn.NewStreamSession(s.tunName, s.samKeys, s.rtcOptions()); err != nil {
		return nil, err
	}
	s.Log("SAM stream session established")
	if s.publishListen, err = s.publishStream.Listen(); err != nil {
		return nil, err
	}
	s.Log("SAM Listener created")
	log.Println(s.GetServerAddresses())
	return &s, nil
}

//Log outputs only if verbose==true
func (s *SamRTCServer) Log(i ...interface{}) {
	if s.verbose == true {
		log.Println(i...)
	}
}
