package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

var arr []time.Duration

func main() {

	http.HandleFunc("/add", addHandler)
	http.HandleFunc("/schedule", scheduleHandler)
	http.HandleFunc("/time", timeHandler)

	go handleTaskListener()

	fmt.Println("Server is listening...")
	http.ListenAndServe(":8181", nil)

}

func handleTaskListener() {
	for {
		if len(arr) > 0 {
			makeTask(arr[0])
			arr = removeIndex(arr, 0)
		}
	}
}

func addHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var addReq AddReq
	err := json.NewDecoder(r.Body).Decode(&addReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	duration, err := time.ParseDuration(addReq.TimeDuration)
	if err != nil {
		log.Fatal(err)
	}

	arr = append(arr, duration)
}

func scheduleHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	fmt.Fprint(w, arr)
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	seconds := 0.0
	for _, element := range arr {
		seconds += element.Seconds()
	}

	duration, _ := time.ParseDuration(strconv.Itoa(int(seconds)) + "s")

	data := TimeResp{}
	data.timeDuration = duration.String()

	fmt.Fprint(w, data)
}

type AddReq struct {
	TimeDuration string
	Sync         bool
}

type TimeResp struct {
	timeDuration string
}

func makeTask(duration time.Duration) {
	log.Println("started task at " + time.Now().GoString() + " duration = " + duration.String())
	time.Sleep(duration)
	log.Println("completed task at " + time.Now().GoString())
}

func removeIndex(s []time.Duration, index int) []time.Duration {
	return append(s[:index], s[index+1:]...)
}
