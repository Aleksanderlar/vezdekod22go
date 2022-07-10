package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	file, err := os.Open("durations.txt")
	if err != nil {
		log.Fatalf("Error when opening file: %s", err)
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)

	nCh := make(chan int)
	durCh := make(chan string)

	var freeProcessors int
	fmt.Println("Enter count processors")
	_, err = fmt.Scanf("%d", &freeProcessors)
	if err != nil {
		log.Fatalf("Error when parsing processor count: %s", err)
	}

	for i := 0; i < freeProcessors; i++ {
		go func() {
			for {
				select {
				case v, open := <-nCh:
					if !open {
						return
					}
					makeTask(<-durCh, v)
				}
			}
		}()
	}

	var n = 0
	for fileScanner.Scan() {
		dur := fileScanner.Text()
		n++
		if dur == "" {
			continue
		}
		nCh <- n
		durCh <- dur
	}

	close(nCh)
	close(durCh)
}

func makeTask(dur string, n int) {
	duration, err := time.ParseDuration(dur)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("started task №" + strconv.Itoa(n) + " at " + time.Now().GoString() + " duration = " + dur)
	time.Sleep(duration)
	log.Println("completed task №" + strconv.Itoa(n) + " at " + time.Now().GoString())
}
