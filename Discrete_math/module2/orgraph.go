package main 

import (
	"fmt"
)

var (
	N int64
	M int64
)

func createArrays(in *[][]int64, out *[][]int64) {
	for i := 0; int64(i) < N; i++ {
		(*in)[i] = make([]int64, 0)
		(*out)[i] = make([]int64, 0)
	}
}

func fill(orgB *[]int64, needB *[]int64, outs *[][]int64) {
	var baseSwap func(int64, int64, *[]int64, *[][]int64)
	baseSwap = func(oldIndex int64, newIndex int64, b *[]int64, out *[][]int64) {
		if (*b)[oldIndex] != newIndex {
			(*b)[oldIndex] = newIndex
			for i := 0; i < len((*out)[oldIndex]); i++ {
				baseSwap((*out)[oldIndex][i], newIndex, b, out)
			}
		}
	}

	for i := 0; int64(i) < N; i++ {
		if (*orgB)[i] == -1 {
			(*needB) = append((*needB), int64(i))
			baseSwap(int64(i), int64(i), orgB, outs)
		}
	}
}

func main() {
	fmt.Scanf("%d", &N)
	fmt.Scanf("%d", &M)

	var (
		allInV [][]int64 = make([][]int64, N)
		allOutV [][]int64 = make([][]int64, N)
		orgBase []int64 = make([]int64, N)
		isBase []int64 = make([]int64, N)
	)

	createArrays(&allInV, &allOutV)

	for i := 0; int64(i) < N; i++ {
		orgBase[i] = -1
	}

	for i := 0; int64(i) < M; i++ {
		var (
			v1 int64
			v2 int64
		)

		fmt.Scanf("%d", &v1)
		fmt.Scanf("%d", &v2)

		allOutV[v1] = append(allOutV[v1], v2)
		allInV[v2] = append(allInV[v2], v1)
	}
	
	fill(&orgBase, &isBase, &allOutV)

	for i := 0; i < len(isBase); i++ {
		if orgBase[isBase[i]] == isBase[i] {
			fmt.Print(isBase[i], " ")
			orgBase[isBase[i]] = -1
		}
	}

	fmt.Println()
}