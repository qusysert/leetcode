package main

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"
)

func main() {
	const (
		totalRequests = 10000
		duration      = 10 * time.Second
		url           = "https://10.8.141.108:9993/api/v1/oauth2/token"
	)

	// Calculate requests per second
	requestsPerSecond := float64(totalRequests) / duration.Seconds()
	interval := time.Duration(float64(time.Second) / requestsPerSecond)

	fmt.Printf("Sending %d requests over %v (%.2f req/sec)\n", totalRequests, duration, requestsPerSecond)

	var wg sync.WaitGroup
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	// Create HTTP client with timeout and skip certificate verification
	client := &http.Client{
		Timeout: 5 * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true, // This fixes the certificate error
			},
		},
	}

	// Stats tracking
	var successCount, errorCount int64
	var mu sync.Mutex

	start := time.Now()

	for i := 0; i < totalRequests; i++ {
		<-ticker.C // Wait for next tick

		wg.Add(1)
		go func(reqNum int) {
			defer wg.Done()

			// Create request
			req, err := http.NewRequest("POST", url, strings.NewReader(""))
			if err != nil {
				mu.Lock()
				errorCount++
				mu.Unlock()
				return
			}

			// Set headers
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			req.Header.Set("Accept", "application/json")
			req.Header.Set("X-Forwarded-For", "203.0.113.5")

			// Send request
			resp, err := client.Do(req)
			if err != nil {
				mu.Lock()
				errorCount++
				mu.Unlock()
				fmt.Printf("Request %d failed: %v\n", reqNum, err)
				return
			}
			defer resp.Body.Close()

			mu.Lock()
			if resp.StatusCode >= 200 && resp.StatusCode < 300 {
				successCount++
			} else {
				errorCount++
			}
			mu.Unlock()

			if reqNum%1000 == 0 {
				fmt.Printf("Completed %d requests\n", reqNum)
			}
		}(i + 1)
	}

	// Wait for all requests to complete
	wg.Wait()
	elapsed := time.Since(start)

	fmt.Printf("\nCompleted in %v\n", elapsed)
	fmt.Printf("Success: %d, Errors: %d\n", successCount, errorCount)
	fmt.Printf("Actual rate: %.2f req/sec\n", float64(totalRequests)/elapsed.Seconds())
}
