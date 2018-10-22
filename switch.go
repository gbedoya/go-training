package main

import (
	"fmt"
	"time"
)

func Switch(){

	i := 2
	fmt.Println("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It is before noon")
	default:
		fmt.Println("It is after noon")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I am a bool")
		case int:
			fmt.Println("I am an int")
		default:
			fmt.Printf("Don't know type %T\n", t)

		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}