package main

import (
	"fmt"
)

func main()  {
	//Declaration and initialization
	var m  = make(map[string]int)

	//update
	m["liu_ming_ming"] = 100
	m["liu_xing_xing"] = 90


	//get
	score, ok := m["liu_ming_ming"]
	if ok {
		fmt.Println(score)
	}

	score2, ok := m["liu_yuan_yuan"]
	fmt.Println(ok)
	if !ok {
		fmt.Println(score2)
	}

	//traverse
	for key, value := range m{
		fmt.Println(key, value)
	}

	//delete
	delete(m, "liu_ming_ming")
	fmt.Println(m)
}



