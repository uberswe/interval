# interval

Interval is a package which uses Golangs time.Duration to perform different functions.

With the DoEvery functions it is possible to call a function in Golang every x seconds. Depending on the function you can do this in a blocking or asyncronous way.

It is also possible to check if enough time has passed between two times. Such as checking if something has executed 7 days before today for example.

To call a function repeatedly at a set interval simply do this

```go
func main() {
	lambda := func(interval time.Duration, time time.Time) {
		log.Printf("do function called %s at %s\n", interval.String(), time.String())
	}
	err := interval.DoEvery("1s", lambda, -1)
	if err != nil {
		log.Panicf("Error: %v", err)
	}
	log.Println("This is never called because DoEvery is blocking")
}
```

By passing -1 to interval.DoEvery the function will run in an infinite loop.

You can use DoEveryAsync to call a function repeatedly in a non-blocking way like so

```go
func main() {
	lambda := func(interval time.Duration, time time.Time) {
		log.Printf("do function called %s at %s\n", interval.String(), time.String())
	}
	exit, err := interval.DoEveryAsync("1s", lambda, -1)
	if err != nil {
		log.Panicf("Error: %v", err)
	}
	log.Println("This is called because DoEveryAsync is non-blocking")
	// sleep for 10 seconds
	time.Sleep(time.Second * 10)
	// exit the async
	exit <- 1
}
```

