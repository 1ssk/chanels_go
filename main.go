package main

import (
	"fmt"
)

func Print(s chan<- string) {
	for {

		s <- fmt.Sprintf("How")
		s <- fmt.Sprintf("When")
	}

}
func main() {
	s := make(chan string)
	go Print(s)
	for i := 0; i < 1; i++ {
		fmt.Println("Hi")
		fmt.Println(<-s, ",", <-s)
	}
}
