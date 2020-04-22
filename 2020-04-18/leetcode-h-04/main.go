// [4. 寻找两个有序数组的中位数 - 力扣（LeetCode）](https://leetcode-cn.com/problems/median-of-two-sorted-arrays/)

// class Solution {
//     public double findMedianSortedArrays(int[] A, int[] B) {
//         int m = A.length;
//         int n = B.length;
//         if (m > n) {
//             return findMedianSortedArrays(B,A); // 保证 m <= n
//         }
//         int iMin = 0, iMax = m;
//         while (iMin <= iMax) {
//             int i = (iMin + iMax) / 2;
//             int j = (m + n + 1) / 2 - i;
//             if (j != 0 && i != m && B[j-1] > A[i]){ // i 需要增大
//                 iMin = i + 1;
//             }
//             else if (i != 0 && j != n && A[i-1] > B[j]) { // i 需要减小
//                 iMax = i - 1;
//             }
//             else { // 达到要求，并且将边界条件列出来单独考虑
//                 int maxLeft = 0;
//                 if (i == 0) { maxLeft = B[j-1]; }
//                 else if (j == 0) { maxLeft = A[i-1]; }
//                 else { maxLeft = Math.max(A[i-1], B[j-1]); }
//                 if ( (m + n) % 2 == 1 ) { return maxLeft; } // 奇数的话不需要考虑右半部分

//                 int minRight = 0;
//                 if (i == m) { minRight = B[j]; }
//                 else if (j == n) { minRight = A[i]; }
//                 else { minRight = Math.min(B[j], A[i]); }

//                 return (maxLeft + minRight) / 2.0; //如果是偶数的话返回结果
//             }
//         }
//         return 0.0;
//     }
// }

// 作者：windliang
// 链接：https://leetcode-cn.com/problems/median-of-two-sorted-arrays/solution/xiang-xi-tong-su-de-si-lu-fen-xi-duo-jie-fa-by-w-2/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
package main

import (
	"fmt"
	"math"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)
	if m > n {
		m, n = n, m
		nums1, nums2 = nums2, nums1
	}

	iMin := 0
	iMax := m
	for iMin <= iMax {
		i := (iMin + iMax) / 2
		j := (m+n+1)/2 - i
		if j != 0 && i != m && nums2[j-1] > nums1[i] {
			// i需要增大
			iMin = i + 1
		} else if i != 0 && j != n && nums1[i-1] > nums2[j] {
			// i需要减小
			iMax = i - 1
		} else {
			// 达到条件，边界值列出来单独考虑
			maxLeft := 0.0
			if i == 0 {
				maxLeft = float64(nums2[j-1])
			} else if j == 0 {
				maxLeft = float64(nums1[i-1])
			} else {
				maxLeft = math.Max(float64(nums1[i-1]), float64(nums2[j-1]))
			}
			if (m+n)%2 == 1 {
				// 奇数不需要考虑右半部分
				return maxLeft
			}

			minRight := 0.0
			if i == m {
				minRight = float64(nums2[j])
			} else if j == n {
				minRight = float64(nums1[i])
			} else {
				minRight = math.Min(float64(nums2[j]), float64(nums1[i]))
			}
			// 如果是偶数直接返回
			return (maxLeft + minRight) / 2.0
		}
	}
	return 0.0

}

func main() {
	fmt.Println(findMedianSortedArrays([]int{1, 3}, []int{2}))
	fmt.Println(findMedianSortedArrays([]int{1, 2}, []int{3, 4}))
	a := 1
	b := 2
	a, b = b, a
	fmt.Println(a, b)

	fmt.Println((a + b) / 2)
}
