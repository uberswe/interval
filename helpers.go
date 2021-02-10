package interval

import (
	"errors"
	"time"
)

// interfaceToDuration takes a string or time.Duration and returns a time.Duration
func interfaceToDuration(interval interface{}) (time.Duration, error) {
	if t, ok := interval.(time.Duration); ok {
		return t, nil
	} else if t, ok := interval.(string); ok {
		return time.ParseDuration(t)
	}
	return time.Duration(0), errors.New("interval is of invalid type")
}
