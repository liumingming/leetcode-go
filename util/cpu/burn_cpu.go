package main

import (
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/process"
	"os"
	"runtime"
	"strconv"
	"time"
)

func main ()  {
	burnCpu()
	time.Sleep(10*time.Minute)
}

var cpuCount = 1
var cpuPercent = 60

func burnCpu() {

	runtime.GOMAXPROCS(cpuCount)

	var totalCpuPercent []float64
	var curProcess *process.Process
	var curCpuPercent float64
	var err error

	totalCpuPercent, err = cpu.Percent(time.Second, false)
	fmt.Println("totalcpu", totalCpuPercent)
	if err != nil {
		fmt.Println(err)
	}

	curProcess, err = process.NewProcess(int32(os.Getpid()))
	if err != nil {
		return
	}

	curCpuPercent, err = curProcess.CPUPercent()
	fmt.Println("curCpuPercent", curCpuPercent)

	if err != nil {
		return
	}

	otherCpuPercent := (100.0 - (totalCpuPercent[0] - curCpuPercent)) / 100.0

	fmt.Println("otherCpuPercent", otherCpuPercent)
	go func() {
		t := time.NewTicker(3 * time.Second)
		for {
			select {
			case <-t.C:
				totalCpuPercent, err = cpu.Percent(time.Second, false)
				if err != nil {

				}

				curCpuPercent, err = curProcess.CPUPercent()
				if err != nil {

				}
				otherCpuPercent = (100.0 - (totalCpuPercent[0] - curCpuPercent)) / 100.0
			}
		}
	}()
	for i := 0; i < cpuCount; i++ {
		go func() {
			busy := int64(0)
			idle := int64(0)
			all := int64(10000000)  //10豪秒
			ds := time.Duration(0)
			dx := (float64(cpuPercent) - totalCpuPercent[0]) / otherCpuPercent
			busy = busy + int64(dx*100000)
			idle = all - busy
			ds, _ = time.ParseDuration(strconv.FormatInt(idle, 10) + "ns")
			for i := 0; ; i = (i + 1) % 1000 {
				startTime := time.Now().UnixNano()
				for time.Now().UnixNano()-startTime < busy {
				}
				time.Sleep(ds)
				runtime.Gosched()
			}
		}()
	}
}
