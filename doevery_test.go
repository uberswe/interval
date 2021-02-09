package interval

import (
	"testing"
	"time"
)

func TestDoEvery(t *testing.T) {
	start := time.Now().Local()
	expected := 5
	count := 0
	lambda := func(interval time.Duration, time time.Time) {
		t.Logf("do function called")
		count++
	}
	err := DoEvery("1ms", lambda, expected)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	if count != expected {
		t.Errorf("Expected lambda function to run %d times, got %d", expected, count)
	}
	finish := time.Now().Local().Sub(start)
	if finish < time.Millisecond*time.Duration(expected) || finish > time.Millisecond*time.Duration(expected+1) {
		t.Errorf("Expected DoEvery to take at least %dms to run but not more than %dms, got %s", expected, expected+1, finish.String())
	}

}
