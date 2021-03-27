package main

import (
	"fmt"
	"os"
)

// решил отказаться от хранения числителя и знамеенателя в отдельных матрицах
// просто сделал структуру "дробь", т.к скорее всего ошибаюсь где-то в вычислениях
type fraction struct {
	numerator int
	denominator int
}
//

func main(){
	// ввод переменных и заполнение матриц
	var (
		nel int
	)
	fmt.Scanf("%d", &nel)
	var (
		matrix [][]fraction = make([][]fraction, nel);
		answer []fraction = make([]fraction, nel)
	)
	for i := 0; i < nel; i++ {
		matrix[i] = make([]fraction, nel + 1)
	}
	//

	//заполенение матрицы по числителям, знаменатель по умолчанию 1!
	for i := 0; i < nel; i++ {
		for j := 0; j < nel + 1; j++ {
			fmt.Scan(&matrix[i][j].numerator)
			matrix[i][j].denominator = 1
		}
	}
	//

	// вынес часто повторяющиеся действия из прошлого решения в отдельные функции, чтобы было проще отслеживать ошибки

	// старый НОД почему-то в некторых тестах делил на 0
	gcd := func(a int, b int) int {
		for b != 0 {
			t := b
			b = a % b
			a = t
		}
		return a
	}

	createCurrentFrac := func(a int, b int) (result fraction) {
		result.numerator = a
		result.denominator = b

		return
	}

	generalize := func(a fraction, b fraction, position int) (result fraction) {
		if position == 1 {
			result.numerator = a.numerator * b.denominator
			result.denominator = b.numerator * a.denominator
		} else if position == 2 {
			result.numerator = a.numerator * b.numerator
			result.denominator = a.denominator * b.denominator
		} else if position == 3 {
			result.numerator = a.denominator * b.numerator + b.denominator * a.numerator
			result.denominator = a.denominator * b.denominator
		}

		return
	}

	simplify := func(a fraction, b fraction) (result fraction) {
		result = generalize(a, b, 1)

		var localGcd int = gcd(result.numerator, result.denominator)

		result.numerator /= localGcd
		result.denominator /= localGcd

		return
	}

	multiplication := func(a fraction, b fraction) (result fraction) {
		result = generalize(a, b, 2)

		var localGcd int = gcd(result.numerator, result.denominator)

		result.numerator /= localGcd
		result.denominator /= localGcd

		return
	}

	sumfraction := func(a fraction, b fraction) (result fraction) {
		result = generalize(a, b, 3)

		var localGcd int = gcd(result.numerator, result.denominator)

		result.numerator /= localGcd
		result.denominator /= localGcd

		return
	}
	/////

	//преобразования
	var (
		ti int
		tj int
	)
	for ti + 1 < nel && tj < nel + 1{
		if !(matrix[ti][tj].numerator == 0 && matrix[ti][tj].denominator == 1) {
			var tmp fraction = createCurrentFrac(-matrix[ti][tj].numerator, matrix[ti][tj].denominator)
	
			for i := ti + 1; i < nel; i++ {
				var help fraction = simplify(matrix[i][tj], tmp)
				for j := tj; j < nel + 1; j++ {
					matrix[i][j] = sumfraction(matrix[i][j], multiplication(matrix[ti][j], help))
				}
			}
			ti += 1
			tj += 1
		} else {
			flagBool := false

			var str int = -1

			for i := ti; i < nel; i++ {
				if !(matrix[i][tj].numerator == 0 && matrix[i][tj].denominator == 1) {
					str = i
					flagBool = true
					break
				}
			}

			if flagBool == true {
				for i := 0; i < nel + 1; i++ {
					matrix[ti][i], matrix[str][i] = matrix[str][i], matrix[ti][i]
				}
			} else {
				tj += 1
			}
		}
	}

	for i := 0; i < nel; i++ {
		flagBool := false 

		for j := 0; j < nel; j++ {
			if flagBool || !(matrix[i][j].numerator == 0 && matrix[i][j].denominator == 1) {
				flagBool = true
			} else {
				flagBool = false
			}
		}
		if flagBool == false {
			fmt.Println("No solution")
			os.Exit(0) 
		}
	}
	for i := nel - 1; i > -1; i-- {
		for j := nel - 1; j > i; j-- {
			var tmp fraction = createCurrentFrac(-matrix[i][j].numerator, matrix[i][j].denominator)

			matrix[i][nel] = sumfraction(matrix[i][nel], multiplication(answer[j], tmp))
		}
		if matrix[i][i].numerator == 0 && matrix[i][i].denominator == 1 {
			fmt.Println("No solution")
			os.Exit(0)
		}
		answer[i] = simplify(matrix[i][nel], matrix[i][i])
	}

	// без функции ниже будет иногда неправильный ответ из-за знака дроби
	resultingAnswer := func(x fraction) (result fraction) {
		if (x.numerator > -1 && x.denominator > 0) || (x.numerator < 1 && x.denominator > 0) {
			result.numerator = x.numerator
			result.denominator = x.denominator
		}
		
		if (x.numerator > -1 && x.denominator < 0) || (x.numerator < 1 && x.denominator < 0) {
			result.numerator = -x.numerator
			result.denominator = -x.denominator
		}

		return
	}
	for i := 0; i < nel; i++ {
		answer[i] = resultingAnswer(answer[i])
		fmt.Printf("%d/%d\n", answer[i].numerator, answer[i].denominator)
	}
}