package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	n, m     int
	graph    [][]int
	dp       []int
	visited  []bool
	maxHacks int = -1
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(line)
	numbers := strings.Split(line, " ")
	n, _ = strconv.Atoi(numbers[0])
	m, _ = strconv.Atoi(numbers[1])

	graph = make([][]int, n+1)
	for i := 0; i <= n; i++ {
		graph[i] = []int{}
	}

	for i := 0; i < m; i++ {
		line, _ = reader.ReadString('\n')
		line = strings.TrimSpace(line)
		edges := strings.Split(line, " ")
		a, _ := strconv.Atoi(edges[0])
		b, _ := strconv.Atoi(edges[1])
		graph[b] = append(graph[b], a)
	}

	dp = make([]int, n+1)
	for i := 1; i <= n; i++ {
		visited = make([]bool, n+1)
		solve(i)
	}

	for i := 1; i <= n; i++ {
		if dp[i] == maxHacks {
			fmt.Fprintf(writer, "%d ", i)
		}
	}
}

func solve(start int) {
	queue := []int{start}
	visited[start] = true

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		for _, next := range graph[current] {
			if visited[next] {
				continue
			}
			queue = append(queue, next)
			dp[start]++
			visited[next] = true
		}
	}

	if dp[start] == 0 {
		dp[start]++
	}

	if dp[start] > maxHacks {
		maxHacks = dp[start]
	}
}
