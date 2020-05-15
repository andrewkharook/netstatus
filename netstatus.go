package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/justincampbell/anybar"
	ping "github.com/sparrc/go-ping"
)

const defaultHost = "1.1.1.1"
const iconUp = "black"
const iconDown = "white"

var port int
var host string

func init() {
	flag.Usage = usage

	flag.IntVar(&port, "port", anybar.DefaultPort, "The port of an AnyBar instance")
	flag.StringVar(&host, "host", defaultHost, "A host to ping")

	flag.Parse()
}

func main() {
	icon := iconUp
	if !isUp() {
		icon = iconDown
	}

	anybar.SendTo(icon, port)
}

func isUp() bool {
	pinger, err := ping.NewPinger(host)
	if err != nil {
		panic(err)
	}
	pinger.Timeout, _ = time.ParseDuration("2s")
	pinger.Run()
	stats := pinger.Statistics()
	if stats.PacketLoss == 100 {
		return false
	}

	return true
}

func usage() {
	fmt.Fprintf(os.Stderr, "usage: %s [options]\n", os.Args[0])
	flag.PrintDefaults()
}
