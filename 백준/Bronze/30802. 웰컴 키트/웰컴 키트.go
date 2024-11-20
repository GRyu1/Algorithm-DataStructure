package main

import (
	"fmt"
	"math"
)

func main() {
	var n, t, p int
	var sizeAmount []int

	fmt.Scan(&n)

	var size int
	for i := 0; i < 6; i++ {
		if _, err := fmt.Scan(&size); err != nil {
			break
		}
		sizeAmount = append(sizeAmount, size)
	}

	fmt.Scan(&t, &p)

	sumValue := 0
	for _, value := range sizeAmount {
		sumValue += int(math.Ceil(float64(value) / float64(t)))
	}
	
	fmt.Println(sumValue)
	fmt.Printf("%d %d\n", n/p, n%p)
}