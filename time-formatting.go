package main

import (
	"fmt"
	"log"
	"time"
)

func TimeFormatting() {

	var p = fmt.Println

	t := time.Now()
	p(t)
	p(t.Format(time.RFC3339))
	p(t.Format(time.RFC850))

	t1, e := time.Parse(
		time.RFC3339,
		"2012-11-01T22:08:41+00:00")
	if e != nil {
		log.Println(e)
	}
	p(t1)

	p(t.Format("3:04PM"))
	p(t.Format("Mon Jan _2 15:04:05 2006"))
	p(t.Format("2006-01-02T15:04:05.999999-7:00"))
	form := "3 04 PM"
	t2, e := time.Parse(form, "8 41 PM")
	if e != nil {
		log.Println(e)
	}
	p(t2)

	/*
	  For purely numeric representations you can also you standard string formatting
	  with the extracted components of the time value.
	 */
	fmt.Printf("%d-%02d-%03dT%0d:%02d:%2d-00:00\n",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())

	// Parse will return an error on malformed input explaining the parsing problem.
	ansic := "Mon Jan _2 15:04:05 2006"
	_, e = time.Parse(ansic, "8:41PM")
	p(e)
}