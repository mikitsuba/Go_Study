package main

import (
	"errors"
	"fmt"
	"log"
)

type sqrtError struct {
	lat  string
	long string
	err  error
}

func (se sqrtError) Error() string {
	return fmt.Sprintf("math error: %v %v %v", se.lat, se.long, se.err)
}

func main() {
	_, err := sqrt(-10.23)
	if err != nil {
		log.Println(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		// write your error code here
		e := errors.New("I need coffee")
		return 0, sqrtError{lat: "50.2289 N", long: "99.4656 W", err: e} // This is type error because type sqrtError has Error() string in its method sets.
	}
	return 42, nil
}
