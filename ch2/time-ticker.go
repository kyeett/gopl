package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {
	tick := time.Tick(1000 * time.Millisecond)
	abort := make(chan struct{})

	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()

	clearScreen()
	fmt.Println("Starting countdown")
	for countdown := 10; countdown > 0; countdown-- {

		select {
		case <-abort:
			clearScreen()
			fmt.Println("\nAborting")
			return
		case <-tick:
			clearScreen()
			fmt.Println("\nCountdown: ", countdown)
		}

	}
	launch()
}

func clearScreen() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}

func launch() {
	fmt.Println("Launching space shuttle")
}
