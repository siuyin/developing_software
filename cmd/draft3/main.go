package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	fmt.Println("draft3: refactor after writing tests")

	temp := make(chan int)
	sensor(temp)
	actuator(temp)

	select {}
}

func sensor(temp chan<- int) { // sensor only writes to the temp channel.
	go func() {
		for {

			sense(temp)

		}
	}()
}

func sense(temp chan<- int) { // sense only writes to the temp channel.
	for i := 0; i <= 100; i++ {
		temp <- i
		time.Sleep(90 * time.Millisecond)
	}
}

func actuator(temp <-chan int) { // actuator only reads from the temp channel.
	go func() {
		for {

			t := <-temp
			o := actuate(t)
			display(t, o) // display(<-temp, actuate(<-temp))

		}
	}()
}

const (
	lowTemp  = 25
	highTemp = 30
)

func actuate(t int) int {
	if t < lowTemp {
		return -1
	}
	if t > highTemp {
		return 1
	}
	return 0
}

func display(t, o int) {
	switch o {
	case -1:
		fmt.Printf("too cold: %d deg. C:  heating..\n", t)
	case 0:
		fmt.Printf("just right: %d deg. C: not doing anything.\n", t)
	case 1:
		fmt.Printf("too hot: %d deg. C:  cooling..\n", t)
	default:
		log.Fatalf("fatal: invalid output received by display: %d", o)
	}
}
