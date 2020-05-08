// Go programs express error state with error values.

// The error type is a built-in interface.:

// type error interface {
//     Error() string
// }

package main

import (
	"errors"
	"fmt"
)

//constructs a basic error value with the given error message.
func isEven(i int) (bool, error) { //error is Interface
	if i%2 != 0 {
		return false, errors.New("Not an even number")
	}
	return true, nil
}

//using custom types as errors by implementing the Error() method on them.
type myError struct {
	number int
	msg    string
}

func (e *myError) Error() string {
	return fmt.Sprintf("%d %s", e.number, e.msg)
}

func isOdd(i int) (bool, error) {
	if i%2 == 0 {
		return false, &myError{
			number: i,
			msg:    "is not an odd number",
		}
	}
	return true, nil
}

func main() {
	ans, err := isEven(9)
	// /Functions often return an error value,
	//and calling code should handle errors by
	//testing whether the error equals nil
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(ans)
	}

	ans, err = isOdd(10)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(ans)
	}

}
