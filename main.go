package main

import (
	"fmt"
	"github.com/accesstoken/go-training/chapter4/section4.2/logger"
	"os"
)

func main() {
	var log logger.Logger
	log.Writer = os.Stdout
	//log.Writer = ioutil.Discard
	log.Begin()
	var name string = "Joe"
	log.Logf("The name was %q\n", name)
	var msg string = "Hello world!"
	fmt.Println(msg)
	log.Log("I said:", msg)

	log.End()
}