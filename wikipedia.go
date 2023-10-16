package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func fetchWiki(city string) {
	url := fmt.Sprintf("https://en.wikipedia.org/wiki/%s", city)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error fetching page for city %s: %s\n", city, err)
	}

	defer resp.Body.Close()
}

func fetchWikiAsync(city string, ch chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()

	url := fmt.Sprintf("https://en.wikipedia.org/wiki/%s", city)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error fetching page for city %s: %s\n", city, err)
	}
	defer resp.Body.Close()

	ch <- fmt.Sprintf("City: %s", city)
}

func main() {
	// Download http urls
	// 1. sequentially
	// 2. concurrently
	// Compare the time

	start := time.Now()

	cities := []string{"Detroit", "Seoul", "Paris", "Manila"}

	ch := make(chan string)
	var wg sync.WaitGroup

	for _, city := range cities {
		// fetchWiki(city)

		// WHY?: can we call a function on an object that
		//       was declared, but not defined
		wg.Add(1)
		go fetchWikiAsync(city, ch, &wg)
	}

	// WHY?: do we call the the waitgroup wait in its own thread
	go func() {
		wg.Wait()
		close(ch)
	}()

	for result := range ch {
		fmt.Println(result)
	}

	fmt.Println("Time: ", time.Since(start))
}
