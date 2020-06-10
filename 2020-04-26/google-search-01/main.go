// demo - Google Search 01
// Google Search 1.0
// Synchronous
// The Google function takes a query and returns a slice of Results(which are just strings)
// Google invokes Web, Image and Video searches serially, appending them to the results slice.
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
	results = append(results, web(query))
	results = append(results, image(query))
	results = append(results, video(query))
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
