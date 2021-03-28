package main

// это не задача, а казино 

// slave: oh shit i'm sorry.
// master: sorry for what?..

import (
	"fmt"
)

func main() {
	var (
		startString string
		symb1 rune
		symb2 rune
	)

	fmt.Scanf("%s\n%c %c", &startString, &symb1, &symb2)

	fmt.Println(findMinDistance(startString, symb1, symb2))
}

func findMinDistance(str string, symb1 rune, symb2 rune) (distance int) {
	distance = 100000002
	var (
		indexSymb1 int = -1 // индекс последнего встретившегося 1го символа
		indexSymb2 int = -1 // ... 2го символа
		counter int
	)
	for _, value := range str {
		if value == symb1 {
			indexSymb1 = counter
		}
		if value == symb2 {
			indexSymb2 = counter
		}
		if indexSymb1 != -1 && indexSymb2 != -1 {
			tmpDistance := abs(indexSymb1 - indexSymb2) - 1
			if tmpDistance < distance {
				distance = tmpDistance
			}
		}
		counter++
	}
	return 
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}