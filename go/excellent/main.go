package main

import "fmt"

func EvenOrOdd(number int) string {
	if number%2 == 0 {
		return "even"
	} else {
		return "odd"
	}
}

func main() {
	result := EvenOrOdd(10)
	fmt.Println(result)
}
