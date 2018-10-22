package main

import (
	"errors"
	"fmt"
)

func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("Can't work with 42")
	}
	return arg + 3, nil
}

// It's possible the use custom types of errors by implementing the Error() method on them.
type argError struct {
	arg int
	prob string
}

func (e *argError) Error() string{
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "Can't work with 42"}
	}
	return arg + 3, nil
}

func Errors() {
	// Note the use of an inline error check on the if line is  a common idiom in Go code.
	// (_, i_) = position, value

	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}

	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}

	// If you want to programmatically use the data in the custom error,
	// you'll need to get the error as an instance of the custom  error type
	// via type assertion.
	// Type assertions: https://tour.golang.org/methods/15
	// t, ok := i.(T)
	// If i holds a T,
	// then t will be the underlying value and ok will be true
	_, e := f2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
		fmt.Println(ok)
	}
}
