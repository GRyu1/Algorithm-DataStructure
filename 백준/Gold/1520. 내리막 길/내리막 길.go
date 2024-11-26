package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

var (
	dp    [][]int
	route [][]int
	m, n  int
	dx    [4]int = [4]int{0, 0, 1, -1}
	dy    [4]int = [4]int{1, -1, 0, 0}
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	line, _ := reader.ReadString('\n')
	arrLine := strings.Fields(line)
	m, _ = strconv.Atoi(arrLine[0])
	n, _ = strconv.Atoi(arrLine[1])

	route = make([][]int, m)
	dp = make([][]int, m)

	for i := 0; i < m; i++ {
		route[i] = make([]int, n)
		dp[i] = make([]int, n)

		line, _ = reader.ReadString('\n')
		arrLine = strings.Fields(line)
		for j := 0; j < n; j++ {
			route[i][j], _ = strconv.Atoi(arrLine[j])
			dp[i][j] = -1
		}
	}
	writer.WriteString(strconv.Itoa(countRoute(0, 0)))
}

func countRoute(x int, y int) int {
	if x == m-1 && y == n-1 {
		return 1
	}
	if dp[x][y] != -1 {
		return dp[x][y]
	}
	dp[x][y] = 0
	for i := 0; i < 4; i++ {
		nx := x + dx[i]
		ny := y + dy[i]
		if nx < 0 || nx >= m || ny < 0 || ny >= n {
			continue
		}
		if route[nx][ny] < route[x][y] {
			dp[x][y] += countRoute(nx, ny)
		}
	}
	return dp[x][y]
}
