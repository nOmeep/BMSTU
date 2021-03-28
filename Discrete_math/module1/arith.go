package main

/*
	Здесь <number> и <var> обозначают множества терминальных символов, соответствующих числам и именам переменных.

	После удаления левой рекурсии получается LL(1)-грамматика
	<E>  ::= <T> <E’>.
	<E’> ::= + <T> <E’> | - <T> <E’> | .
	<T>  ::= <F> <T’>.
	<T’> ::= * <F> <T’> | / <F> <T’> | .
	<F>  ::= <number> | <var> | ( <E> ) | - <F>.

	Алгоритм вычисления значения выражения должен быть разбит на две части: лексический анализатор и синтаксический анализатор.
*/

import (
	"fmt" 
	"strconv" // для atoi
	"github.com/skorobogatov/input" // для Gets()
)

// Реализую стек для хранения чисел и снятия их с верхушки 

type Stack []int

func (s *Stack) isEmpty() bool { // проверка на пустоту стека во избежание ошибок 
	return len(*s) == 0 // если стек пустой, будет "panic"
}

func (s *Stack) push(x int) { // добавить элемент на верхушку стека
	*s = append(*s, x) // добавляю новое значение в конец стека
}

func (s *Stack) pop() int { // снять элемент с верхушки стека
	if s.isEmpty() {
		fmt.Println("Panic in pop()")
		return 0
	} else {
		var index int = len(*s) - 1 // берем индекс у самого верхнего элемента
		var topElement int = (*s)[index] // достаем элемент по индексу выше
		*s = (*s)[:index] // не забываем удалять элемент из стека
		return topElement
	}
}

var stack Stack // экземпляр stack'a

// конец реализации stack'a

// необходимые для задачи структуры
type Tag int

type Lexem struct { 
    Tag 
    Image string 
} 

type Variable struct { // для переменных их значений
	value int
	title string
}

// конец описания необходимых структур

//начало ввода global переменных для парсинига

const ( 
    ERROR Tag = 1 << iota  // Неправильная лексема 
    NUMBER                 // Целое число 
    VAR                    // Имя переменной 
    PLUS                   // Знак + 
    MINUS                  // Знак - 
    MUL                    // Знак * 
    DIV                    // Знак / 
    LPAREN                 // Левая круглая скобка 
    RPAREN                 // Правая круглая скобка 
) 

var (
	lexem []Lexem
	lexemLength int 
	lexemIndex int 

	token []string

	variable []Variable

	errorStatus bool
)

// Конец global переменных для парсинга

// Основные функции парсинга

func NumericValue(curTerm string, x int) (result int){
	var (
		index int
	)

	for index = x; index < len(curTerm); index++ {
		if (curTerm[index] > 47 && curTerm[index] < 58) { // если цифра - все GOOD, иначе прекращаем
			continue 
		} else {
			break 
		}
	}

	result = index
	return
}	

func VarIndex(curTerm string, x int) (result int) {
	var (
		index int
	)

	for index = x; index < len(curTerm); index++ {
		if (curTerm[index] > 47 && curTerm[index] < 58) || (curTerm[index] > 64 && curTerm[index] < 91) || (curTerm[index] > 96 && curTerm[index] < 123) { // a-z; A-Z; 0-9
			continue
		} else { // если не 0-9, a-z или A-Z
			break 
		}
	}

	result = index
	return
}

func lexer(expr string) {
	var (
		lx Lexem
	)

	for i := 0; i < len(expr); i++ {

		if (expr[i] == 32) { // пробел 

			continue

		} else if (expr[i] == 40) { // знак "("

			lx.Tag = LPAREN
			lx.Image = "("
			lexem = append(lexem, lx)

		} else if (expr[i] == 41) { // знак ")"

			lx.Tag = RPAREN
			lx.Image = ")"
			lexem = append(lexem, lx)

		} else if (expr[i] == 42) { // знак "*"

			lx.Tag = MUL
			lx.Image = "*"
			lexem = append(lexem, lx)

		} else if (expr[i] == 43) { // знак "+"

			lx.Tag = PLUS
			lx.Image = "+"
			lexem = append(lexem, lx)

		} else if (expr[i] == 45) { // знак "-"

			lx.Tag = MINUS
			lx.Image = "-"
			lexem = append(lexem, lx)

		} else if (expr[i] == 47) { // знак "/"

			lx.Tag = DIV
			lx.Image = "/"
			lexem = append(lexem, lx)

		} else { // значение, имя или ERROR

			if (expr[i] > 47 && expr[i] < 58) {

				var (
					valueIndex int = NumericValue(expr, i)
				)

				lx.Tag = NUMBER
				lx.Image = expr[i:valueIndex]

				lexem = append(lexem, lx)

				i = valueIndex - 1

			} else if (expr[i] > 64 && expr[i] < 91) || (expr[i] > 96 && expr[i] < 123) {

				var (
					variableIndex int = VarIndex(expr, i)
				)

				lx.Tag = VAR
				lx.Image = expr[i: variableIndex]

				lexem = append(lexem, lx)

				i = variableIndex - 1

			} else {

				lx.Tag = ERROR
				lexem = append(lexem, lx)

				errorStatus = true

			}

		}

		lexemLength += 1
	} 
}

func EParse() {
	TParse()
	ECommaParse()
}

func TParse() {
	FParse()
	TCommaParse()
}

func TCommaParse() {
	if lexemLength > lexemIndex {

		var (
			lx = lexem[lexemIndex]
		)

		if lx.Tag&(DIV|MUL) != 0 {
			lexemIndex += 1

			FParse()

			token = append(token, lx.Image)

			TCommaParse()
		}

	}
}

func ECommaParse() {
	if lexemLength > lexemIndex {

		var (
			lx = lexem[lexemIndex]
		)

		if lx.Tag&(PLUS|MINUS) != 0 {
			lexemIndex += 1

			TParse()

			token = append(token, lx.Image)

			ECommaParse()

		} else if lx.Tag&(VAR|NUMBER) != 0 {

			errorStatus = true

		}

	}
}

func FParse() {
	if lexemLength > lexemIndex {

		var (
			lx = lexem[lexemIndex]
		)

		if lx.Tag&(NUMBER|VAR) != 0 {
			lexemIndex += 1
			token = append(token, lx.Image)
		} else if lx.Tag&MINUS != 0 {
			lexemIndex += 1
			token = append(token, "-1")
			FParse()
			token = append(token, "*")
		} else if lx.Tag&LPAREN != 0 {
			lexemIndex += 1
			EParse()
			if lexemLength > lexemIndex {
				lx = lexem[lexemIndex]
				lexemIndex += 1
				if lx.Tag&RPAREN == 0 {
					errorStatus = true
				}
			} else {
				errorStatus = true
			}
		} else {
			errorStatus = true
		}
	} else {
		errorStatus = true
	}
}

func getVal(term string) int {
	for _, curVar := range variable {
		if curVar.title == term {
			return curVar.value
		}
	}
	return 0
}

func Vocabular(term string) int {
	for _, curVar := range variable {
		if curVar.title == term {
			return 1
		}
	}
	return 0
}

func SchemeEval() int {
	for _, tmp := range token{
		if (tmp[0] > 47 && tmp[0] < 58) || (tmp[0] == '-' && len(tmp) > 1){
			num, _ := strconv.Atoi(tmp)
			stack.push(num)
		} else if ( tmp[0] > 64 &&  tmp[0] < 91) || ( tmp[0] > 96 && tmp[0] < 123) {
			num := getVal(tmp)
			stack.push(num)
		} else {
			switch tmp {

				case "+":

					stack.push(stack.pop() + stack.pop())

				case "-":

					var (
						second int = stack.pop()
						first int = stack.pop()
					)

					stack.push(first - second)

				case "*":

					stack.push(stack.pop() * stack.pop())

				case "/":

					var (
						second int = stack.pop()
						first int = stack.pop()
					)

					if first == 0 {

						stack.push(0)

					} else {

						stack.push(first/second)
					}

			}
		}
	}
	return stack.pop()
}

// Конец функций

// Начало gachibass'a в main'e

func main(){
	var (
		term string = input.Gets()
	)
	lexer(term)
	EParse()
	lexemIndex = len(token) - 1
	if errorStatus {
		fmt.Println("error")
	} else {
		var x int
		var v Variable
		for _, lx := range lexem{
			if lx.Tag & VAR != 0 && Vocabular(lx.Image) == 0{
				input.Scanf("%d", &x)
				v.title = lx.Image
				v.value = x
				variable = append(variable, v)
			}
		}
		fmt.Println(SchemeEval())
		//fmt.Println(errorStatus)
	}
}

// конец всего...

