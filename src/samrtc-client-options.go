package samrtc

import (
	"fmt"
	"strconv"
	//"time"
)

//Option is a SamRTCClient option
type ClientOption func(*SamRTCClient) error

//SetSamClientHost sets the host of the client's SAM bridge
func SetSamClientHost(s string) func(*SamRTCClient) error {
	return func(c *SamRTCClient) error {
		c.samHost = s
		return nil
	}
}

//SetSamClientTunName sets the host of the client's SAM bridge
func SetSamClientTunName(s string) func(*SamRTCClient) error {
	return func(c *SamRTCClient) error {
		c.tunName = s
		return nil
	}
}

//SetSamClientPort sets the port of the client's SAM bridge
func SetSamClientPort(v string) func(*SamRTCClient) error {
	return func(c *SamRTCClient) error {
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

//SetSamClientPortInt sets the port of the client's SAM bridge
func SetSamClientPortInt(v int) func(*SamRTCClient) error {
	return func(c *SamRTCClient) error {
		if v < 65536 && v > -1 {
			c.samPort = strconv.Itoa(v)
			return nil
		}
		return fmt.Errorf("Invalid port")
	}
}

//SetSamClientWhitelist adds a destination to the samRTC Whitelist.
/*func SetSamClientWhitelist(s string) func(*SamRTCClient) error {
	return func(c *SamRTCClient) error {
		for _, w := range c.whitelist {
			if w == s {
				return fmt.Errorf("Destination already exists on whitelist: %s", s)
			}
		}
		c.whitelist = append(c.whitelist, s)
		return nil
	}
}*/

//SetSamClientVerbose sets the verbosity of the server
func SetSamClientVerbose(b bool) func(*SamRTCClient) error {
	return func(c *SamRTCClient) error {
		c.verbose = b
		return nil
	}
}
