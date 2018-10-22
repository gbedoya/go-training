package main

import (
	"fmt"
	"time"
)

func Time() {
	p := fmt.Println
	now := time.Now()
	p(now)

	then := time.Date(
		2009, 11, 17, 20, 34, 58, 123456789, time.UTC)
	p(then)

	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())
	p(then.Weekday())

	// These methods compare two times
	p(then.Before(now))
	p(then.After(now))
	p(then.Equal(now))

	// The sub methods returns a Duration representing the interval between two times.
	diff := now.Sub(then)
	p(diff)
	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())

	p(then.Add(diff))
	p(then.Add(-diff))
}
