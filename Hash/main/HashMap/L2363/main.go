package main

import "sort"

func mergeSimilarItems(items1 [][]int, items2 [][]int) [][]int {
	hashMap := make(map[int]int)
	for _, item := range items1 {
		v := item[0]
		w := item[1]
		hashMap[v] += w
	}
	for _, item := range items2 {
		v := item[0]
		w := item[1]
		hashMap[v] += w
	}
	var res [][]int
	for v, w := range hashMap {
		res = append(res, []int{v, w})
	}
	sort.Slice(res, func(i, j int) bool {
		return res[i][0] < res[j][0]
	})
	return res
}
