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

	var n = 0
	for fileScanner.Scan() {
		dur := fileScanner.Text()
		n++
		if dur == "" {
			continue
		}
		go makeTask(dur, n)
	}

	fmt.Scanln() // для ожидания завершения работы горутин

	if err := fileScanner.Err(); err != nil {
		log.Fatalf("Error while reading file: %s", err)
	}
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
