package interval

import "time"

func IsDue(last time.Time, interval interface{}) (bool, error) {
	duration, err := interfaceToDuration(interval)
	if err != nil {
		return false, err
	}
	return IsDueWithDuration(last, duration), nil
}

func IsDueWithDuration(last time.Time, interval time.Duration) bool {
	return IsDueAtTimeWithDuration(last, interval, time.Now().Local())
}

func IsDueAtTime(last time.Time, interval interface{}, at time.Time) (bool, error) {
	duration, err := interfaceToDuration(interval)
	if err != nil {
		return false, err
	}
	return IsDueAtTimeWithDuration(last, duration, at), nil
}

func IsDueAtTimeWithDuration(last time.Time, interval time.Duration, at time.Time) bool {
	return last.Local().Add(interval).Before(at)
}
