package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("draft2: sensor and actuator")

	temp := make(chan int)
	sense(temp)
	actuate(temp)

	select {}
}

func sense(temp chan<- int) { // sense only writes to the temp channel.
	go func() {
		for {

			for i := 0; i <= 100; i++ {
				temp <- i
				time.Sleep(90 * time.Millisecond)
			}

		}
	}()
}

func actuate(temp <-chan int) { // actuate only reads from the temp channel.
	go func() {
		for {

			t := <-temp
			if t < 25 {
				fmt.Printf("too cold: %d deg. C:  heating..\n", t)
				continue
			}
			if t > 30 {
				fmt.Printf("too hot: %d deg. C:  cooling..\n", t)
				continue
			}
			fmt.Printf("just right: %d deg. C: not doing anything.\n", t)

		}
	}()
}
