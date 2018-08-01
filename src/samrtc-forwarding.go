package samrtc

import (
	"fmt"
	//"io"
	"log"
	//"net"
	"strings"
)

import (
	"github.com/eyedeekay/sam-forwarder"
	"github.com/eyedeekay/sam-forwarder/config"
)

//SamRTCServer is an object representing a server who's sole purpose is to make
//the SAM bridge accessible in the browser.
type SamRTCServer struct {
	server *samforwarder.SAMForwarder

	samHost string
	samPort string
	tunName string
	iniFile string

	verbose bool

	whitelist []string
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
	go s.server.Serve()
	return nil
}

//GetServerAddresses returns the base64 and base32 addresses of the server
func (s *SamRTCServer) GetServerAddresses() (string, string) {
	return s.server.SamKeys.Addr().Base64(), s.server.SamKeys.Addr().Base32()
}

//AddWhitelistDestination adds a client destination to the server whitelist
func (s *SamRTCServer) AddWhitelistDestination(dest string) error {
	//var err error
	for _, w := range s.whitelist {
		if w == dest {
			return fmt.Errorf("Destination already exists on whitelist: %s", dest)
		}
	}

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
	s.iniFile = "/etc/samrtc/samrtc.conf"
	s.verbose = false
	for _, o := range opts {
		if err := o(&s); err != nil {
			return &s, err
		}
	}
	if s.server, err = i2ptunconf.NewSAMForwarderFromConfig(s.iniFile, s.samHost, s.samPort); err != nil {
		return nil, err
	}
	log.Println(s.GetServerAddresses())
	return &s, nil
}

//Log outputs only if verbose==true
func (s *SamRTCServer) Log(i ...interface{}) {
	if s.verbose == true {
		log.Println(i...)
	}
}
