package main

import (
	".."
	"flag"
	"log"
	"strings"
)

type arrayFlags []string

func (i *arrayFlags) String() string {
	var r string
	for _, s := range *i {
		r += s + ","
	}
	return strings.TrimSuffix(r, ",")
}

func (i *arrayFlags) Set(value string) error {
	*i = append(*i, value)
	return nil
}

func main() {
	var whitelistAddrs arrayFlags
	samAddrString := flag.String("shost", "127.0.0.1",
		"host: of the SAM bridge")
	samPortString := flag.String("sport", "7656",
		":port of the SAM bridge")
	samTunName := flag.String("sname", "serverTun",
		":port of the SAM bridge")
	localHost := flag.String("lhost", "127.0.0.1",
		"host: of the local whitelisting control tunnel")
	localPort := flag.String("lport", "7681",
		":port of the local whitelisting control tunnel")
	verbosity := flag.Bool("verbose", false, "enable verbose output")
	flag.Var(&whitelistAddrs, "addrs", "Subscription URL(Can be specified multiple times)")
	flag.Parse()

	if samForwarder, err := samrtc.NewSamRTCHostFromOptions(
		samrtc.SetHostSamHost(*samAddrString),
		samrtc.SetHostSamPort(*samPortString),
		samrtc.SetHostLocalHost(*localHost),
		samrtc.SetHostLocalPort(*localPort),
		samrtc.SetHostSamTunName(*samTunName),
		samrtc.SetHostSamWhitelist(whitelistAddrs.String()),
		samrtc.SetHostSamVerbose(*verbosity),
	); err != nil {
		log.Fatal(err.Error())
	} else {
		//defer samforwarder.Close()
		samForwarder.Serve()
	}

}
