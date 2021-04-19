package main

import (
	"fmt"
	"strings"
	"github.com/skorobogatov/input"
)

var brNel int = 0;

type Vertex struct {
	color string // white or black
	comp int
	edge []int
	parent *Vertex
}

func initializeVertex() Vertex {
	var tmp Vertex
	tmp.color = "white"
	tmp.comp = -1
	tmp.edge = make([]int, 0)
	return tmp
}

// Реализация очереди

type Queue struct {
	buf []*Vertex
	cap int
	count int 
	head int
	tail int
}

// Решение задачи

func DFS1(v *[]Vertex, q *Queue) {
	for _, value := range *v {
		value.color = "white"
	}
	for index, value := range *v {
		if strings.Compare(value.color, "white") == 0 {
			//fmt.Println("I WAS HERE DFS1")
			brNel--;
			//fmt.Println(brNel)
			VV1(v, q, &(*v)[index])
		}
	}
}

func VV1(v *[]Vertex, q *Queue, singleV *Vertex) {
	var Enqueue func(*Queue, *Vertex) // добавление элемента в очередь
	Enqueue = func(q *Queue, v *Vertex) {
		(*q).buf[(*q).tail] = v;
		(*q).tail++
		if (*q).tail == (*q).cap {
			(*q).tail = 0;
		}
		(*q).count++; // увеличиваем количество элементов 
	}

	(*singleV).color = "black"
	Enqueue(q, singleV);
	for _, value := range (*singleV).edge {
		if strings.Compare((*v)[value].color, "white") == 0 {
			(*v)[value].parent = singleV
			VV1(v, q, &(*v)[value])
		}
	}
}

func DFS2(v *[]Vertex, q *Queue) {
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

	var c int = 0;
	for (*q).count > 0 {
		var tmp *Vertex = Dequeue(q);
		if (*tmp).comp == -1 {
			//fmt.Println("I WAS HERE DFS 2")
			VV2(v, tmp, c)
			brNel++
			//fmt.Println(brNel)
			c++
		}
	}
}

func VV2(v *[]Vertex, singleV *Vertex, c int) {
	(*singleV).comp = c;
	for _, value := range (*singleV).edge {
		if ((*v)[value].comp == -1) && ((*v)[value].parent != singleV) {
			VV2(v, &((*v)[value]), c);
		}
	}
}

func main() {
	var (
		N int 
		M int

		u int
		v int 

		vertexArr []Vertex = make([]Vertex, 0);
		queue Queue
	)

	input.Scanf("%d", &N);
	input.Scanf("%d", &M);

	queue.buf = make([]*Vertex, N);
	queue.cap = N;
	queue.count = 0
	queue.head = 0
	queue.tail = 0

	for i := 0; i < N; i++ {
		vertexArr = append(vertexArr, initializeVertex())
	}	

	for i := 0; i < M; i++ {
		input.Scanf("%d %d", &u, &v)
		vertexArr[u].edge = append(vertexArr[u].edge, v)
		vertexArr[v].edge = append(vertexArr[v].edge, u)
	}

	DFS1(&vertexArr, &queue);
	DFS2(&vertexArr, &queue);

	fmt.Println(brNel)
}
