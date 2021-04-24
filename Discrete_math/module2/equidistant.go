package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/skorobogatov/input"
)

type Vertex struct {
	color string 
	index int
	edge []int
}

// Очередь
type Queue struct {
	buf []*Vertex
	cap int 
	count int 
	head int 
	tail int
}

// Конец очереди


func main() {
	var (
		N int // vertexes
		M int // ребра 
		K int // опорные вершины

		u int 
		v int 

		u1 int 

		vArr []Vertex = make([]Vertex, 0)

		queue Queue
	)

	input.Scanf("%d", &N)
	input.Scanf("%d", &M)

	queue.buf = make([]*Vertex, N)
	queue.count = 0
	queue.cap = N
	queue.head = 0
	queue.tail = 0

	for i := 0; i < N; i++ {
		vArr = append(vArr, initVertex(i))
	}
	for i := 0; i < M; i++ {
		input.Scanf("%d %d", &u, &v)

		vArr[u].edge = append(vArr[u].edge, v)
		vArr[v].edge = append(vArr[v].edge, u)
	}

	input.Scanf("%d", &K)
	distances := make([][]int, K)
	for i := 0; i < K; i++ {
		input.Scanf("%d", &u1)

		distances[i] = make([]int, N)

		BFS(vArr, u1, &distances[i], &queue)
	}

	var (
		nedfulDistances []int = make([]int, 0) 
		flag bool = false
	)

	for i := 0; i < N; i++ {
		flag = true

		for j := 0; j < K - 1; j++ {
			if distances[j][i] != distances[j + 1][i] || distances[j][i] == 0 {
				flag = false
				//fmt.Print("Before break\n")
				break
			}
		}
		if flag == true {
			nedfulDistances = append(nedfulDistances, i)
			//fmt.Printf("Appended\n")
		}
	}

	sort.Ints(nedfulDistances)
	if len(nedfulDistances) == 0 {
		fmt.Printf("минус")
	} else {
		for index := range nedfulDistances {			fmt.Printf("%d ", nedfulDistances[index])
		}
	}

	fmt.Printf("\n");
}

func BFS(v []Vertex, b int, dist *[]int, queue *Queue) {
	
	var Enqueue func(*Queue, *Vertex) // добавление элемента в очередь
	Enqueue = func(q *Queue, v *Vertex) {
		//fmt.Printf("Enqueded\n");
		(*q).buf[(*q).tail] = v;
		(*q).tail++
		if (*q).tail == (*q).cap {
			(*q).tail = 0;
		}
		(*q).count++; // увеличиваем количество элементов 
	}


	var Dequeue func(*Queue) *Vertex // yдаление и возврат элемента
	Dequeue = func(q *Queue) *Vertex {
		var v *Vertex = (*q).buf[(*q).head];
		(*q).head++
		if (*q).head == (*q).cap {
			(*q).head = 0;
		}
		(*q).count-- // не забываем уменьшать количество
		return v;
	}

	var Empty func(*Queue) bool
	Empty = func(q *Queue) bool {
		return (*q).count <= 0;
	}

	for index := range v {
		v[index].color = "black"
	}

	v[b].color = "white"
	//fmt.Printf("Enqueded %d\n", b);
	Enqueue(queue, &v[b])

	for !Empty(queue) {
		//fmt.Printf("count > 0\n")

		tmp := Dequeue(queue)
		//fmt.Printf("Dequeded\n")

		for _, value := range (*tmp).edge {
			if strings.Compare(v[value].color, "black") == 0 {
				(*dist)[value] = (*dist)[tmp.index] + 1
				v[value].color = "white"
				Enqueue(queue , &v[value])
				//fmt.Printf("Enqueded inside %d\n", value);
			}
		}
	}
}

func initVertex(index int) Vertex {
	var tmp Vertex

	tmp.index = index
	tmp.color = "black"
	tmp.edge = make([]int, 0)

	return tmp
}