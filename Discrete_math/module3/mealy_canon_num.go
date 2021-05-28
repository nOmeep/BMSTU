package main

import (
	"fmt"
	"sort"
	"strings"
)

type Edge struct {
	to int64
	symbol string
}

type Vertex struct {
	color string
	canon int64
	edge []Edge
}

func increase(i *int) {
	*i++;
}

func show_result(automat *[]Vertex, n int64, m int64) {
	fmt.Println(n)
	fmt.Println(m)
	fmt.Println(0)
 	for index := range (*automat) {
 		if (*automat)[index].canon != -1 {
			for i := 0; int64(i) < m; i++ {
				fmt.Printf("%d ", (*automat)[index].edge[i].to)
			}
			fmt.Println()
		}
 	}
	for index := range (*automat) {
		if (*automat)[index].canon != -1 {
			for i := 0; int64(i) < m; i++ {
				fmt.Printf("%s ", (*automat)[index].edge[i].symbol)
			}
			fmt.Println()
		}
	}
}

var (
	count int = 0
)

func DFS(automat *[]Vertex, v int64)  {
	(*automat)[v].color = "white"
	(*automat)[v].canon = int64(count)
	increase(&count)
	for _, e := range (*automat)[v].edge {
		if strings.Compare((*automat)[e.to].color, "black") == 0 {
			DFS(automat, e.to)
		}
	}
}

func create_vertex() (result Vertex) {
	result.color = "black"
	result.canon = -1
	result.edge = make([]Edge, 0)
	return
}

func scan_and_dfs(automat *[]Vertex, n int64, m int64, q0 int64) {
	for i := 0; int64(i) < n; i++ {
		var tmpVertex Vertex = create_vertex();
		for j := 0; int64(j) < m; j++ {
			var (
				tmpEdge Edge
				out int64
			)
			fmt.Scan(&out)
			tmpEdge.to = out
			tmpVertex.edge = append(tmpVertex.edge, tmpEdge)
		}
		(*automat) = append((*automat), tmpVertex)
 	}
 	for i := 0; int64(i) < n; i++ {
 		for j := 0; int64(j) < m; j++ {
			var (
				symbol string
			)
 			fmt.Scan(&symbol)
			(*automat)[i].edge[j].symbol = symbol
		}
 	}

	DFS(automat, q0);
}

func do_sort(a *[]Vertex) {
	sort.Slice(*a, func(i, j int) bool {
		return (*a)[i].canon < (*a)[j].canon
	})
}

func main() {
	var (
		n int64
		m int64
		q0 int64
	)
	var automat = make([]Vertex, 0)
	fmt.Scan(&n)
	fmt.Scan(&m)
	fmt.Scan(&q0)
	
	scan_and_dfs(&automat, n, m, q0)

 	for v, _ := range automat {
 		if automat[v].canon == -1 {
 			n--
		}
 		for e, _ := range automat[v].edge {
 			automat[v].edge[e].to = automat[automat[v].edge[e].to].canon
		}
	}

	do_sort(&automat)

 	show_result(&automat, n, m)
}