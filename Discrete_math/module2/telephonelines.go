package main

import (
	"fmt"
	"os"
	"github.com/skorobogatov/input"
)

type Edge struct {
	length int64
	next int64
}

func setEdge(l int64, n int64) (tmp Edge) {
	tmp.length = l
	tmp.next = n
	return
}

type Vertex struct {
	edges []Edge
	name int64
	// на что указывает heap
	index int64
	key int64
	value int64
}

func setVertex(x int64) (tmp Vertex) {
	tmp.name = x
	tmp.index = -1
	tmp.edges = make([]Edge, 0)
	return
}

// Реализация очереди с приоритетами

type PriorityQueue struct {
	qHeap []*Vertex
	count int64 // количество элементов очереди
	cap int64
}

func queueEmpty(q *PriorityQueue) bool {
	return (*q).count == 0
}

func insert(q *PriorityQueue, v *Vertex) { // добавление вершины в очередь
	var i int64 = (*q).count

	if i == (*q).cap {
		fmt.Printf("PANIC\n")
		os.Exit(1)
	}

	(*q).count++ // (*q).count = i + 1
	(*q).qHeap[i] = v

	for i > 0 && (*q).qHeap[(i - 1) / 2].key > (*q).qHeap[i].key {
		(*q).qHeap[(i - 1) / 2], (*q).qHeap[i] = (*q).qHeap[i], (*q).qHeap[(i - 1) / 2]
		(*q).qHeap[i].index = i
		i = (i - 1) / 2
	}

	(*q).qHeap[i].index = i
}

func extractMin(q *PriorityQueue) *Vertex {
	var heapify func(*PriorityQueue, int64, int64);
	heapify = func(pq *PriorityQueue, i int64, n int64) {
		var (
			l int64
			j int64
			r int64
		)
		for {
			l = 2 * i + 1
			r = l + 1
			j = i
			if l < n && (*pq).qHeap[i].key > (*pq).qHeap[l].key {
				i = l
			}
			if r < n && (*pq).qHeap[i].key > (*pq).qHeap[r].key {
				i = r
			}
			if i == j {
				break
			}
			(*pq).qHeap[i], (*pq).qHeap[j] = (*pq).qHeap[j], (*pq).qHeap[i]
			(*pq).qHeap[i].index = i
			(*pq).qHeap[j].index = j
		}
	}

	if queueEmpty(q) {
		fmt.Printf("UNDERFLOW\n")
		os.Exit(1)
	}

	var ptr *Vertex = (*q).qHeap[0]
	(*q).count--

	if !queueEmpty(q) {
		(*q).qHeap[0] = (*q).qHeap[(*q).count]
		(*q).qHeap[0].index = 0
		heapify(q, 0, (*q).count)
	}

	return ptr
}

func decreaseKey(q *PriorityQueue, index int64, k int64) {
	var i int64 = index
	(*q).qHeap[index].key = k

	for i > 0 && (*q).qHeap[(i - 1) / 2].key > k {
		(*q).qHeap[(i - 1) / 2], (*q).qHeap[i] = (*q).qHeap[i], (*q).qHeap[(i - 1) / 2]
		(*q).qHeap[i].index = i
		i = (i - 1) / 2
	}

	(*q).qHeap[i].index = i
}

func MST_PRIM(vArr []*Vertex) (l int64) {
	var (
		v *Vertex = vArr[0]
		q PriorityQueue

		tmpN int64 = int64(len(vArr) - 1)
	)

	//fmt.Println("LENGTH TMP ", len(vArr) - 1)

	for ind,_ := range vArr { // !!!!!!!
		vArr[ind].index = -1 
	}

	q.qHeap = make([]*Vertex, tmpN)
	q.count = 0
	q.cap = tmpN;

	for ;; {
		//fmt.Println("I WAS HERE")
		v.index = -2 // 
		for _, val := range v.edges { // научен опытом, теперь валуе беру только у указателей 
			//fmt.Println("vArr[val.next].index = ", vArr[val.next].index)
			if vArr[val.next].index == -1 {
				vArr[val.next].key = val.length
				vArr[val.next].value = v.name
				insert(&q, vArr[val.next])
			} else {
				if vArr[val.next].index != -2 && val.length <= vArr[val.next].key {
					vArr[val.next].value = v.name
					decreaseKey(&q, vArr[val.next].index, val.length)
				}
			}
		}
		if queueEmpty(&q) {
			break;
		}
		v = extractMin(&q)
		//fmt.Println("ITS NAME - ", v.name)
		//fmt.Println("ITS KEY - ", v.key)
		l += v.key
		//fmt.Printf("Lenght is - %d\n", l)
	}
	
	return 
}

func main() {
	var (
		N int64 
		M int64
		vArr []*Vertex = make([]*Vertex, 0)
	)

	input.Scanf("%d", &N)
	input.Scanf("%d", &M)

	for i := 0; int64(i) < N; i++ {
		var tmp Vertex = setVertex(int64(i))
		vArr = append(vArr, &tmp)
	}

	var (
		u int64
		v int64 
		l int64 
	)

	for i := 0; int64(i) < M; i++ {
		input.Scanf("%d", &u)
		input.Scanf("%d", &v)
		input.Scanf("%d", &l)

		var tmp1 Edge = setEdge(l, v)
		vArr[u].edges = append(vArr[u].edges, tmp1)
		var tmp2 = setEdge(l, u)
		vArr[v].edges = append(vArr[v].edges, tmp2)
	}

	fmt.Println(MST_PRIM(vArr))
}
