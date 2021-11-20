package main

import "fmt"

func main() {
	fmt.Println("first draft: maintain comfortable temperature")

	for i := 0; i <= 100; i++ {
		if i < 25 {
			fmt.Printf("too cold: %d deg. C:  heating..\n", i)
			continue
		}
		if i > 30 {
			fmt.Printf("too hot: %d deg. C:  cooling..\n", i)
			continue
		}
		fmt.Printf("just right: %d deg. C: not doing anything.\n", i)
	}
}
