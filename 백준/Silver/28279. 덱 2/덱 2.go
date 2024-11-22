package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Deque struct {
	deque    []int
	size     int
	head     int
	tail     int
	capacity int
}

func NewDeque(capacity int) *Deque {
	return &Deque{
		deque:    make([]int, capacity),
		size:     0,
		head:     0,
		tail:     0,
		capacity: capacity,
	}
}

func (d *Deque) addFirst(value int) {
	if d.size == 0 {
		d.deque[0] = value
		d.head = 0
		d.tail = 0
	} else {
		d.head = d.decrease(d.head)
		d.deque[d.head] = value
	}
	d.size++
}

func (d *Deque) addLast(value int) {
	if d.size == 0 {
		d.deque[0] = value
		d.head = 0
		d.tail = 0
	} else {
		d.tail = d.increase(d.tail)
		d.deque[d.tail] = value
	}
	d.size++
}

func (d *Deque) peekFirst() int {
	if d.size > 0 {
		return d.deque[d.head]
	}
	return -1
}

func (d *Deque) peekLast() int {
	if d.size > 0 {
		return d.deque[d.tail]
	}
	return -1
}

func (d *Deque) pollFirst() int {
	if d.size > 0 {
		value := d.deque[d.head]
		d.head = d.increase(d.head)
		d.size--
		return value
	}
	return -1
}

func (d *Deque) pollLast() int {
	if d.size > 0 {
		value := d.deque[d.tail]
		d.tail = d.decrease(d.tail)
		d.size--
		return value
	}
	return -1
}

func (d *Deque) getSize() int {
	return d.size
}

func (d *Deque) increase(pointer int) int {
	if pointer == d.capacity-1 {
		return 0
	}
	return pointer + 1
}

func (d *Deque) decrease(pointer int) int {
	if pointer == 0 {
		return d.capacity - 1
	}
	return pointer - 1
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	line, _ := reader.ReadString('\n')
	count, _ := strconv.Atoi(strings.TrimSpace(line))

	deque := NewDeque(count * 2)

	for i := 0; i < count; i++ {
		line, _ := reader.ReadString('\n')
		tokens := strings.Fields(line)
		command, _ := strconv.Atoi(tokens[0])

		switch command {
		case 1:
			value, _ := strconv.Atoi(tokens[1])
			deque.addFirst(value)
		case 2:
			value, _ := strconv.Atoi(tokens[1])
			deque.addLast(value)
		case 3:
			fmt.Fprintln(writer, deque.pollFirst())
		case 4:
			fmt.Fprintln(writer, deque.pollLast())
		case 5:
			fmt.Fprintln(writer, deque.getSize())
		case 6:
			if deque.getSize() == 0 {
				fmt.Fprintln(writer, 1)
			} else {
				fmt.Fprintln(writer, 0)
			}
		case 7:
			fmt.Fprintln(writer, deque.peekFirst())
		case 8:
			fmt.Fprintln(writer, deque.peekLast())
		}
	}
}
