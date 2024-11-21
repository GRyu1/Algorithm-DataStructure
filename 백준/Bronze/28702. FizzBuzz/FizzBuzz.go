package main

import (
	"fmt"
	"strconv"
)

func main() {
	var indexValue, index int
	var data string
	for i := 0; i < 3; i++ {
		fmt.Scan(&data)
		value, _ := strconv.Atoi(data)
		if value != 0 {
			indexValue = value
			index = i
		}
	}
	indexValue += 4 - (index + 1)

	if indexValue%3 == 0 && indexValue%5 == 0 {
		fmt.Println("FizzBuzz")
	} else if indexValue%3 == 0 {
		fmt.Println("Fizz")
	} else if indexValue%5 == 0 {
		fmt.Println("Buzz")
	} else {
		fmt.Println(indexValue)
	}
}
