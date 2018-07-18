package samrtc

import (
	"fmt"
	"strconv"
	//"time"
)

type Option func(*SamRTCServer) error

//SetSamHost sets the host of the client's SAM bridge
func SetSamHost(s string) func(*SamRTCServer) error {
	return func(c *SamRTCServer) error {
		c.samHost = s
		return nil
	}
}

//SetSamPort sets the port of the client's SAM bridge
func SetSamPort(v string) func(*SamRTCServer) error {
	return func(c *SamRTCServer) error {
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

//SetSamPortInt sets the port of the client's SAM bridge
func SetSamPortInt(v int) func(*SamRTCServer) error {
	return func(c *SamRTCServer) error {
		if v < 65536 && v > -1 {
			c.samPort = strconv.Itoa(v)
			return nil
		}
		return fmt.Errorf("Invalid port")
	}
}

//SetSamWhitelist adds a destination to the samRTC Whitelist.
func SetSamWhitelist(s string) func(*SamRTCServer) error {
	return func(c *SamRTCServer) error {
        for _, w := range s.whitelist {
            if w == dest {
                return fmt.Errorf("Destination already exists on whitelist: %s", dest)
            }
        }
        s.whitelist = append(s.whitelist, dest)
        return nil
    }
}
