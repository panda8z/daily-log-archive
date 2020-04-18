
**第一遍：徒手写的，错误连篇**

```go
package main

import "fmt"

func addTwo(nums []int, target int) []int {
  var m map[int]int
  for i, v := range nums {
    if j,ok := m[target-v]; ok {
      return [i,j]
    }
    m[v] = i
  }
  return nil
}

func main() {
  fmt.Println(addTwo([]int{1,2,3,4,5,6}, 7))
  fmt.Println(addTwo([]int{1,2,3,4,5,6}, 9))
}
```

**正确的跑通的如下**

```go
package main

import "fmt"

func addTwo(nums []int, target int) []int {
	m := map[int]int{}
	for i, v := range nums {
		if j, ok := m[target-v]; ok {
			return []int{i, j}
		}
		m[v] = i
	}
	return nil
}

func main() {
	fmt.Println(addTwo([]int{1, 2, 3, 4, 5, 6}, 7))
	fmt.Println(addTwo([]int{1, 2, 3, 4, 5, 6}, 9))
}

```

**第二遍：徒手写**

```go

package main

import "fmt"

func addTwo(nums []int, target int) []int {
  m := map[int]int{}

  for i, v := range nums {
      if j, ok := map[target-v]; ok {
        return []int{i,j}
      }
      m[v] = i
  }
  return nil
}

func main() {
  fmt.Println([]int{1,2,3,4,5,6,7,8,9}, 10)
  fmt.Println([]int{1,2,3,4,5,6,7,8,9}, 7)
}

//这一遍
// # command-line-arguments
// 2020-04-17/demo02/main.go:9:29: syntax error: unexpected -, expecting ]

```