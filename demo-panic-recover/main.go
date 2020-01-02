package main

import (
	"errors"
	"fmt"
)

func recove()  {
	defer func() {
		r := recover()
		if err, ok := r.(error); ok {
			fmt.Println(err)
		} else {
			panic(errors.New("not known error"))
		}
	}()

	a :=5/0
	fmt.Println(a)
}

func main() {
	recove()
}
