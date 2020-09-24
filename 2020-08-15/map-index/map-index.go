/*
基于hash 的多键索引
*/

package main

import "fmt"

type Profile struct {
	Name    string
	Age     int
	Married bool
}

func main() {
	list := []*Profile{
		{Name: "P1", Age: 30, Married: true},
		{Name: "P2", Age: 34},
		{Name: "P3", Age: 23},
	}
	// 基于hash的多键索引
	// buildIndexByHash(list)
	// queryDataHash("P1", 30)

	// 基于 map 特性的多键索引 start
	buildIndexByMap(list)
	queryDataMap("P1", 30)

}

/*
基于hash的多键索引 start
*/

func simpleHash(str string) (ret int) {
	for i := 0; i < len(str); i++ {
		c := str[i]
		ret += int(c)
	}
	return
}

type classicQueryKey struct {
	Name string
	Age  int
}

func (c *classicQueryKey) hash() int {
	return simpleHash(c.Name) + c.Age*1000000
}

var mapper = make(map[int][]*Profile)

func buildIndexByHash(list []*Profile) {
	for _, profile := range list {
		/*
		 */
		key := classicQueryKey{profile.Name, profile.Age}
		exitValue := mapper[key.hash()]        // take out the value of the hash
		exitValue = append(exitValue, profile) // add new profile
		mapper[key.hash()] = exitValue         // reset the hash-key`s value
	}
}

func queryDataHash(name string, age int) {
	keyToQuery := classicQueryKey{name, age}
	resultList := mapper[keyToQuery.hash()]
	for _, result := range resultList {
		if result.Age == age && result.Name == name {
			fmt.Println(result)
			return
		}
	}
	fmt.Println("not found!")
}

/*
基于hash的多键索引 end
*/

/*
基于 map 特性的多键索引 start
*/
type queryKey struct {
	Name string
	Age  int
}

var mapperMap = make(map[queryKey]*Profile)

func buildIndexByMap(list []*Profile) {
	for _, profile := range list {
		key := queryKey{
			Name: profile.Name,
			Age:  profile.Age,
		}
		mapperMap[key] = profile
	}
}

func queryDataMap(name string, age int) {
	key := queryKey{name, age}
	result, ok := mapperMap[key]
	if ok {
		fmt.Println(result)
	} else {
		fmt.Println("not found!")
	}
}

/*
基于 map 特性的多键索引 end
*/
