package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

var lis []int
var arr []int
var n int

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	temp, _ := reader.ReadString('\n')
	n, _ = strconv.Atoi(strings.TrimSpace(temp))

	temp, _ = reader.ReadString('\n')
	inputs := strings.Fields(temp)
	arr = make([]int, n)
	for index, value := range inputs {
		arr[index], _ = strconv.Atoi(value)
	}

	lis = []int{arr[0]}
	for i := 0; i < n; i++ {
		if arr[i] > lis[len(lis)-1] {
			lis = append(lis, arr[i])
		} else {
			idx := binarySearch(arr[i])
			lis[idx] = arr[i]
		}
	}

	writer.WriteString(strconv.Itoa(len(lis)))
}

func binarySearch(target int) int {
	left, right := 0, len(lis)-1
	for left <= right {
		mid := (left + right) / 2
		if lis[mid] == target {
			return mid
		} else if lis[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}