package main

import (
	"runtime"
	"time"
)

func main()  {
	var cpuCount = runtime.NumCPU()
	var count int64
	for i:=0; i<cpuCount;i++ {
		go func() {
			for {
				count++
				count--
			}
		}()
	}
	time.Sleep(10*time.Minute)
}
