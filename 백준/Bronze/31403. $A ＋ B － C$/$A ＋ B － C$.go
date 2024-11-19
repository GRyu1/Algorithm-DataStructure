package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	a, _ := reader.ReadString('\n')
	b, _ := reader.ReadString('\n')
	c, _ := reader.ReadString('\n')

	numA, _ := strconv.Atoi(strings.TrimSpace(a))
	numB, _ := strconv.Atoi(strings.TrimSpace(b))
	numC, _ := strconv.Atoi(strings.TrimSpace(c))
	numAB, _ := strconv.Atoi(strings.TrimSpace(a) + strings.TrimSpace(b))

	fmt.Println(numA + numB - numC)
	fmt.Println(numAB - numC)
}
