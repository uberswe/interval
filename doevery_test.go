package interval

import (
	"testing"
	"time"
)

func TestDoEvery(t *testing.T) {
	start := time.Now().Local()
	expected := 5
	count := 0
	lambda := func(interval time.Duration, time time.Time, extra interface{}) {
		t.Logf("do function called")
		count++
	}
	err := DoEvery("1ms", nil, lambda, expected)
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

func TestDoEveryAsync(t *testing.T) {
	start := time.Now().Local()
	expected := 5
	count := 0
	lambda := func(interval time.Duration, time time.Time, extra interface{}) {
		t.Logf("do function called")
		count++
	}
	_, err := DoEveryAsync("1ms", nil, lambda, expected)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	finish := time.Now().Local().Sub(start)
	if finish > time.Millisecond {
		t.Errorf("Expected DoEveryAsync not to block but more than %dms, got %s", 1, finish.String())
	}
	if count > 0 {
		t.Errorf("Expected lambda function to run %d times async, got %d", 0, count)
	}
	time.Sleep(time.Millisecond * time.Duration(expected+1))
	if count != expected {
		t.Errorf("Expected lambda function to run %d times, got %d", expected, count)
	}
}

func TestDoEveryAsyncWithExit(t *testing.T) {
	start := time.Now().Local()
	expected := 5
	count := 0
	lambda := func(interval time.Duration, time time.Time, extra interface{}) {
		t.Logf("do function called")
		count++
	}
	exit, err := DoEveryAsync("1ms", nil, lambda, expected)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	finish := time.Now().Local().Sub(start)
	if finish > time.Millisecond {
		t.Errorf("Expected DoEveryAsync not to block but more than %dms, got %s", 1, finish.String())
	}
	if count > 0 {
		t.Errorf("Expected lambda function to run %d times async, got %d", 0, count)
	}
	exit <- 1
	time.Sleep(time.Millisecond * time.Duration(expected+1))
	if count > 0 {
		t.Errorf("Expected lambda function to exit before running, got %d", count)
	}
}
