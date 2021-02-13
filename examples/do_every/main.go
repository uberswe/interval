package main

import (
	"github.com/uberswe/interval"
	"log"
	"time"
)

// This is an example of how to use the DoEvery function

func main() {
	lambda := func(interval time.Duration, time time.Time, extra interface{}) {
		log.Printf("do function called %s at %s\n", interval.String(), time.String())
	}
	err := interval.DoEvery("1s", nil, lambda, -1)
	if err != nil {
		log.Panicf("Error: %v", err)
	}
	log.Println("This is never called because DoEvery is blocking")
}
