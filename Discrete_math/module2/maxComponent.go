package main

import (
	"fmt"
	"strings"
	"github.com/skorobogatov/input"
)

type Component struct {
	compNum   int
	vertexNum int
	edgeNum   int
	minVertex int
}

type inputEdge struct {
	u int
	v int
}

type Vertex struct {
	color string // был или не был
	comp  int    // компонента
	name  int
	edges []int // ребра
}

func main() {
	var (
		component Component
		N         int // количество вершин графа
		M         int // количество ребер графа

		u int
		v int

		tmpEdge inputEdge

		allEdges []inputEdge
	)
	component.compNum = -1
	component.edgeNum = 0
	component.minVertex = -1
	component.vertexNum = 0

	input.Scanf("%d ", &N)
	input.Scanf("%d", &M)

	var vertexArr []Vertex = make([]Vertex, 0)

	for i := 0; i < N; i++ {
		var (
			tmp Vertex
		)
		tmp.color = "black"
		tmp.comp = -1
		tmp.name = i
		tmp.edges = make([]int, 0)

		vertexArr = append(vertexArr, tmp)
	}

	for i := 0; i < M; i++ {
		input.Scanf("%d", &u)
		input.Scanf("%d", &v)

		tmpEdge.u = u
		tmpEdge.v = v

		allEdges = append(allEdges, tmpEdge)

		vertexArr[u].edges = append(vertexArr[u].edges, v)
		vertexArr[v].edges = append(vertexArr[v].edges, u)

	}

	// ФУНКЦИИ ОБХОДА В ГЛУБИНУ
	var VisitVertex1 func(*[]Vertex, *Component, *Vertex)
	VisitVertex1 = func(verArr *[]Vertex, comp *Component, v *Vertex) {
		(*v).color = "white"
		(*v).comp = (*comp).compNum
		(*comp).vertexNum++
		if (*v).name < (*comp).minVertex {
			(*comp).minVertex = (*v).name
		}
		for _, value := range (*v).edges {
			(*comp).edgeNum++
			if strings.Compare((*verArr)[value].color, "black") == 0 {
				VisitVertex1(verArr, comp, &((*verArr)[value]))
			}
		}
	}

	var DFS func(*[]Vertex, *Component)
	DFS = func(verArr *[]Vertex, comp *Component) {
		var (
			tmp Component
		)

		tmp.compNum = 0
		for index, value := range *verArr {
			if strings.Compare(value.color, "black") == 0 {
				VisitVertex1(verArr, &tmp, &(*verArr)[index])
				if tmp.vertexNum > comp.vertexNum {
					*comp = tmp
				}
				if tmp.vertexNum == comp.vertexNum {
					if tmp.edgeNum > comp.edgeNum {
						*comp = tmp
					}
					if tmp.edgeNum == comp.edgeNum {
						if tmp.minVertex > comp.minVertex {
							*comp = tmp
						}
					}
				}
				tmp.vertexNum = 0
				tmp.edgeNum = 0
				tmp.minVertex = -1
				tmp.compNum++
			}
		}
	}
	// КОНЕЦ ФУНКЦИЙ ОБХОДА В ГЛУБИНУ

	DFS(&vertexArr, &component)

	// РИСОВАНИЕ САМОГО ГРАФА
	fmt.Println("graph {")
	for _, value := range vertexArr {
		if value.comp == component.compNum {
			fmt.Printf("\t%d [color = red]\n", value.name)
		} else {
			fmt.Printf("\t%d\n", value.name)
		}
	}
	for _, value := range allEdges {
		if vertexArr[value.u].comp == component.compNum {
			fmt.Printf("\t%d -- %d [color = red]\n", value.u, value.v)
		} else {
			fmt.Printf("\t%d -- %d\n", value.u, value.v)
		}
	}
	fmt.Println("}")
}