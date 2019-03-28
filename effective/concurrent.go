package effective

import "fmt"

//how to kill a goroutine
func Generator() chan int {
	ch := make(chan int)
	go func() {
		n := 1
		for {
			select {
			case ch <- n:
				n++
			case c := <-ch:
				// the c is zero value
				fmt.Println("channel is close ", c)
				return
			}
		}
	}()
	return ch
}
