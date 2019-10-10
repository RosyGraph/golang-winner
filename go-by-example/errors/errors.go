// In Go it's idiomatic to communicate errors via an explicit,
// separate return value. This contrasts with exceptions
// used in languages like Java and Ruby and the overloaded
// single result / error value sometimes used in C. Go's
// approach makes it easy to see which functions return
// errors and to handle them using the same language
// constructs employed for any other, non-error tasks.
package main

import (
	"errors"
	"fmt"
)

func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("can't work with 42")
	}
	return arg + 3, nil
}

type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func main() {
	fmt.Println("vim-go")
}
