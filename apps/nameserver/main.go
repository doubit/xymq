package main

import (
	"flag"
	"fmt"
	//"github.com/doubit/xymq/nameserver"
	"github.com/doubit/xymq/util"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	flagSet = flag.NewFlagSet("nameserver", flag.ExitOnError)

	showVersion      = flagSet.Bool("version", false, "print version")
	verbose          = flagSet.Bool("verbose", false, "enable verbose logging")
	tcpAddress       = flagSet.String("tcp-address", "0.0.0.0:9812", "<addr>:<port>")
	broadcastAddress = flagSet.String("broadcast-address", "", "address of this nameserver node, (default to the OS hostname)")

	inactiveProducerTimeout = flagSet.Duration("inactive-producer-timeout", 300*time.Second, "duration of time a broker will remain in the active list since its last ping")
	tombstoneLifetime       = flagSet.Duration("tombstone-lifetime", 45*time.Second, "duration of time a broker will remain tombstoned if registration remains")
)

func main() {
	fmt.Println("nameserver....")
	flagSet.Parse(os.Args[1:])

	if *showVersion {
		fmt.Println(util.Version("nameserver"))
		return
	}

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

}
