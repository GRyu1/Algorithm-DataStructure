package main

import (
	"bufio"
	"container/list"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	temp, _ := reader.ReadString('\n')
	temp, _ = reader.ReadString('\n')
	qs := strings.Fields(temp)

	temp, _ = reader.ReadString('\n')
	data := strings.Fields(temp)
	l := list.New()
	for i, q := range qs {
		isQueue, _ := strconv.Atoi(q)
		if isQueue == 0 {
			datum, _ := strconv.Atoi(data[i])
			l.PushBack(datum)
		}
	}

	temp, _ = reader.ReadString('\n')
	m, _ := strconv.Atoi(strings.TrimSpace(temp))
	temp, _ = reader.ReadString('\n')
	inputs := strings.Fields(temp)
	for i := 0; i < m; i++ {
		input, _ := strconv.Atoi(inputs[i])
		l.PushFront(input)
		value := l.Back().Value.(int)
		_, _ = writer.WriteString(strconv.Itoa(value) + " ")
		l.Remove(l.Back())
	}

}
