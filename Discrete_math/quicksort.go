package main

import ("fmt")

func main(){
	var n int
	fmt.Scan(&n)
	var mas []int = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&mas[i])
	}
	qsort(n, 
		func (i, j int) bool {
			return mas[i] < mas[j]
		}, 
		func (i, j int) {
			mas[i], mas[j] = mas[j], mas[i]
		})
	for _, value := range mas {
		fmt.Print(value, " ")
	}

}

//если я все правильно понял, суть задачи реализовать весь алгоритм сортировки в одной функции qsort

func qsort(n int, less func(i, j int) bool, swap func(i, j int)) { 
	var low, high int = 0, (n - 1)
	
	var Partition func(low, high int) int
	Partition = func(low, high int) int{
		var i, j int = low, low
		for j < high {
			if less(j, high) {
				swap(i, j)
				i++
			}
			j++
		}
		swap(i, high)
		return i
	}

	var QuickSortRec func(low, high int)
	QuickSortRec = func(low, high int) {
		if low < high {
			var q int = Partition(low, high)
			QuickSortRec(low, q - 1)
			QuickSortRec(q + 1, high)
		}
	}

	QuickSortRec(low, high)

} 