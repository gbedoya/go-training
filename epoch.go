package main

import (
	"fmt"
	"time"
)

func Epoch() {

	p := fmt.Println

	now := time.Now()
	secs := now.Unix()
	nanos := now.UnixNano()

	p(now)
	p(secs)
	millis := nanos / 1000000
	p(millis)
	p(nanos)

	// You can also convert integer seconds or nanoseconds since
	// the epoch into the corresponding time
	//time.Unix(sec int64, nsec int64)
	p(time.Unix(secs,0))
	p(time.Unix(0, nanos))

}