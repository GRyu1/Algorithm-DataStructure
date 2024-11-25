package main

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	m := make(map[int]int)
	maxLen := 0
	left := 0
	count := 0

	line, _ := reader.ReadString('\n')
	line, _ = reader.ReadString('\n')
	arrStr := strings.Fields(line)

	for right, s := range arrStr {
		v, _ := strconv.Atoi(s)
		_, exist := m[v]
		if !exist {
			m[v] = 1
			count++
		} else {
			m[v]++
		}

		for count > 2 {
			v, _ := strconv.Atoi(arrStr[left])
			m[v]--
			if m[v] == 0 {
				delete(m, v)
				count--
			}
			left++
			break
		}
		maxLen = int(math.Max(float64(maxLen), float64(right-left+1)))
	}
	writer.WriteString(strconv.Itoa(maxLen))
}
