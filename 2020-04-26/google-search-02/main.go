// demo - Google Search 02
// Google Search 2.0
// Run the Web, Image and Video searches concurrently, and wait for all results.
// No locks, No condition variables, No callbacks
package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	web   = fakeSearch("web")
	image = fakeSearch("image")
	video = fakeSearch("video")
)

// Result type
type Result string

// Search fake
type Search func(query string) Result

func fakeSearch(kind string) Search {
	return func(query string) Result {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		return Result(fmt.Sprintf("%s result for %q\n", kind, query))
	}
}

// Google ,
func Google(query string) (results []Result) {
	c := make(chan Result)
	go func() {
		c <- web(query)
	}()
	go func() {
		c <- image(query)
	}()
	go func() {
		c <- video(query)
	}()

	for i := 0; i < 3; i++ {
		result := <-c
		results = append(results, result)
	}
	return results
}

func main() {
	rand.Seed(time.Now().UnixNano())
	start := time.Now()
	result := Google("google")
	elapsed := time.Since(start)
	fmt.Println(result)
	fmt.Println(elapsed)

}
