package interval

import (
	"time"
)

// DoEvery is a blocking function that takes an interval of time.Duration or a string representing a time interval and then calls a provided function every interval
func DoEvery(interval interface{}, extra interface{}, do func(interval time.Duration, time time.Time, extra interface{}), count int) error {
	duration, err := interfaceToDuration(interval)
	if err != nil {
		return err
	}
	DoEveryWithDuration(duration, extra, do, count)
	return nil
}

// DoEveryWithDuration is a blocking function that takes an interval of time.Duration representing a time interval and then calls a provided function every interval
func DoEveryWithDuration(interval time.Duration, extra interface{}, do func(interval time.Duration, time time.Time, extra interface{}), count int) {
	tick := time.Tick(interval)
	for range tick {
		do(interval, time.Now().Local(), extra)
		if count > 0 {
			count--
			if count == 0 {
				return
			}
		}
	}
}

// DoEveryAsync is a non-blocking function that takes an interval of time.Duration or a string representing a time interval and then calls a provided function every interval
func DoEveryAsync(interval interface{}, extra interface{}, do func(interval time.Duration, time time.Time, extra interface{}), count int) (chan int, error) {
	duration, err := interfaceToDuration(interval)
	if err != nil {
		return nil, err
	}
	return DoEveryAsyncWithDuration(duration, extra, do, count)
}

// DoEveryAsyncWithDuration is a non-blocking function that takes an interval of time.Duration representing a time interval and then calls a provided function every interval
func DoEveryAsyncWithDuration(interval time.Duration, extra interface{}, do func(interval time.Duration, time time.Time, extra interface{}), count int) (chan int, error) {
	tick := time.NewTicker(interval)
	exit := make(chan int)
	go func() {
		for {
			select {
			case <-tick.C:
				do(interval, time.Now().Local(), extra)
				if count > 0 {
					count--
					if count == 0 {
						tick.Stop()
						return
					}
				}
			case <-exit:
				tick.Stop()
				return
			}
		}
	}()
	return exit, nil
}
