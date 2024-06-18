package main

import (
	"flag"

	"github.com/golang/glog"
)

func main() {

	flag.Set("logtostderr", "1")
	flag.Parse()

	glog.Infoln("Arrays and Linked Lists .....")
}
