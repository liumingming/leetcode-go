package main

import (
	"fmt"
)

func main()  {
	//Declaration and initialization
	var m map[string]int = make(map[string]int)

	//update
	m["liumingming"] = 100
	m["liuxingxing"] = 90


	//get
	score, ok := m["liumingming"]
	if ok {
		fmt.Println(score)
	}

	score2, ok := m["liuyuanyuan"]
	fmt.Println(ok)
	if !ok {
		fmt.Println(score2)
	}

	//traverse
	for key, value := range m{
		fmt.Println(key, value)
	}

	//delete
	delete(m, "liumingming")
	fmt.Println(m)




}



