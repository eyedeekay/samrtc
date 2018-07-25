package samrtc

import (
	"fmt"
    "log"
	"net/http"
)

type SamRTCHost struct {
	host string
	port string

	samHost string
	samPort string

    serve bool
    tunName string
    whitelist []string
    verbose bool

	forwarder *SamRTCServer
}

func (s *SamRTCHost) AddWhitelistMember(in string) error {
    if s.serve == false {
        var err error
        s.forwarder, err = NewSamRTCServerFromOptions(
            SetSamHost(s.samHost),
            SetSamPort(s.samPort),
            SetSamWhitelist(in),
            SetSamVerbose(s.verbose),
            SetSamTunName(s.tunName),
        )
        if err != nil {
            return err
        }
        go s.forwarder.Serve()
        return nil
    }
    if err := s.forwarder.AddWhitelistDestination(in); err != nil {
        return err
    }
    go s.forwarder.Serve()
	return nil
}

func (s *SamRTCHost) Serve() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprint(w, "Adding Whitelist Member:", r.URL.Path)
        if s.AddWhitelistMember(r.URL.Path) != nil {
            return
        }
        s.serve=true
		fmt.Fprint(w, "Added Whitelist Member:", r.URL.Path)
	})

    log.Println("serving on", s.host, ":", s.port)
	if e := http.ListenAndServe(s.host+":"+s.port, nil); e != nil {
        log.Println(e.Error())
        return
    }
}

func NewSamRTCHost() (*SamRTCHost, error) {
	return NewSamRTCHostFromOptions()
}

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
	return &s, nil
}
