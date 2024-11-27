package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

const maxWeight = 500

var (
	dp   [][]bool
	n    int
	arr1 []int
	arr2 []int
)

func sol(i, target int) {
	if i > n {
		return
	}
	if dp[i][target] {
		return
	}

	dp[i][target] = true

	sol(i+1, target)
	if i < n {
		sol(i+1, target+arr1[i])
		sol(i+1, abs(target-arr1[i]))
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	line, _ := reader.ReadString('\n')
	n, _ = strconv.Atoi(strings.TrimSpace(line))

	line, _ = reader.ReadString('\n')
	arr1 = toIntArray(strings.Fields(line))

	reader.ReadString('\n')
	line, _ = reader.ReadString('\n')
	arr2 = toIntArray(strings.Fields(line))

	dp = make([][]bool, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]bool, maxWeight*30+1)
	}

	sol(0, 0)

	for _, weight := range arr2 {
		if weight > maxWeight*30 {
			writer.WriteString("N ")
		} else if dp[n][weight] {
			writer.WriteString("Y ")
		} else {
			writer.WriteString("N ")
		}
	}
}

func toIntArray(strs []string) []int {
	arr := make([]int, len(strs))
	for i, s := range strs {
		arr[i], _ = strconv.Atoi(s)
	}
	return arr
}
