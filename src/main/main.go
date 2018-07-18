package main

import (
    "flag"
    "log"
    "strings"
    ".."
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

func main(){
    var whitelistAddrs arrayFlags
    samAddrString := flag.String("addr", "127.0.0.1",
		"host: of the SAM bridge")
	samPortString := flag.String("port", "7656",
		":port of the SAM bridge")
    samTunName := flag.String("name", "serverTun",
		":port of the SAM bridge")
    flag.Var(&whitelistAddrs, "addrs", "Subscription URL(Can be specified multiple times)")
    flag.Parse()

    if samForwarder, err := samrtc.NewSamRTCServerFromOptions(
        samrtc.SetSamHost(*samAddrString),
        samrtc.SetSamPort(*samPortString),
        samrtc.SetSamTunName(*samTunName),
        samrtc.SetSamWhitelist(whitelistAddrs.String()),
        ); err != nil {
            log.Fatal(err.Error())
    }else{
        samForwarder.Serve()
    }

}
