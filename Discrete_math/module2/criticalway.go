package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Vertex struct {
	name int64
	enteredAction string
	dist int64
	enteredjobTime int64
	edges []int64
	parent []int64
	t1 int64
	low int64
	comp int64
	color string
}

type Stack struct {
	buf []*int64
	top int64
}

func (s *Stack) isEmpty() bool { // проверка на пустоту стека во избежание ошибок 
	return len((*s).buf) <= 0 // если стек пустой, будет "panic"
}

func (s *Stack) push(x int64) { // добавить элемент на верхушку стека
	(*s).buf = append((*s).buf, &x) // добавляю новое значение в конец стека
	(*s).top++
}

func (s *Stack) pop() int64 { // снять элемент с верхушки стека
	if s.isEmpty() {
		fmt.Println("UNDERFLOW: PANIC IN POP()")
		os.Exit(1)
	} 
	var index int64 = int64(len((*s).buf) - 1) // берем индекс у самого верхнего элемента
	var topElement *int64 = (*s).buf[index] // достаем элемент по индексу выше
	(*s).buf = (*s).buf[:index] // не забываем удалять элемент из стека
	(*s).top--
	return *topElement
}

var stack Stack

func Tarjan(V_G *[]Vertex)  {
	stack.top = 0
	stack.buf = make([]*int64, len(*V_G))
	for e, _ := range *V_G {
		if (*V_G)[e].t1 == 0 {
			visitVertexTarjan(V_G, int64(e))
		}
	}
}

func DOTgraph(vg *[]Vertex) {
	fmt.Println("digraph {")
	for _, value := range *vg {
		if strings.Compare(value.color, "black") == 0 {
			fmt.Printf("\t%s [label = \"%s(%d)\"]\n", value.enteredAction, value.enteredAction, value.enteredjobTime)
		}
		if strings.Compare(value.color,"red") == 0 {
			fmt.Printf("\t%s [label = \"%s(%d)\", color = red]\n", value.enteredAction, value.enteredAction, value.enteredjobTime)
		}
		if strings.Compare(value.color, "blue") == 0 {
			fmt.Printf("\t%s [label = \"%s(%d)\", color = blue]\n", value.enteredAction, value.enteredAction, value.enteredjobTime)
		}
	}
	for _, value1 := range *vg {
		for _, e := range value1.edges {
			if strings.Compare(value1.color, "blue") == 0 {
				fmt.Printf("\t%s -> %s [color = blue]\n", value1.enteredAction, (*vg)[e].enteredAction)
			}
			if strings.Compare(value1.color, "red") == 0 && strings.Compare((*vg)[e].color, "red") == 0 {
				var (
					flag bool = false
				)
				for _, valueTmp := range (*vg)[e].parent {
					if valueTmp == value1.name {
						flag = true
					}
				}
				if flag {
					fmt.Printf("\t%s -> %s [color = red]\n", value1.enteredAction, (*vg)[e].enteredAction)
				} else {
					fmt.Printf("\t%s -> %s\n", value1.enteredAction, (*vg)[e].enteredAction)
				}
			}
			if strings.Compare(value1.color, "black") == 0 || (strings.Compare(value1.color, "red") == 0 && (strings.Compare((*vg)[e].color, "black") == 0 || strings.Compare((*vg)[e].color, "blue") == 0)) {
				fmt.Printf("\t%s -> %s\n", value1.enteredAction, (*vg)[e].enteredAction)
			}
		}
	}
	fmt.Println("}")

	яусталрешатьграфы = true
}

func visitVertexTarjan(V_G *[]Vertex, v int64)  {
	(*V_G)[v].t1 = time
	(*V_G)[v].low = time
	time++
	stack.push(v)
	for _, e := range (*V_G)[v].edges {
		
		if (*V_G)[e].t1 == 0 {
			
			visitVertexTarjan(V_G, e)
		}
		if ((*V_G)[e].comp == -1) && ((*V_G)[v].low > (*V_G)[e].low) {
			(*V_G)[v].low = (*V_G)[e].low
		}
	}
	if (*V_G)[v].low == (*V_G)[v].t1 {
		var (
			u int64 = stack.pop()
		)
		var (
			flag bool = false
		)
		for _, value := range (*V_G)[u].edges {
			if value ==	u {
				flag = true
			}
		}
		if u == v && flag {
			coloredBlue(V_G, u)
		}
		(*V_G)[u].comp = count
		for u != v {
			u = stack.pop()
			(*V_G)[u].comp = count
			coloredBlue(V_G, u)
		}
		count++
	}
}

func relayRace(V_G *[]Vertex)  {
	var (
		end bool
	)
	for i, _ := range *V_G {
		if strings.Compare((*V_G)[i].color, "blue") != 0 {
			for _, e := range (*V_G)[i].edges {
				if relax(&(*V_G)[i], &(*V_G)[e], (*V_G)[e].enteredjobTime) {
					end = true
				}
			}
		}
	}
	if !end {
		return
	} else if end {
		relayRace(&(*V_G))
	}
}

func relax(u *Vertex, v *Vertex, w int64) bool {
	var (
		predicate bool = u.dist + w > v.dist
	)
	if predicate {
		(*v).dist = (*u).dist + w
		(*v).parent = make([]int64, 0)
		(*v).parent = append((*v).parent, (*u).name)
	}
	if (*u).dist + w == (*v).dist {
		var (
			flag bool
		)
		for _, value := range (*v).parent {
			if value == (*u).name {
				flag = true
			}
		}
		if !flag {
			(*v).parent = append((*v).parent, (*u).name)
		}
	}
	
	return predicate
}

func coloredBlue(V_G *[]Vertex, i int64)  {
	
	(*V_G)[i].color = "blue"
	(*V_G)[i].dist = -1 
	for _, e := range (*V_G)[i].edges {
		if strings.Compare((*V_G)[e].color, "blue") != 0 {
			coloredBlue(V_G, e)
		}
	}
}

func coloredRed(V_G *[]Vertex, i int64)  {
	(*V_G)[i].color = "red"
	for _, e := range (*V_G)[i].parent {
		if strings.Compare((*V_G)[e].color, "red") != 0 {
			coloredRed(V_G, e)
		}
	}
}

func addVertex(v_g *[]Vertex, act string, name int64, jobT int64) {
	var tmp Vertex
	tmp.color = "black"
	tmp.name = name
	tmp.edges = make([]int64, 0)
	tmp.parent = make([]int64, 0)
	tmp.enteredAction = act
	tmp.enteredjobTime = jobT
	tmp.comp = -1
	tmp.dist = jobT
	*v_g = append(*v_g, tmp)
}

func doColoredRed(v_g *[]Vertex, currmax int64) {
	for i, _ := range *v_g {
		if (*v_g)[i].dist == currmax && (strings.Compare((*v_g)[i].color, "red") == 0 || strings.Compare((*v_g)[i].color, "black") == 0) {
			coloredRed(&(*v_g), int64(i))
		}
	}
}

var (
	count int64 = 0
	time int64 = 1
	яусталрешатьграфы bool = true
)

func main()  {
	
	var (
		next bool
		connect bool
		V_G []Vertex = make([]Vertex, 0)
		vertexMap map[string]int64 = make(map[string]int64)
		firstAction string
		tokenProm string
		secondAction string
	)
	fmt.Scanf("%s %s", &firstAction, &tokenProm)

	if strings.Index(firstAction, ";") != -1 {
		next = true
		firstAction = strings.Replace(firstAction, firstAction[strings.Index(firstAction, ";"):], "",-1)
	}

	jobTimeInd, _:= strconv.Atoi(firstAction[strings.Index(firstAction, "(") + 1 : strings.Index(firstAction, ")")])

	firstAction = strings.Replace(firstAction, firstAction[strings.Index(firstAction, "("):], "", -1)
	vertexMap[firstAction] = 0
	addVertex(&V_G, firstAction, 0, int64(jobTimeInd))
	var (
		c int64 = 1
	)
	
	if !(tokenProm == "" && !next) {
		for {
			if next {
				connect = true
				next = false
			}
			tokenProm = ""
			fmt.Scanf("%s %s", &secondAction, &tokenProm)
			
			if strings.Index(secondAction, ";") != -1 {
				next = true
				secondAction = strings.Replace(secondAction, secondAction[strings.Index(secondAction, ";"):], "",-1)
			}

			if strings.Index(secondAction, "(") != -1 {
				
				jobTimeIndex, _ := strconv.Atoi(secondAction[strings.Index(secondAction, "(") + 1 : strings.Index(secondAction, ")")])
				secondAction = strings.Replace(secondAction, secondAction[strings.Index(secondAction, "("):], "", -1)
				vertexMap[secondAction] = c
				addVertex(&V_G, secondAction, int64(c), int64(jobTimeIndex))
				if !connect {
					var (
						t int64 = vertexMap[firstAction]
					)
					var (
						flag bool = false
					)
					for _, value := range V_G[t].edges {
						if value == c {
							flag = true
						}
					}
					if !flag {
						V_G[t].edges = append(V_G[t].edges, c)
					}
				}
				c++
			} else if strings.Index(secondAction, "(") == -1 {
				if !connect {
					var (
						t int64 = vertexMap[firstAction]
					)
					var (
						flag bool = false
					)
					for _, value := range V_G[t].edges {
						if value ==	vertexMap[secondAction] {
							flag = true
						}
					}
					if !flag {
						V_G[t].edges = append(V_G[t].edges, vertexMap[secondAction])
					}

				}
			}
			firstAction = secondAction
			if strings.Compare(tokenProm, "") == 0 && !next{
				break
			}
			connect = false
		}
	}

	Tarjan(&V_G)
	
	relayRace(&V_G)

	var (
		maximum int64 = V_G[0].dist
	)
	for i, _ := range V_G {
		if V_G[i].dist > maximum && (strings.Compare(V_G[i].color, "red") == 0 || strings.Compare(V_G[i].color, "black") == 0){
			maximum = V_G[i].dist
		}
	}
	
	doColoredRed(&V_G, maximum)

	if яусталрешатьграфы {
		DOTgraph(&V_G)
		os.Exit(1);
	}

	fmt.Println("Я")
	fmt.Println("устал")
	fmt.Println("решать")
	fmt.Println("задачи")
	fmt.Println("пощады")
}