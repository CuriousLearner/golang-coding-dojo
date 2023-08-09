package main

import "fmt"

// Fizz buzz

func FizzBuzz(n int) (msg string, err string) {
	msg = ""
	err = ""
	if n < 0 {
		err = string("Negative values not allowed")
	} else if n > 15 {
		err = "Values greater than 15 are not allowed"
	}
	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			msg += fmt.Sprintf("%s\n", "FizzBuzz")
		} else if i%5 == 0 {
			msg += fmt.Sprintf("%s\n", "Buzz")
		} else if i%3 == 0 {
			msg += fmt.Sprintf("%s\n", "Fizz")
		} else {
			msg += fmt.Sprintf("%d\n", i)
		}
	}
	return msg, err
}
