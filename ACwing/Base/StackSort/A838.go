//  838. 堆排序
//     题目
//     提交记录
//     讨论
//     题解
//     视频讲解
//
// 输入一个长度为 n
// 的整数数列，从小到大输出前 m
// 小的数。
//
// 输入格式
// 第一行包含整数 n
// 和 m
// 。
//
// 第二行包含 n
// 个整数，表示整数数列。
//
// 输出格式
// 共一行，包含 m
// 个整数，表示整数数列中前 m
// 小的数。
//
// 数据范围
// 1≤m≤n≤105
// ，
// 1≤数列中元素≤109
// 输入样例：
// 5 3
// 4 5 1 3 2
// 输出样例：
// 1 2 3
package main

import (
	_ "Algorithm/ACwing/Base/Stack"
	"fmt"
)

var (
	n, m int
	N    = 100010
	h    = make([]int, N)
	size = 0
)

// 左右子树
func down(u int) {
	//t表示的为
	t := u
	if u*2 <= size && h[u*2] < h[t] {
		t = u * 2
	}
	if u*2+1 <= size && h[u*2+1] < h[t] {
		t = u*2 + 1
	}
	if u != t {
		swap(h, u, t)
		down(t)
	}
}

// 父节点
func up(u int) {
	for u/2 > 0 && h[u/2] > h[u] {
		swap(h, u/2, u)
		u /= 2
	}
}

func main() {

	fmt.Scan(&n, &m)
	for i := 1; i <= n; i++ {
		fmt.Scan(&h[i])
	}
	size = n

	for i := n / 2; i != 0; i-- {
		down(i)
	}

	for ; m > 0; m-- {
		fmt.Print(h[1], " ")
		swap(h, 1, size)
		size--
		down(1)
	}
}

func swap(a []int, i, j int) {
	tmp := a[i]
	a[i] = a[j]
	a[j] = tmp
}
