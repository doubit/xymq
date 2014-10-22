package main

import (
	"fmt"
	"github.com/doubit/xymq/broker"
)

func main() {
	broker.DoNothing()
	fmt.Println("broker...")
}
