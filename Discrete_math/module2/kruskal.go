package main

import (
	"fmt"
	"math"
	"github.com/skorobogatov/input"
)

type Edge struct {
	u int64 
	v int64 
	len float64
}

type Vertex struct {
	x int64 
	y int64 
	d int64

	parent *Vertex
	edge *Edge 
}

var (
	N int64
	result float64 = 0
)

func main() {
	var (
		x int64
		y int64
	)

	input.Scanf("%d", &N)

	var (
		allV []Vertex = make([]Vertex, N)
		allE []Edge = make([]Edge, N * (N - 1) / 2)
	)

	for i := 0; int64(i) < N; i++ {
		input.Scanf("%d", &x)
		input.Scanf("%d", &y)

		allV[i].x = x;
		allV[i].y = y
		allV[i].parent = &allV[i] // сам себе, так сказать, будет поначалу
	}

	var (
		count int64 = 0
	)

	for i := 0; int64(i) < N; i++ {
		for j := i + 1; int64(j) < N; j++ {
			allE[count].len = math.Sqrt(math.Pow(float64(allV[j].x - allV[i].x), 2) + math.Pow(float64(allV[j].y - allV[i].y), 2)) // находим длину по классической формуле 
			//fmt.Printf("LENGTH - %f\n" , allE[count].len)
			allE[count].v = int64(i)
			allE[count].u = int64(j)
			count++
		}
	}

	count = makeHeapify(count, &allE)
	

	MST_Kruskal(&allE, &allV, count)

	fmt.Printf("%.2f\n", result)
}

func MST_Kruskal(e *[]Edge, v *[]Vertex, c int64) {
	var i int64
	for i < N - 1{
		var (
			v1 int64 = (*e)[0].u
			v2 int64 = (*e)[0].v
		)

		if find(&(*v)[v1]) != find(&(*v)[v2]) {
			result += (*e)[0].len
			//fmt.Printf("RESULT IS %f\n", result)
			union(find(&(*v)[v1]), find(&(*v)[v2]))
			i++
		}

		(*e)[0], (*e)[c] = (*e)[c], (*e)[0]
		heapify(e, 0, c)
		c--
	}
}

func makeHeapify(c int64, e *[]Edge) int64 {
	for i := c / 2 - 1; i > -1; i-- {
		heapify(e, i, c)
	}
	return c - 1
}

func heapify(arr *[]Edge, i int64, n int64) {
	var (
		l int64
		j int64
		r int64
	)
	for {
		l = 2 * i + 1
		r = l + 1
		j = i
		if l < n && (*arr)[i].len > (*arr)[l].len {
			i = l
		}
		if r < n && (*arr)[i].len > (*arr)[r].len {
			i = r
		}
		if i == j {
			break
		}
		(*arr)[i], (*arr)[j] = (*arr)[j], (*arr)[i]
		//fmt.Printf("CHANGES - %d %d\n", (*arr)[j].u, (*arr)[i].u)
	}
}

func union(x *Vertex, y *Vertex) {
	var (
		rX *Vertex = find(x)
		rY *Vertex = find(y)
	)
	if (*rX).d < (*rY).d {
		(*rX).parent = rY
	} else {
		(*rY).parent = rX
		if ((*rX).d == (*rY).d) && (rX != rY) {
			(*rX).d++
		}
	}
}

func find(x *Vertex) *Vertex {
	if *(*x).parent == (*x) {
		return x
	}
	return find((*x).parent)
}