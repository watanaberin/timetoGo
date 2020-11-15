package main

import (
	"fmt"

	"rsc.io/quote"
)

var (
	ch1 = make(chan string)
	ch2 = make(chan string)
)

func download(url string) {
	var sum int = 0
	for i := 0; i < 5; i++ {
		sum += i
	}
	ch1 <- "Downloaded" + url + string(sum)
}
func main() {
	fmt.Println(quote.Hello())
	go download("/a")
	go func() {//需要在线程里完成 否则会报fatal error: all goroutines are asleep - deadlock!
		v := <-ch1
		ch2 <- v
	}()
	fmt.Println(<-ch2)
}
