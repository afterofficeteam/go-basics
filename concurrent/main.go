package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

// Global variable to package main
// Exported rules also apply
// var URLs = []string{
// 	"https://jsonplaceholder.typicode.com/posts?userId=1",    // 10
// 	"https://jsonplaceholder.typicode.com/comments?postId=1", // 5
// 	"https://jsonplaceholder.typicode.com/albums?userId=1",   // 10
// }

var URLs = []string{
	"https://jsonplaceholder.typicode.com/posts",    // 100
	"https://jsonplaceholder.typicode.com/comments", // 500
	"https://jsonplaceholder.typicode.com/albums",   // 100
}

func main() {
	defer fmt.Println("main thread end")
	defer func() {
		fmt.Println("defer")
	}()
	fmt.Println("main thread start")

	// basicGoroutine()
	// basicChan()

	// fetchURL(URLs[0])

	// fetchSequentially()
	// fmt.Println("")
	// fetchConcurrently()
}

func basicGoroutine() {
	// spawn 1 goroutine with "anonymous func" / Closure
	go func() {
		fmt.Println("anonymous goroutine")
	}()

	// spawn 1 goroutine that run fmt.Println()
	go fmt.Println("hello")

	// spawn 10 goroutine that each goroutine will run fmt.Println() 0 to 9 respectively
	for i := 0; i < 10; i++ {
		go fmt.Println(i)
	}
	time.Sleep(time.Second * 1) // blocking 1 sec to make sure other goroutines running
}

func basicChan() {
	/* UNBUFFERED CHANNEL */
	messageChan := make(chan string) // unbuffered channel
	go func() {
		messageChan <- "tes" // pass "tes" to messageChan
	}()
	msg := <-messageChan // output 1 data from messageChan // BLOCKING
	fmt.Println(msg)
	// msg = <-messageChan // Blocking main goroutine, waiting for new data from messageChan

	/* BUFFERED CHANNEL */
	bufferedMessageChan := make(chan string, 2)

	go func() {
		for i := 1; i <= 10; i++ {
			bufferedMessageChan <- fmt.Sprintf("data %d", i)
			fmt.Println("done sending data", i)
			time.Sleep(time.Second)
		}
		close(bufferedMessageChan)
	}()

	// loop over channel. quit loop when channel is CLOSED
	for str := range bufferedMessageChan {
		fmt.Println("received", str)
		// time.Sleep(time.Second)
	}

	/* CHANNEL DIRECTION in func */
	chPertama := make(chan int, 1)
	chKedua := make(chan string, 1)
	passMsg1(chPertama, 2)
	passMsg2(chPertama, chKedua)

	fmt.Println("chKedua:", <-chKedua)
}

func passMsg1(ch chan<- int, data int) {
	ch <- data
}

func passMsg2(ch1 <-chan int, ch2 chan<- string) {
	angka := <-ch1
	texts := []string{"nol", "satu", "dua", "tiga", "empat"}
	ch2 <- texts[angka]
}

// fetchURL launch HTTP call GET to url and return response
func fetchURL(url string) (string, error) {
	// Start recording the time when this function is executed
	start := time.Now()

	// Send an HTTP GET request to the given URL
	resp, err := http.Get(url)
	if err != nil {
		// If there is an error while sending the request, return the error
		return "", fmt.Errorf("error fetching %s: %v", url, err)
	}
	// Close resp.Body after the function is finished executing
	defer resp.Body.Close()

	// Read the entire response body from the server
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		// If there is an error while reading the response body, return the error
		return "", fmt.Errorf("error reading response body from %s: %v", url, err)
	}

	// Calculate the duration it took to fetch the data from the URL
	duration := time.Since(start).Seconds()

	// Return the number of bytes fetched, the URL, and the duration it took
	fmt.Printf("Fetched %d bytes from %s in %.2f seconds\n", len(body), url, duration)
	return string(body), nil
}

func fetchSequentially() {
	fmt.Println("fetchSequentially")
	start := time.Now()
	// fmt.Println(start)

	var results []string

	// Iterate over each URL and fetch them one by one
	for _, url := range URLs {
		result, err := fetchURL(url)
		if err != nil {
			// Append the error message to results if an error occurs
			results = append(results, err.Error())
		} else {
			// Append the result to results if successful
			results = append(results, result)
		}
	}
	// fmt.Println("results", results[0]) // posts
	// fmt.Println("results", results[1]) // comments
	// fmt.Println("results", results[2]) // albums

	duration := time.Since(start).Seconds()
	fmt.Println("diff:", duration, "second")
}

func fetchConcurrently() {
	fmt.Println("fetchConcurrently")
	start := time.Now()
	// fmt.Println(start)

	// Create a channel to communicate results from goroutines
	ch := make(chan string)

	// Iterate over each URL
	for _, url := range URLs {
		// Launch a goroutine to fetch each URL concurrently
		go func(url string) {
			result, err := fetchURL(url)
			if err != nil {
				// Send the error message to the channel if an error occurs
				ch <- err.Error()
			} else {
				// Send the result to the channel if successful
				ch <- result
			}
		}(url)
	}

	var results []string
	// Collect results from the channel for each URL
	for range URLs { // iterate len(URLs) times without assign any new variable.
		results = append(results, <-ch)
	}

	// do something with results

	duration := time.Since(start).Seconds()
	fmt.Println("diff:", duration, "second")
}

func fetchConcurrentlyWaitgroup() {
	fmt.Println("fetchConcurrentlyWaitgroup")
	start := time.Now()
	// fmt.Println(start)

	// Create a channel to communicate results from goroutines
	ch := make(chan string)
	wg := sync.WaitGroup{}

	// Iterate over each URL
	for _, url := range URLs {
		// Launch a goroutine to fetch each URL concurrently
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			result, err := fetchURL(url)
			if err != nil {
				// Send the error message to the channel if an error occurs
				ch <- err.Error()
			} else {
				// Send the result to the channel if successful
				ch <- result
			}
		}(url)
	}

	var results []string
	// Collect results from the channel for each URL
	wg.Wait()

	for msg := range ch { // iterate len(URLs) times without assign any new variable.
		results = append(results, msg)
	}

	// do something with results

	duration := time.Since(start).Seconds()
	fmt.Println("diff:", duration, "second")
}
