package main

import (
	"fmt"
	"sort"
)

type Student struct {
	Name string
	Weight  int
	Height int
}

func main()  {
	var students = make([]Student,0)
	tom := Student{
		Name:   "tom",
		Weight: 100,
		Height: 178,
	}
	students = append(students, tom)

	bob := Student{
		Name:   "bob",
		Weight: 86,
		Height: 180,
	}
	students = append(students, bob)

	jack := Student{
		Name:   "tom",
		Weight: 121,
		Height: 176,
	}
	students = append(students, jack)

	sort.Slice(students, func(i, j int) bool {
		return students[i].Weight < students[j].Weight
	})

	fmt.Println("0", students)
	sort.Slice(students, func(i, j int) bool {
		return students[i].Weight > students[j].Weight
	})
	fmt.Println("1", students)

}
