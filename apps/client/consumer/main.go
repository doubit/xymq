package main

import (
	"fmt"
	"github.com/doubit/xymq/client/consumer"
)

func main() {
	consumer.DoNothing()
	fmt.Println("consumer...")
}
