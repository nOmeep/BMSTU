package main 

import (
	"fmt"
	//"github.com/skorobogatov/input" 
	"os"
	"strings"
)

var (
	N int64
)

type Vertex struct {
	x int64
	y int64
	key int64
	index int64
}

// ОЧЕРЕДЬ С ПРОИРИТЕТАМИ

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

func changeColorToWhite(c *string) {
	(*c) = "white"
}

func makeTmpVertex(i int64, j int64, swtp *[][]int64) (tmp Vertex) { 
	tmp.x = i
	tmp.y = j
	tmp.key = (*swtp)[i][j]

	return
}

func findShortestWay(q *PriorityQueue, swtp *[][]int64, lba *[][]int64) {
	for !queueEmpty(q) {
		var currMin *Vertex = extractMin(q)
		var (
			i int64 = (*currMin).x
			j int64 = (*currMin).y
		)

		changeColorToWhite(&colors[i][j]) // посещаем

		
		if i > 0 && (*swtp)[i - 1][j] > (*swtp)[i][j] + (*lba)[i - 1][j] {
			(*swtp)[i - 1][j] = (*swtp)[i][j] + (*lba)[i - 1][j]
			if strings.Compare(colors[i - 1][j], "black") == 0 {
				var tmp Vertex = makeTmpVertex(i - 1, j, swtp)
				insert(q, &tmp)
			}
		}
		
		if i < N - 1 && (*swtp)[i + 1][j] > (*swtp)[i][j] + (*lba)[i + 1][j] {
			(*swtp)[i + 1][j] = (*swtp)[i][j] + (*lba)[i + 1][j]
			if strings.Compare(colors[i + 1][j], "black") == 0 {
				var tmp Vertex = makeTmpVertex(i + 1, j, swtp)
				insert(q, &tmp)
			}
		}
	
		if j > 0 && (*swtp)[i][j - 1] > (*swtp)[i][j] + (*lba)[i][j - 1] {
			(*swtp)[i][j - 1] = (*swtp)[i][j] + (*lba)[i][j - 1]
			if strings.Compare(colors[i][j - 1], "black") == 0 {
				var tmp Vertex = makeTmpVertex(i, j - 1, swtp)
				insert(q, &tmp)
			}
		}
	
		if j < N - 1 && (*swtp)[i][j + 1] > (*swtp)[i][j] + (*lba)[i][j + 1] {
			(*swtp)[i][j + 1] = (*swtp)[i][j] + (*lba)[i][j + 1]
			if strings.Compare(colors[i][j + 1], "black") == 0 {
				var tmp Vertex = makeTmpVertex(i, j + 1, swtp)
				insert(q, &tmp)
			}
		}
	}
}

var (
	colors [][]string
)

func startCreate(lengthsBetweenAdjacent *[][]int64, shortestWayToPoints *[][]int64) {
	*lengthsBetweenAdjacent = make([][]int64, N)
	*shortestWayToPoints = make([][]int64, N)
	colors = make([][]string, N)
	for i := 0; int64(i) < N; i++ {
		(*lengthsBetweenAdjacent)[i] = make([]int64, N)
		(*shortestWayToPoints)[i] = make([]int64, N)
		colors[i] = make([]string, N)
		for i1 := 0; int64(i1) < N; i1++ {
			(*shortestWayToPoints)[i][i1] = 9223372036854775807
			colors[i][i1] = "black"
			fmt.Scan(&(*lengthsBetweenAdjacent)[i][i1])
		}
	}

	(*shortestWayToPoints)[0][0] = (*lengthsBetweenAdjacent)[0][0]
}

func initQ(q *PriorityQueue) {
	(*q).count = 0
	(*q).cap = N * N
	(*q).qHeap = make([]*Vertex, N * N)
}

func main() {
	fmt.Scan(&N)

	var (
		q PriorityQueue
	)

	initQ(&q)

	var (
		lengthsBetweenAdjacent [][]int64
		shortestWayToPoints [][]int64 
	)

	startCreate(&lengthsBetweenAdjacent, &shortestWayToPoints);

	var (
		tmpV Vertex
	)

	tmpV.key = shortestWayToPoints[0][0]
	tmpV.x = 0
	tmpV.y = 0

	insert(&q, &tmpV)

	findShortestWay(&q, &shortestWayToPoints, &lengthsBetweenAdjacent)

	fmt.Println(shortestWayToPoints[N - 1][N - 1])
}