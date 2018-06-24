package main

import (
	"fmt"
	"os"
	"os/exec"
	"sync"
	"time"
)

var sema = make(chan struct{}, 30)

func DoSomeWork() {
	sema <- struct{}{}        //Aquire token
	defer func() { <-sema }() //Release token
	//fmt.Println("Work started")
	time.Sleep(1000 * time.Millisecond)
	//fmt.Println("Work done")
}

func clearScreen() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}

func main() {
	var wg sync.WaitGroup
	jobsStarted := 0
	jobsDone := 0
	tick := time.NewTicker(20 * time.Millisecond)

	var printTick <-chan time.Time
	printTick = time.Tick(50 * time.Millisecond)

	go func() {
		for ; true; <-printTick {
			clearScreen()
			inQueue := len(sema)
			fmt.Println(jobsStarted-inQueue-jobsDone, inQueue, jobsDone)
		}
	}()

	for i := 0; i < 100; i++ {
		<-tick.C
		wg.Add(1)
		jobsStarted++
		go func() {
			defer wg.Done()
			DoSomeWork()
			jobsDone++
		}()
	}
	fmt.Println("All work added")

	wg.Wait()
	inQueue := len(sema)
	fmt.Println(jobsStarted-inQueue, inQueue, jobsDone)
}
