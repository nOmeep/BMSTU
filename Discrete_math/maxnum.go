package main

import (
	"fmt"
)

func main() {
	var nel int
	fmt.Scan(&nel)
	var biggestVal []int = make([]int, nel)
	for i := 0; i < nel; i++ {
		fmt.Scan(&biggestVal[i])
	}

	var reverse func([]int)
	reverse = func(mas []int) {
		for i, j := 0, len(mas)-1; i < j; i, j = i+1, j-1 {
			mas[i], mas[j] = mas[j], mas[i]
		}
	}

	var compare func(int, int) int
	compare = func(a, b int) int {
		var (
			firstNum []int = make([]int, 0)
			secondNum []int = make([]int, 0)
			i, i1 int = 0, 0
		)

		//fmt.Print(a, b, "\n")
		
		for a > 9 {
			firstNum = append(firstNum, a % 10)
			a /= 10
		}
		firstNum = append(firstNum, a % 10)
		reverse(firstNum)
		
		for b > 9 { 
			secondNum = append(secondNum, b % 10)
			b /= 10
		}
		secondNum = append(secondNum, b % 10) 
		reverse(secondNum)

		for ;i < len(firstNum) && i1 < len(secondNum); i, i1 = i + 1, i1 + 1 {
			if firstNum[i] > secondNum[i1] {
				return 1
			} else if firstNum[i] < secondNum[i1] {
				return -1
			}
		}
		
		var length1, length2 int = len(firstNum), len(secondNum)

		if length1 > length2 {
			i1--
			if firstNum[i] > secondNum[i1] {
				return 1
			}
			return -1
		} else if length1 < length2 {
			i--
			if firstNum[i] > secondNum[i1] {
				return 1
			} 
			return -1
		} 
		
		return 0
	}

	var sorting func(int, func(int, int) int)
	sorting = func(n int, compare func(int, int) int) {
		var (
			t int = (n - 1)
			bound int
			i int
		)
		for t > 0 {
			bound = t
			t = 0
			i = 0
			for i < bound {
				if 1 == compare(biggestVal[i + 1], biggestVal[i]) {
					biggestVal[i + 1], biggestVal[i] = biggestVal[i], biggestVal[i + 1]
					t = i
				}
				i++
			}
		}
	}

	sorting(nel, compare)
	
	for _, value := range biggestVal {
		fmt.Print(value)
	}
	fmt.Println()
}