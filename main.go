package main

import (
	"flag"
	"fmt"
	"net/http"
	"sync"
	"time"
)

type Result struct {
	statusCode int
	duration   time.Duration
}

func worker(url string, requests int, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < requests; i++ {
		start := time.Now()
		resp, err := http.Get(url)
		duration := time.Since(start)

		if err != nil {
			results <- Result{statusCode: 0, duration: duration}
			continue
		}
		results <- Result{statusCode: resp.StatusCode, duration: duration}
		resp.Body.Close()
	}
}

func main() {
	url := flag.String("url", "", "URL do serviço a ser testado")
	requests := flag.Int("requests", 1, "Total de requests")
	concurrency := flag.Int("concurrency", 1, "Total de requests concorrentes")
	flag.Parse()

	if *url == "" {
		fmt.Println("Informe a URL")
		return
	}

	totalRequests := *requests
	numWorkers := *concurrency

	var wg sync.WaitGroup
	results := make(chan Result, totalRequests)

	fmt.Printf("Iniciando %d requests para %s com %d chamadas simultâneas\n", totalRequests, *url, numWorkers)
	start := time.Now()

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(*url, totalRequests/numWorkers, results, &wg)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	var totalDuration time.Duration
	var totalRequestsMade int
	var status200 int
	statusCodesCount := make(map[int]int)

	for result := range results {
		totalRequestsMade++
		totalDuration += result.duration
		if result.statusCode == 200 {
			status200++
		}
		if result.statusCode != 0 {
			statusCodesCount[result.statusCode]++
		}
	}

	duration := time.Since(start)

	fmt.Printf("Tempo total: %s\n", duration)
	fmt.Printf("Quantidade total de requests realizados: %d\n", totalRequestsMade)
	fmt.Printf("Quantidade de requests com status 200: %d\n", status200)
	fmt.Println("Distribuição de códigos de status HTTP:")
	for statusCode, count := range statusCodesCount {
		fmt.Printf("%d: %d\n", statusCode, count)
	}
}
