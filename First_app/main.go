package main

import (
	"fmt"
	"time"
)

func slow_draw(total int, SYMBOL string, channel chan bool) {
	for i := 0; i < total; i++ {
		var row_string string
		for n := 0; n < total-i; n++ {
			row_string += " "
		}
		for j := 0; j < i*2+1; j++ {
			row_string += SYMBOL
		}
		fmt.Println(row_string)
		time.Sleep(50 * time.Millisecond)
	}

	channel <- true
}

func main() {
	row_channel := make(chan bool)
	const SYMBOL string = "*"

	go slow_draw(118, SYMBOL, row_channel)

	<-row_channel
}
