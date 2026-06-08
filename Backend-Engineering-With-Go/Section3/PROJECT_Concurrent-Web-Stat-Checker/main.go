package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(5)
	var webURLS = []string{
		"https://google.com",
		"https://github.com",
		"https://fakesite-12345.com",
		"https://golang.org",
		"https://stackoverflow.com",
	}

	resultChan := make(chan string)
	for _, url := range webURLS {
		// if isWebsiteUp(url) {
		// 	fmt.Println(url, "is UP")
		// } else {
		// 	fmt.Println(url, "is DOWNer wg.Done()
		// }()")
		// }

		go func() {
			if isWebsiteUp(url) {
				resultChan <- url + " is UP."
			} else {
				resultChan <- url + " is DOWN."
			}
			defer wg.Done()
		}()

	}
	go func() {
		wg.Wait()
		close(resultChan)
	}()

	for msg := range resultChan {
		fmt.Println(msg)
	}

	// go func() {
	// 	if isWebsiteUp(webURLS[0]) {
	// 		fmt.Println(webURLS[0], "is UP")
	// 	} else {
	// 		fmt.Println(webURLS[0], "is DOWN")
	// 	}
	// 	defer wg.Done()
	// }()

	// go func() {
	// 	if isWebsiteUp(webURLS[1]) {
	// 		fmt.Println(webURLS[1], "is UP")
	// 	} else {
	// 		fmt.Println(webURLS[1], "is DOWN")
	// 	}
	// 	defer wg.Done()
	// }()

	// go func() {
	// 	if isWebsiteUp(webURLS[2]) {
	// 		fmt.Println(webURLS[2], "is UP")
	// 	} else {
	// 		fmt.Println(webURLS[2], "is DOWN")
	// 	}
	// 	defer wg.Done()
	// }()

	// go func() {
	// 	if isWebsiteUp(webURLS[3]) {
	// 		fmt.Println(webURLS[3], "is UP")
	// 	} else {
	// 		fmt.Println(webURLS[3], "is DOWN")
	// 	}
	// 	defer wg.Done()
	// }()

	// go func() {
	// 	if isWebsiteUp(webURLS[4]) {
	// 		fmt.Println(webURLS[4], "is UP")
	// 	} else {
	// 		fmt.Println(webURLS[4], "is DOWN")
	// 	}
	// 	defer wg.Done()
	// }()
}

func isWebsiteUp(url string) bool {
	response, err := http.Get(url)
	// fmt.Println("Response is:",response)
	if err != nil {
		return false
	} else {
		response.Body.Close()
		return true
	}
}
