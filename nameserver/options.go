package nameserver

// import (
// 	"log"
// 	"os"
// 	"time"
// )

// type Options struct {
// 	Verbose bool `flag:"verbose"`

// 	TCPAddress       string `flag:"tcp-address"`
// 	BroadcastAddress string `flag:"broadcast-address"`

// 	InactiveProducerTimeout time.Duration `flag:"inactive-producer-timeout"`
// 	TombstoneLifetime       time.Duration `flag:"tombstone-lifetime"`

// 	Logger logger
// }

// func NewOptions() *xymqNameServerOptions {
// 	hostname, err := os.Hostname()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	return &nsqlookupdOptions{
// 		TCPAddress:       "0.0.0.0:9812",
// 		BroadcastAddress: hostname,

// 		InactiveProducerTimeout: 300 * time.Second,
// 		TombstoneLifetime:       45 * time.Second,

// 		Logger: log.New(os.Stderr, "[nameserver] ", log.Ldate|log.Ltime|log.Lmicroseconds),
// 	}
// }
