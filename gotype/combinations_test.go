package gotype

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestFunc2( t *testing.T)  {
	rand.Seed(time.Now().Unix())

	for i:=0;i<16 ;i++  {
		fmt.Print(func2(), "")
	}

	fmt.Println()

	for i:=0;i<16 ;i++  {
		fmt.Print(func2(), "")
	}
}
