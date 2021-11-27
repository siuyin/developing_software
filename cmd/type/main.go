package main

import "fmt"

type Temperature int

const (
	lowTemp  Temperature = 25
	highTemp             = Temperature(30)
)

func pTemp(t Temperature) {
	fmt.Println(t)
}

func main() {
	fmt.Println("type safety")
	t := 27
	pTemp(t)
}
