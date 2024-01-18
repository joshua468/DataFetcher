package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

func main() {
	urls := []string{
		"https://www.google.com",
		"https://www.example.org",
		"https://www.example.net",
	}

	var wg sync.WaitGroup
	var mu sync.Mutex

	results := make(map[string]string)

	for _, url := range urls {
		wg.Add(1)
		go fetchData(url, &wg, &mu, results)
	}

	wg.Wait()

	fmt.Println("Aggregated Results:")
	for url, data := range results {
		fmt.Printf("URL: %s\nData: %s\n\n", url, data)
	}
}

func fetchData(url string, wg *sync.WaitGroup, mu *sync.Mutex, results map[string]string) {
	defer wg.Done()

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error fetching data from %s: %v\n", url, err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading data from %s: %v\n", url, err)
		return
	}

	mu.Lock()
	results[url] = string(body)
	mu.Unlock()

	fmt.Printf("Data fetched from %s\n", url)
}
