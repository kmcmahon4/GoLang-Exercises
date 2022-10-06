package exercises

import (
	"fmt"
	"time"
)

/*
Create an ongoing process that will exit after 2 seconds.
*/
func Exercise5() {

	quit := make(chan int)

	go func() {
		spawnProcess(quit)
	}()
	<-quit

	fmt.Println("Here!")

}

func spawnProcess(quit chan int) {
	for {
		select {
		case <-quit:
			print("hi")
		case <-time.After(2 * time.Second):
			fmt.Println("leaving")
			quit <- 1
			return
		}

	}
}
