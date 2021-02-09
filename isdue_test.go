package interval

import (
	"testing"
	"time"
)

type IsDueTestStruct struct {
	Last             time.Time
	StringInterval   string
	DurationInterval time.Duration
	At               time.Time
	Expected         bool
}

func TestIsDue(t *testing.T) {
	parameters := []IsDueTestStruct{
		{Last: time.Now().Local().Add(-10 * time.Hour), StringInterval: "5h", DurationInterval: time.Hour * 5, Expected: true},
		{Last: time.Now().Local().Add(-10 * time.Hour), StringInterval: "11h", DurationInterval: time.Hour * 11, Expected: false},
	}

	for _, p := range parameters {
		result, err := IsDue(p.Last, p.StringInterval)
		if err != nil {
			t.Errorf("Error: %v", err)
		}
		if result != p.Expected {
			t.Errorf("Unexpected result from IsDue(%s, %s), got: %t, want: %t.", p.Last.String(), p.StringInterval, result, p.Expected)
		}
	}

	for _, p := range parameters {
		result, err := IsDue(p.Last, p.DurationInterval)
		if err != nil {
			t.Errorf("Error: %v", err)
		}
		if result != p.Expected {
			t.Errorf("Unexpected result from IsDue(%s, %s), got: %t, want: %t.", p.Last.String(), p.DurationInterval, result, p.Expected)
		}
	}
}

func TestIsDueAtTime(t *testing.T) {
	parameters := []IsDueTestStruct{
		{Last: time.Now().Local().Add(-10 * time.Hour), StringInterval: "5h", DurationInterval: time.Hour * 5, Expected: true, At: time.Now().Local()},
		{Last: time.Now().Local().Add(-10 * time.Hour), StringInterval: "11h", DurationInterval: time.Hour * 11, Expected: false, At: time.Now().Local()},
	}

	for _, p := range parameters {
		result, err := IsDueAtTime(p.Last, p.StringInterval, p.At)
		if err != nil {
			t.Errorf("Error: %v", err)
		}
		if result != p.Expected {
			t.Errorf("Unexpected result from IsDueAtTime(%s, %s), got: %t, want: %t.", p.Last.String(), p.StringInterval, result, p.Expected)
		}
	}

	for _, p := range parameters {
		result, err := IsDueAtTime(p.Last, p.DurationInterval, p.At)
		if err != nil {
			t.Errorf("Error: %v", err)
		}
		if result != p.Expected {
			t.Errorf("Unexpected result from IsDueAtTime(%s, %s), got: %t, want: %t.", p.Last.String(), p.DurationInterval, result, p.Expected)
		}
	}
}
