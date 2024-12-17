package main

import (
	"fmt"
	"log"

	"idp4897/greeting"
)

func main() {
	// setting log
	log.SetPrefix("greetings: ")
	log.SetFlags(0)


	names := []string{"Care", "Surachok", "Hemadhulin"}

	messages, err := greeting.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}

	for _, message:= range messages {
		fmt.Println(message)
	}

}