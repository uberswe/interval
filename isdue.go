package interval

import "time"

// IsDue takes a time and an interval of time.Duration or a string and checks if the duration has passed since the last time and now
func IsDue(last time.Time, interval interface{}) (bool, error) {
	duration, err := interfaceToDuration(interval)
	if err != nil {
		return false, err
	}
	return IsDueWithDuration(last, duration), nil
}

// IsDueWithDuration takes a time and an interval of time.Duration and checks if the duration has passed since the last time and now
func IsDueWithDuration(last time.Time, interval time.Duration) bool {
	return IsDueAtTimeWithDuration(last, interval, time.Now().Local())
}

// IsDueAtTime takes a time and an interval of time.Duration or a string and checks if the duration has passed since the last time and the at time provided
func IsDueAtTime(last time.Time, interval interface{}, at time.Time) (bool, error) {
	duration, err := interfaceToDuration(interval)
	if err != nil {
		return false, err
	}
	return IsDueAtTimeWithDuration(last, duration, at), nil
}

// IsDueAtTimeWithDuration takes a time and an interval of time.Duration and checks if the duration has passed since the last time and the at time provided
func IsDueAtTimeWithDuration(last time.Time, interval time.Duration, at time.Time) bool {
	return last.Local().Add(interval).Before(at)
}
