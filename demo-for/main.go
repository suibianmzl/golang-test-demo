package main

import (
	"fmt"
)

func main() {
	ergodicSlice()
}

func ergodicSlice()  {

	arr := []int{
		100,
		200,
		300,
	}

	for i, value := range arr {
		fmt.Println(i, ":", value)
	}

}