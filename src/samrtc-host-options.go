package samrtc

import (
	"fmt"
	"strconv"
	//"time"
)

//HostOption is a SamRTCHost option
type HostOption func(*SamRTCHost) error

//SetHostSamHost sets the host of the client's SAM bridge
func SetHostSamHost(s string) func(*SamRTCHost) error {
	return func(c *SamRTCHost) error {
		c.samHost = s
		return nil
	}
}

//SetHostLocalHost sets the host of the whitelist-managing server
func SetHostLocalHost(s string) func(*SamRTCHost) error {
	return func(c *SamRTCHost) error {
		c.host = s
		return nil
	}
}

//SetHostSamTunName sets the host of the client's SAM bridge
func SetHostSamTunName(s string) func(*SamRTCHost) error {
	return func(c *SamRTCHost) error {
		c.tunName = s
		return nil
	}
}

//SetHostLocalPort sets the port of the whitelist-managing server
func SetHostLocalPort(v string) func(*SamRTCHost) error {
	return func(c *SamRTCHost) error {
		port, err := strconv.Atoi(v)
		if err != nil {
			return fmt.Errorf("Invalid port; non-number")
		}
		if port < 65536 && port > -1 {
			c.port = v
			return nil
		}
		return fmt.Errorf("Invalid port")
	}
}

//SetHostSamPort sets the port of the client's SAM bridge
func SetHostSamPort(v string) func(*SamRTCHost) error {
	return func(c *SamRTCHost) error {
		port, err := strconv.Atoi(v)
		if err != nil {
			return fmt.Errorf("Invalid port; non-number")
		}
		if port < 65536 && port > -1 {
			c.samPort = v
			return nil
		}
		return fmt.Errorf("Invalid port")
	}
}

//SetHostSamPortInt sets the port of the client's SAM bridge
func SetHostSamPortInt(v int) func(*SamRTCHost) error {
	return func(c *SamRTCHost) error {
		if v < 65536 && v > -1 {
			c.samPort = strconv.Itoa(v)
			return nil
		}
		return fmt.Errorf("Invalid port")
	}
}

//SetHostSamWhitelist adds a destination to the samRTC Whitelist.
func SetHostSamWhitelist(s string) func(*SamRTCHost) error {
	return func(c *SamRTCHost) error {
		for _, w := range c.whitelist {
			if w == s {
				return fmt.Errorf("Destination already exists on whitelist: %s", s)
			}
		}
		c.whitelist = append(c.whitelist, s)
		return nil
	}
}

//SetHostSamVerbose sets the verbosity of the server
func SetHostSamVerbose(b bool) func(*SamRTCHost) error {
	return func(c *SamRTCHost) error {
		c.verbose = b
		return nil
	}
}
