package main

import (
    "flag"
    "log"
    ".."
)

func main(){
    samAddrString := flag.String("bridge-addr", "127.0.0.1",
		"host: of the SAM bridge")
	samPortString := flag.String("bridge-port", "7656",
		":port of the SAM bridge")
    flag.Parse()

    if samForwarder, err := samrtc.NewSamRTCServerFromOptions(
        samrtc.SetSamHost(*samAddrString),
        samrtc.SetSamPort(*samPortString),
        ); err != nil {
            log.Fatal(err.Error())
    }else{
        samForwarder.Serve()
    }

}
