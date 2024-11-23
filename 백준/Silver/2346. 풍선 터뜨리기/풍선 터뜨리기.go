package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type deque struct {
	queue    [][]int
	size     int
	front    int
	rear     int
	capacity int
}

func newDeque(capacity int) *deque {
	return &deque{
		queue:    make([][]int, capacity),
		size:     0,
		front:    0,
		rear:     0,
		capacity: capacity,
	}
}

func (d *deque) addFirst(value int, index int) {
	if d.size == 0 {
		d.queue[0][0] = value
		d.queue[0][1] = index
		d.rear = 0
		d.front = 0
	} else {
		d.front = d.inc(d.front)
		d.queue[d.front][0] = value
		d.queue[d.front][1] = index
	}
	d.size++
}

func (d *deque) addLast(value int, index int) {
	if d.size == 0 {
		d.queue[0][0] = value
		d.queue[0][1] = index
		d.rear = 0
		d.front = 0
	} else {
		d.rear = d.dec(d.rear)
		d.queue[d.rear][0] = value
		d.queue[d.rear][1] = index
	}
	d.size++
}

func (d *deque) pollFirst() []int {
	_front := d.front
	d.front = d.dec(_front)
	d.size--
	return d.queue[_front]
}

func (d *deque) pollLast() []int {
	_rear := d.rear
	d.rear = d.inc(_rear)
	d.size--
	return d.queue[_rear]
}

func (d *deque) inc(pointer int) int {
	if pointer == d.capacity-1 {
		return 0
	}
	return pointer + 1
}

func (d *deque) dec(pointer int) int {
	if pointer == 0 {
		return d.capacity - 1
	}
	return pointer - 1
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	str, _ := reader.ReadString('\n')
	capacity, _ := strconv.Atoi(strings.TrimSpace(str))

	str, _ = reader.ReadString('\n')
	data := strings.Fields(str)

	d := newDeque(capacity)

	for i, _ := range data {
		d.queue[i] = make([]int, 2)
	}

	for i, value := range data {
		num, _ := strconv.Atoi(value)
		d.addLast(num, i+1)
	}
	for i := 0; i < capacity-1; i++ {
		displacement := d.pollFirst()
		distance := displacement[0]

		writer.WriteString(strconv.Itoa(displacement[1]) + " ")
		if distance > 0 {
			for j := 0; j < distance-1; j++ {
				temp := d.pollFirst()
				d.addLast(temp[0], temp[1])
			}
		} else {
			for j := 0; j < -distance; j++ {
				temp := d.pollLast()
				d.addFirst(temp[0], temp[1])
			}
		}
	}
	writer.WriteString(strconv.Itoa(d.pollFirst()[1]))
}
