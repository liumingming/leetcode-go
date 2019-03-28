package effective

import (
	"fmt"
	"testing"
)

func TestGen(t *testing.T) {
	number := Generator()
	fmt.Println(<-number)
	fmt.Println(<-number)
	fmt.Println(<-number)
	fmt.Println(<-number)
	fmt.Println(<-number)
	fmt.Println(<-number)


	close(number)
}

func TestClose(t *testing.T)  {
	ch := make(chan string)
	go func() {
		ch <- "Hello!"
		close(ch)
	}()

	fmt.Println(<-ch) // Print "Hello!".
	fmt.Println("wo bei guanbi le ",<-ch) // Print the zero value "" without blocking.
	fmt.Println(<-ch) // Once again print "".
	//v, ok := <-ch     // v is "", ok is false.
	//
	//// Receive values from ch until closed.
	//for v := range ch {
	//	fmt.Println(v) // Will not be executed.
	//}
}

