package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const N = 100010
const M = N * 2

var h [N]int
var e [M]int
var ne [M]int
var idx int

func add(a, b int) {
	e[idx] = b
	ne[idx] = h[a]
	h[a] = idx
	idx++
}

func dfs(u int, stl []bool, n int) int {
	stl[u] = true
	res := 0
	size := 1
	for i := h[u]; i != -1; i = ne[i] {
		j := e[i]
		if !stl[j] {
			child := dfs(j, stl, n)
			res = max(res, child)
			size += child
		}
	}
	res = max(res, n-size)
	ans = min(ans, res)
	return size
}

var ans int

func main() {
	reader := bufio.NewReader(os.Stdin)
	n, _ := strconv.Atoi(readLine(reader))
	ans = n + 1
	stl := make([]bool, N)
	for i := 0; i < n-1; i++ {
		line := strings.Split(readLine(reader), " ")
		a, _ := strconv.Atoi(line[0])
		b, _ := strconv.Atoi(line[1])
		add(a, b)
		add(b, a)
	}
	for i := 0; i < N; i++ {
		h[i] = -1
	}
	dfs(1, stl, n)
	fmt.Println(ans)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err != nil {
		return ""
	}
	return string(str)
}
