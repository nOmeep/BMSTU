package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)
var (
	result int64 = 0
	N int64
)

type Stack struct {
	buf []*Vertex
	top int64
}

func (s *Stack) isEmpty() bool { // проверка на пустоту стека во избежание ошибок 
	return len((*s).buf) <= 0 // если стек пустой, будет "panic"
}

func (s *Stack) push(x *Vertex) { // добавить элемент на верхушку стека
	(*s).buf = append((*s).buf, x) // добавляю новое значение в конец стека
	(*s).top++
}

func (s *Stack) pop() *Vertex { // снять элемент с верхушки стека
	if s.isEmpty() {
		fmt.Println("UNDERFLOW: PANIC IN POP()")
		os.Exit(1)
	} 
	var index int64 = int64(len((*s).buf) - 1) // берем индекс у самого верхнего элемента
	var topElement *Vertex = (*s).buf[index] // достаем элемент по индексу выше
	(*s).buf = (*s).buf[:index] // не забываем удалять элемент из стека
	(*s).top--
	return topElement
}

var stack Stack

type Vertex struct {
	color string  
	command string

	index int64
	move int64   

	dom *Vertex
	sdom *Vertex
	label *Vertex
	parent *Vertex
	ancestor *Vertex

	bucket []*Vertex
	outgoing []*Vertex 
	incoming []*Vertex 
}

var (
	globalVertexNel int64 = 0
	vertexMap map[int64]int64 = make(map[int64]int64)
)

func main()  {
	var (
		U int64
		V int64
		currentAct string
	) 
	
	fmt.Scanf("%d", &N)
	V_G := make([]*Vertex, 0)
	V_G = setAllVertex()

	for i := range V_G {
		fmt.Scan(&U)
		fmt.Scan(&currentAct)
		V_G[i].command = currentAct
		if currentAct != "ACTION" {
			fmt.Scan(&V)
			V_G[i].move = V
		}
		vertexMap[U] = int64(i)
	}

	for index := range V_G {
		if strings.Compare(V_G[index].command, "ACTION") == 0 {
			if int64(index) != N - 1 {
				V_G[index].outgoing = append(V_G[index].outgoing, V_G[index + 1])
				V_G[index + 1].incoming = append(V_G[index+1].incoming, V_G[index])
			}
		} else if strings.Compare(V_G[index].command, "BRANCH") == 0 {
			var (
				tmp int64 = vertexMap[V_G[index].move]
			)
			V_G[index].outgoing = append(V_G[index].outgoing, V_G[tmp])
			V_G[tmp].incoming = append(V_G[tmp].incoming, V_G[index])
			if int64(index) != N - 1 {
				V_G[index].outgoing = append(V_G[index].outgoing, V_G[index + 1])
				V_G[index + 1].incoming = append(V_G[index+1].incoming, V_G[index])
			}
		} else if strings.Compare(V_G[index].command, "JUMP") == 0 {
			var (
				tmp int64 = vertexMap[V_G[index].move]
			)
			V_G[index].outgoing = append(V_G[index].outgoing, V_G[tmp])
			V_G[tmp].incoming = append(V_G[tmp].incoming, V_G[index])
		} else {
			fmt.Println("chto eto takoe")
		}
	}

	globalVertexNel++
	DFS(V_G[0])

	for i := 0; i < len(V_G); i++ {
		if strings.Compare(V_G[i].color, "black") == 0 {
			V_G[i] = V_G[len(V_G) - 1]
			V_G[len(V_G) - 1] = nil
			V_G = V_G[:len(V_G) - 1]
			i--
		} else {
			for j := 0; j < len(V_G[i].incoming); j++ {
				if strings.Compare(V_G[i].incoming[j].color, "black") == 0 {
					V_G[i].incoming[j] = V_G[i].incoming[len(V_G[i].incoming) - 1]
					V_G[i].incoming = V_G[i].incoming[:len(V_G[i].incoming) - 1]
					j--
				}
			}
		}
	}

	doSort(&V_G)

	N = int64(len(V_G))

	dominate(&V_G)
	
	for _, v := range V_G {
		for _, e := range v.incoming {
			for e != v && e!= nil {
				e = e.dom
			}
			if e == v {
				result++
				break 
			}
		}
	}
	fmt.Println(result)
}

func setSingleVertex() (tmp Vertex) {
	tmp.color = "black"
	tmp.command = ""
	tmp.outgoing = make([]*Vertex, 0)
	tmp.incoming = make([]*Vertex, 0)
	tmp.bucket = make([]*Vertex, 0)
	tmp.ancestor = nil

	return
}

func setAllVertex() []*Vertex {
	var result []*Vertex
	for i := 0; int64(i) < N; i++ {
		var tmp Vertex = setSingleVertex();
		tmp.sdom = &tmp
		tmp.label = &tmp
		result = append(result, &tmp)
	}

	return result
}

func FindMin(v *Vertex) (min *Vertex) {
	if (*v).ancestor == nil {
		min = v
	} else {
		var (
			u *Vertex
		)
		u = v
		for (*(*u).ancestor).ancestor != nil {
			stack.push(u)
			u = (*u).ancestor
		}
		for stack.top != 0 {
			v = stack.pop()
			if (*(*(*(*v).ancestor).label).sdom).index < (*(*(*v).label).sdom).index {
				(*v).label = (*(*v).ancestor).label
			}
			v.ancestor = u.ancestor
		}
		min = v.label
	}
	return
}

func dominate(allV *[]*Vertex) {
	*allV = Dominators(*allV)
}

func Dominators(V_G []*Vertex) []*Vertex{
	stack.buf = make([]*Vertex, N)
	
	for _, w := range V_G {
		if (*w).index == 1 {
			continue
		}
		for _, v := range (*w).incoming {
			var u *Vertex = FindMin(v)
			if (*(*u).sdom).index < (*(*w).sdom).index {
				(*w).sdom = (*u).sdom
			}
		}
		(*w).ancestor = (*w).parent
		(*(*w).sdom).bucket = append((*(*w).sdom).bucket, w)
		for _, v := range (*(*w).parent).bucket {
			u := FindMin(v)
			if (*u).sdom == (*v).sdom {
				(*v).dom = (*v).sdom
			} else {
				(*v).dom = u
			}
		}
		(*(*w).parent).bucket = nil
	}
	
	for _, w := range V_G {
		if (*w).index == 1 {
			continue
		}
		if (*w).dom != (*w).sdom {
			(*w).dom = (*(*w).dom).dom
		}
	}
	
	V_G[len(V_G) - 1].dom = nil

	return V_G
}

func DFS(r *Vertex)  {
	(*r).color = "white"
	globalVertexNel++
	(*r).index = globalVertexNel - 1
	for e := range (*r).outgoing {
		if strings.Compare((*(*r).outgoing[e]).color, "black") == 0 {
			(*(*r).outgoing[e]).parent = r
			DFS(r.outgoing[e])
		}
	}
}

func doSort(allV *[]*Vertex) {
	sort.Slice(*allV, func(i, j int) bool {
		return (*allV)[i].index > (*allV)[j].index
	})
}