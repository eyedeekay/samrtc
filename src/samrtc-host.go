package samrtc

import (
	"fmt"
	"log"
	"net/http"
)

//SamRTCHost is a server that manages a whitelist-managed SAM over i2p destination.
type SamRTCHost struct {
	host string
	port string

	samHost string
	samPort string
	iniFile string

	serve   bool
	tunName string
	verbose bool

	forwarder *SamRTCServer
}

//AddWhitelistMember adds a whitelist member to the forwarder and restarts it
func (s *SamRTCHost) AddWhitelistMember(in string) error {
	if s.serve == false {
		go s.forwarder.Serve()
		return nil
	}
	if err := s.forwarder.AddWhitelistDestination(in); err != nil {
		return err
	}
	go s.forwarder.Serve()
	return nil
}

//Serve starts the server
func (s *SamRTCHost) Serve() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Adding Whitelist Member:", r.URL.Path)
		if s.AddWhitelistMember(r.URL.Path) != nil {
			return
		}
		s.serve = true
		fmt.Fprint(w, "Added Whitelist Member:", r.URL.Path)
	})
	http.HandleFunc("/base32", func(w http.ResponseWriter, r *http.Request) {
		_, b := s.forwarder.GetServerAddresses()
		fmt.Fprint(w, b)
	})

	log.Println("serving on", s.host, ":", s.port)
	if e := http.ListenAndServe(s.host+":"+s.port, nil); e != nil {
		log.Println(e.Error())
		return
	}
}

//NewSamRTCHost creates a new default server
func NewSamRTCHost() (*SamRTCHost, error) {
	return NewSamRTCHostFromOptions()
}

//NewSamRTCHostFromOptions creates a new server with functional options
func NewSamRTCHostFromOptions(opts ...func(*SamRTCHost) error) (*SamRTCHost, error) {
	var s SamRTCHost
	s.host = "127.0.0.1"
	s.port = "7681"
	s.samHost = "127.0.0.1"
	s.samPort = "7656"
	s.tunName = "SAMTun"
	s.verbose = false
	s.serve = false
	for _, o := range opts {
		if err := o(&s); err != nil {
			return &s, err
		}
	}
	var err error
	s.forwarder, err = NewSamRTCServerFromOptions(
		SetSamHost(s.samHost),
		SetSamPort(s.samPort),
		SetSamVerbose(s.verbose),
		SetSamTunName(s.tunName),
		SetSamIniFile(s.iniFile),
	)
	if err != nil {
		return nil, err
	}
	return &s, nil
}

//NewEmbedSamRTCHostFromOptions creates a new server with functional options
func NewEmbedSamRTCHostFromOptions(opts ...func(*SamRTCHost) error) error {
	var s SamRTCHost
	s.host = "127.0.0.1"
	s.port = "7681"
	s.samHost = "127.0.0.1"
	s.samPort = "7656"
	s.tunName = "SAMTun"
	s.verbose = false
	s.serve = false
	for _, o := range opts {
		if err := o(&s); err != nil {
			return err
		}
	}
	var err error
	s.forwarder, err = NewSamRTCServerFromOptions(
		SetSamHost(s.samHost),
		SetSamPort(s.samPort),
		SetSamVerbose(s.verbose),
		SetSamTunName(s.tunName),
		SetSamIniFile(s.iniFile),
	)
	if err != nil {
		return err
	}
	go s.Serve()
	return nil
}
