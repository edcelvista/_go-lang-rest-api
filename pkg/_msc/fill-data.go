package main

import (
	"bytes"
	"fmt"
	"net/http"
	"sync"
	"time"
)

var threads = 20
var batch = 100

func main() {
	var wg sync.WaitGroup
	client := &http.Client{}

	startAll := time.Now()
	var successCount int64
	var mu sync.Mutex // to safely count successes

	baseURL := "https://localhost:8443/crud"

	for i := 1; i <= threads; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			for j := 1; j <= batch; j++ {
				// Build URL with query params
				// params := url.Values{}
				// params.Add("user", fmt.Sprintf("user%d", id))
				// params.Add("id", fmt.Sprintf("%d", id))

				// POST body
				body := []byte(fmt.Sprintf(`{"message": { "field%v": "value%v" }}`, id, id))

				// Create request
				// req, err := http.NewRequest("POST", baseURL+"?"+params.Encode(), bytes.NewBuffer(body))
				req, err := http.NewRequest("POST", baseURL, bytes.NewBuffer(body))
				if err != nil {
					fmt.Println("Request error:", err)
					return
				}
				req.Header.Set("Content-Type", "application/json")

				// Measure time
				start := time.Now()
				resp, err := client.Do(req)
				elapsed := time.Since(start)

				if err == nil && resp.StatusCode == http.StatusOK {
					mu.Lock()
					successCount++
					mu.Unlock()
				}
				defer resp.Body.Close()

				fmt.Printf("Worker %d -> Status: %s | Time: %v\n", id, resp.Status, elapsed)
			}

		}(i)
	}

	wg.Wait()
	totalTime := time.Since(startAll).Seconds()
	totalRequests := threads * batch
	tps := float64(successCount) / totalTime

	fmt.Printf("\n--- Benchmark Result ---\n")
	fmt.Printf("Threads: %d\n", threads)
	fmt.Printf("Batch per thread: %d\n", batch)
	fmt.Printf("Total requests: %d\n", totalRequests)
	fmt.Printf("Success: %d\n", successCount)
	fmt.Printf("Elapsed time: %.2f sec\n", totalTime)
	fmt.Printf("TPS: %.2f\n", tps)
}
