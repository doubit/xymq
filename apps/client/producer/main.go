package main

import (
	"fmt"
	"github.com/doubit/xymq/client/producer"
)

func main() {
	producer.DoNothing()
	fmt.Println("producer....")
}
