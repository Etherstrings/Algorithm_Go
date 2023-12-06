package main

func waysToBuyPensPencils(total int, cost1 int, cost2 int) int64 {
	var res = 0
	for i := 0; i <= total/cost1; i++ {
		res += (total-i*cost1)/cost2 + 1
	}
	return int64(res)
}
