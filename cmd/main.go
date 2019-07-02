package main

import (
	"fmt"
	"github.com/monimena/daikin-ir"
	"log"
)

func main() {
	manager := daikin.NewManager()
	serial, err := daikin.NewSerial(manager)

	if err != nil {
		log.Fatal(err)
	}

	api := daikin.NewApi(manager)

	c := make(chan error)

	go func() { for { c <- <-serial.Run() }}()
	go func() { for { c <- <-api.Run() }}()

	fmt.Sprintf("%v", <-c)
}
