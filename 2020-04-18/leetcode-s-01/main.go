// 1. 两数之和

package main

import "fmt"

func addTwo(nums []int, target int) []int {
	//1. 创建map来存 【值，索引】
	m := map[int]int{}
	// 2. 遍历每一个数
	for i, v := range nums {
		// 3. 根据map的特性是否能找到map里有相应的值
		if j, ok := m[target-v]; ok {
			// 如果存在就 返回存储的值的索引 j和当前的索引 i组成的切片。
			return []int{j, i}
		}
		// 4. 如果没有就 按[值，索引]存进map
		m[v] = i
	}
	return nil
}
func main() {
	fmt.Println(addTwo([]int{1, 2, 5, 4}, 5))
	fmt.Println(addTwo([]int{1, 2, 5, 4}, 7))
}