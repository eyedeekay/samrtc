package samrtc

import ()

type SamRTCClient struct {
	samHost string
	samPort string
	tunName string

	verbose bool
}

func (s *SamRTCClient) samAddress() string {
	return s.samHost + ":" + s.samPort
}

func (s *SamRTCClient) Get() {

}

func NewSamRTCClient() (*SamRTCClient, error) {
	return NewSamRTCClientFromOptions()
}

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
