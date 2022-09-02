package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for {
		fmt.Println("Waiting worker", id)
		j, open := <-jobs
		if !open {
			break
		}
		fmt.Println("worker", id, "started  job", j)
		results <- j * 2
		fmt.Println("worker", id, "finished job", j)

	}
}

func main() {

	const numJobs = 10000
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j <= numJobs; j++ {
		time.Sleep(time.Second * 5)
		jobs <- j
	}
	close(jobs)
	for a := 1; a <= numJobs; a++ {
		fmt.Println(<-results)
	}
}
