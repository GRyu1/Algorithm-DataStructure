package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	t := time.Now()

	var yyyy, mm, dd string
	y, m, d := t.Date()
	yyyy = dateformat(y)
	mm = dateformat(int(m))
	dd = dateformat(d)

	fmt.Printf("%s-%s-%s", yyyy, mm, dd)
}

func dateformat(i int) string {
	var o string
	if i >= 10 {
		o = strconv.Itoa(i)
	} else {
		o = "0" + strconv.Itoa(i)
	}
	return o
}
