package main

import (
	"fmt"
	"time"
)

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	go worker(1, jobs, results)
	go worker(2, jobs, results)
	go worker(3, jobs, results)
	for i := 0; i < 10; i++ {
		jobs <- i
	}

	close(jobs)

	for a := range results {
		fmt.Println(a)
	}

}

func worker(id int, jobs <-chan int, results chan<- int) {
	for a := range jobs {

		time.Sleep(time.Second * 2)
		results <- a
		fmt.Println("worker", id, "processed job", a)
	}
}
