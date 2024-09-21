package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type ShareData struct {
	mu   sync.Mutex
	data map[string]float64
}

func ScraperWebsite(url string, shared *ShareData, wg *sync.WaitGroup, r *rand.Rand) {
	defer wg.Done()

	time.Sleep(time.Duration(r.Intn(1000)) * time.Millisecond)

	scraperData := r.Float64() * 100

	// code to scrape website
	shared.mu.Lock()
	shared.data[url] = scraperData
	shared.mu.Unlock()

	fmt.Printf("Scraped data from %s:  %f\n", url, scraperData)
}

func main() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	sharedData := &ShareData{
		data: make(map[string]float64),
	}
	websites := []string{
		"https://finance.yahoo.com",
		"https://www.bloomberg.com",
		"https://www.investing.com",
		"https://www.alvantage.co",
		"https://www.google.com/finance",
		"https://www.nasdaq.com",
		"https://www.morningstar.com",
		"https://www.coinmarketcap.com",
		"https://www.data.worldbank.org",
		"https://www.quandl.com",
	}
	var wg sync.WaitGroup

	for _, url := range websites {
		wg.Add(1)
		go ScraperWebsite(url, sharedData, &wg, r)
	}

	wg.Wait()

	fmt.Println("Collected Financial Data!")
	//sharedData.mu.Lock()
	for site, value := range sharedData.data {
		fmt.Printf("%s: %f\n", site, value)
	}
	//sharedData.mu.Unlock()
}
