package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	var (
		n           int                    // число, у которого надо найти все делители
		allDivisors []int = make([]int, 0) // массив со всеми делителями входного числа
	)

	fmt.Scan(&n)

	var limit int = int(math.Sqrt(float64(n)))

	for i := 1; i <= limit; i++ {
		if n%i == 0 {
			allDivisors = append(allDivisors, i)
			if Exist(allDivisors, n / i) == -1 { // избегаем повторения
				allDivisors = append(allDivisors, n / i)
			}
		}
	}

	sort.Ints(allDivisors)
	//fmt.Println(allDivisors)

	fmt.Println("graph {")

	for i := len(allDivisors) - 1; i >= 0; i-- {
		fmt.Printf("\t%d\n", allDivisors[i])
	} 

	for u := len(allDivisors) - 1; u > 0; u-- {
		for v := u - 1; v >= 0; v-- {
			if allDivisors[u] % allDivisors[v] == 0 {
				var exitStatus bool = true

				for w := u - 1; w > v; w-- { // есть ли такой w, что u % w == 0 и w % v == 0
					if allDivisors[u] % allDivisors[w] == 0 && allDivisors[w] % allDivisors[v] == 0 {
						exitStatus = false
						break
					}
				}
				
				if exitStatus {
					fmt.Printf("\t%d -- %d\n", allDivisors[u], allDivisors[v])
				}
			}
		}
	}

	fmt.Println("}")

}


// вспомогательная функция для избежания повторений
func Exist(mas []int, x int) int {
	for i := 0; i < len(mas); i++ {
		if mas[i] == x {
			return mas[i]
		}
	}
	return -1
}
