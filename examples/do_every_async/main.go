package main

import (
	"github.com/uberswe/interval"
	"log"
	"time"
)

// This is an example of how to use the DoEveryAsync function

func main() {
	lambda := func(interval time.Duration, time time.Time, extra interface{}) {
		log.Printf("do function called %s at %s\n", interval.String(), time.String())
	}
	exit, err := interval.DoEveryAsync("1s", nil, lambda, -1)
	if err != nil {
		log.Panicf("Error: %v", err)
	}
	log.Println("This is called because DoEveryAsync is non-blocking")
	// sleep for 10 seconds
	time.Sleep(time.Second * 10)
	// exit the async
	exit <- 1
}
