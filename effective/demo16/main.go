package main

import (
"fmt"
"unsafe"
)

type Post struct {
	published bool // 1 byte
	title     string // 16 Bytes
	content   string // 16 Bytes
	likes int16 // 2 bytes
	// 35 bytes total
}

type OptimisedPost struct {
	title     string // 16 Bytes
	content   string // 16 Bytes
	likes     int16  // 2 bytes
	published bool   // 1 byte
	// 35 bytes total
}

func main() {
	fmt.Printf("Sizeof Post: %d\n", unsafe.Sizeof(Post{}))
	fmt.Printf("Sizeof OptimisedPost: %d\n", unsafe.Sizeof(OptimisedPost{}))
}