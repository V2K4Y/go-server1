package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"sync"
	"time"
)

// Define a struct to represent the JSON payload sent to the server
type requestPayload struct {
	ToSort [][]int `json:"to_sort"`
}

// Define a struct to represent the JSON response sent by the server
type responsePayload struct {
	SortedArrays [][]int `json:"sorted_arrays"`
	TimeNs       int64   `json:"time_ns"`
}

func main() {
	// Register two HTTP endpoints with the server
	http.HandleFunc("/process-single", processSingle)
	http.HandleFunc("/process-concurrent", processConcurrent)

	// Specify the port
	port := ":8000"

	// Start the server and listen on the specified port
	fmt.Printf("Server is running on port %s\n", port)
	if err := http.ListenAndServe("0.0.0.0"+port, nil); err != nil {
		panic(err)
	}
}

// This function handles the /process-single HTTP endpoint
func processSingle(w http.ResponseWriter, r *http.Request) {
	// Check if the HTTP method is POST
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// Decode the JSON payload from the request body
	var payload requestPayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	// Check whether 'to_sort' data is present or not
	if len(payload.ToSort) == 0 {
		http.Error(w, "Error: 'to_sort' array empty / Not found!", http.StatusBadRequest)
		return
	}

	// Sort each sub-array sequentially
	sortedArrays := make([][]int, len(payload.ToSort))
	startTime := time.Now()
	for i, arr := range payload.ToSort {
		sorted := make([]int, len(arr))
		copy(sorted, arr)
		sort.Ints(sorted)
		sortedArrays[i] = sorted
	}
	timeTaken := time.Since(startTime).Nanoseconds()

	// Encode the sorted arrays and time taken as a JSON response
	response := responsePayload{
		SortedArrays: sortedArrays,
		TimeNs:       timeTaken,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// This function handles the /process-concurrent HTTP endpoint
func processConcurrent(w http.ResponseWriter, r *http.Request) {
	// Check if the HTTP method is POST
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// Decode the JSON payload from the request body
	var payload requestPayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	// Check whether 'to_sort' data is present or not
	if len(payload.ToSort) == 0 {
		http.Error(w, "Error: 'to_sort' array empty / Not found!", http.StatusBadRequest)
		return
	}

	// Sort each sub-array concurrently using goroutines and channels
	sortedChan := make(chan []int, len(payload.ToSort))
	var wg sync.WaitGroup
	startTime := time.Now()
	for _, arr := range payload.ToSort {
		wg.Add(1)
		go func(array []int) {
			defer wg.Done()
			sorted := make([]int, len(array))
			copy(sorted, array)
			sort.Ints(sorted)
			sortedChan <- sorted
		}(arr)
	}
	wg.Wait()
	close(sortedChan)
	sortedArrays := make([][]int, len(payload.ToSort))
	for i := range sortedArrays {
		sortedArrays[i] = <-sortedChan
	}
	timeTaken := time.Since(startTime).Nanoseconds()

	// Encode the sorted arrays and time taken as a JSON response
	response := responsePayload{
		SortedArrays: sortedArrays,
		TimeNs:       timeTaken,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
