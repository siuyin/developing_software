package main

import "fmt"

//go:generate stringer -type=Size
type Size int

const (
	small Size = iota
	medium
	large
)

func main() {
	fmt.Println("enums")
	sz := medium
	fmt.Printf("My size is %s\n", sz)
}
