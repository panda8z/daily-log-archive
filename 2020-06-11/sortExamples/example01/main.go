package main

// sort包的示例程序
// 展示了 实现一个 sort。interface接口就能 便捷使用 实现排序
import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%s: %d", p.Name, p.Age)
}

// ByAge implements sort.Interface for []Person based on the Age field
type ByAge []Person

func (b ByAge) Len() int {
	return len(b)
}

func (b ByAge) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}
func (b ByAge) Less(i, j int) bool {
	return b[i].Age < b[j].Age
}

func main() {
	people := []Person{
		{"David", 13},
		{"Peak", 20},
		{"Panda", 18},
		{"Bob", 11},
	}
	// print original
	fmt.Println(people)

	//  print sorted []people
	sort.Sort(ByAge(people))
	fmt.Println(people)

	// print z-a sorted []people
	sort.Slice(people, func(i, j int) bool {
		return people[i].Age > people[j].Age
	})
	fmt.Println(people)
}
