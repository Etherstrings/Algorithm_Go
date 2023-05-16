package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var n, head, idx int
var N = 100010
var e = make([]int, N)
var ne = make([]int, N)

func initList() {
	head = -1
	idx = 1
}

func addToHead(x int) {
	e[idx] = x
	ne[idx] = head
	head = idx
	idx++
}

func remove(k int) {
	if k == 0 {
		head = ne[head]
	} else {
		ne[k] = ne[ne[k]]
	}
}

func insert(k int, x int) {
	e[idx] = x
	ne[idx] = ne[k]
	ne[k] = idx
	idx++
}

func main() {
	initList()

	reader := bufio.NewReader(os.Stdin)
	fmt.Scanln(&n)

	for i := 0; i < n; i++ {
		var k, x int
		s, _ := reader.ReadString('\n')
		s = strings.TrimSpace(s)
		op := strings.Split(s, " ")[0]
		if op == "H" {
			x, _ = strconv.Atoi(strings.Split(s, " ")[1])
			addToHead(x)
		} else if op == "D" {
			k, _ = strconv.Atoi(strings.Split(s, " ")[1])
			if k == 0 {
				head = ne[head]
			} else {
				remove(k - 1)
			}
		} else {
			k, _ = strconv.Atoi(strings.Split(s, " ")[1])
			x, _ = strconv.Atoi(strings.Split(s, " ")[2])
			insert(k-1, x)
		}
	}

	for i := head; i != -1; i = ne[i] {
		fmt.Printf("%d ", e[i])
	}
}
