package samrtc

import (
	"fmt"
	"strconv"
	"strings"
	//"time"
)

//Option is a SamRTCServer option
type Option func(*SamRTCServer) error

//SetSamHost sets the host of the client's SAM bridge
func SetSamHost(s string) func(*SamRTCServer) error {
	return func(c *SamRTCServer) error {
		c.samHost = s
		return nil
	}
}

//SetSamTunName sets the name of the tunnel used by SAM
func SetSamTunName(s string) func(*SamRTCServer) error {
	return func(c *SamRTCServer) error {
		c.tunName = s
		return nil
	}
}

//SetSamIniFile sets the path of the file to load settings from
func SetSamIniFile(s string) func(*SamRTCServer) error {
	return func(c *SamRTCServer) error {
		c.iniFile = s
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
func SetSamWhitelist(x string) func(*SamRTCServer) error {
	return func(c *SamRTCServer) error {
		if x == "" {
			return nil
		}
		s := strings.Split(x, ",")
		if len(s) > 0 {
			for _, i := range s {
				for _, j := range c.whitelist {
					if j == i {
						return nil
					}
				}
				c.whitelist = append(c.whitelist, i)
			}
		}
		return nil
	}
}

//SetSamVerbose sets the verbosity of the server
func SetSamVerbose(b bool) func(*SamRTCServer) error {
	return func(c *SamRTCServer) error {
		c.verbose = b
		return nil
	}
}
