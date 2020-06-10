// demo - Google Search 04
// Google Search 3.0

// Q: How do we avoid discarding results from slow servers?
// A: Replicate the servers. Send requests to multiple replicas, and use the first response.

// Reduce tail latency using replicated search servers.(使用复制的搜索服务器减少尾部延迟。)
package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	web1   = fakeSearch("web")
	web2   = fakeSearch("web")
	image1 = fakeSearch("image")
	image2 = fakeSearch("image")
	video1 = fakeSearch("video")
	video2 = fakeSearch("video")
)

// Result 定义返回值
type Result string

// Search 搜索服务
type Search func(query string) Result

// fakeSearch 虚拟的搜索服务
func fakeSearch(kind string) Search {
	return func(query string) Result {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		return Result(fmt.Sprintf("%s result for %q\n", kind, query))
	}
}

// First  query: 查询字串  replicas: 服务器的复制
func First(query string, replicas ...Search) Result {
	c := make(chan Result)
	searchReplica := func(i int) {
		c <- replicas[i](query)
	}
	for i := range replicas {
		go searchReplica(i)
	}
	return <-c
}

// Google 聚合的Google搜索服务
func Google(query string) (results []Result) {
	c := make(chan Result)
	go func() {
		c <- First(query, web1, web2)
	}()
	go func() {
		c <- First(query, image1, image2)
	}()
	go func() {
		c <- First(query, video1, video2)
	}()

	timeout := time.After(80 * time.Millisecond)
	for i := 0; i < 3; i++ {
		select {
		case result := <-c:
			results = append(results, result)
		case <-timeout:
			fmt.Println("time out")
			return
		}
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
