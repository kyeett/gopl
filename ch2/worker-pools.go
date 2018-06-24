package main

import (
	"math/rand"
	"time"
)

var nJobs int = 1000000
var nWorkers int = 5000
var sleepFactor int = 2

func worker(id int, jobs chan int, results chan int) {

	for j := range jobs {
		//fmt.Printf("worker %d started job %d\n", id, j)
		time.Sleep(time.Duration(rand.Intn(sleepFactor)) * time.Millisecond)

		//fmt.Printf("worker %d completed job %d\n", id, j)
		results <- j * j
	}

}

func main() {
	jobs := make(chan int, nJobs)
	results := make(chan int, nJobs)

	for w := 1; w <= nWorkers; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j <= nJobs; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= nJobs; a++ {
		<-results
	}

}
