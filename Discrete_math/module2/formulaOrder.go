package main

import (
	"fmt"
	"os"
	"sort"
	"bufio"
	"strings"
)

var (
	need int = 0
)

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

type Lexem struct {
	Tag
	Image string
}

type Tag int64
const (
	ERROR Tag = 1 << iota
	NUMBER
	VAR
	PLUS
	MINUS
	MUL
	DIV
	LPAREN
	RPAREN
)

type Vertex struct {
	formula string
	color string
	output string

	edges []int64
	exit int64
}

func numericValue(expr string, ind int64) (result int64) {
	var (
		index int64
	)

	for index = ind; index < int64(len(expr)); index++ {
		if (expr[index] > 47 && expr[index] < 58) { // если цифра - все GOOD, иначе прекращаем
			continue 
		} else {
			break 
		}
	}

	result = index
	return

}

func variableIndex(expr string, ind int64) (result int64) {
	var (
		index int64
	)

	for index = ind; index < int64(len(expr)); index++ {
		if (expr[index] > 47 && expr[index] < 58) || (expr[index] > 64 && expr[index] < 91) || (expr[index] > 96 && expr[index] < 123) { // a-z; A-Z; 0-9
			continue
		} else { // если не 0-9, a-z или A-Z
			break 
		}
	}

	result = index
	return
}

func lexer(expr string, lexems *[]Lexem) {
	var (
		lx Lexem
		i int64
	)

	for ; i < int64(len(expr)); i++ {

		if (expr[i] == 32) { // пробел 

			continue

		} else if (expr[i] == 40) { // знак "("

			lx.Tag = LPAREN
			lx.Image = "("
			*lexems = append(*lexems, lx)

		} else if (expr[i] == 41) { // знак ")"

			lx.Tag = RPAREN
			lx.Image = ")"
			*lexems = append(*lexems, lx)

		} else if (expr[i] == 42) { // знак "*"

			lx.Tag = MUL
			lx.Image = "*"
			*lexems = append(*lexems, lx)

		} else if (expr[i] == 43) { // знак "+"

			lx.Tag = PLUS
			lx.Image = "+"
			*lexems = append(*lexems, lx)

		} else if (expr[i] == 45) { // знак "-"

			lx.Tag = MINUS
			lx.Image = "-"
			*lexems = append(*lexems, lx)

		} else if (expr[i] == 47) { // знак "/"

			lx.Tag = DIV
			lx.Image = "/"
			*lexems = append(*lexems, lx)

		} else { // значение, имя или ERROR

			if (expr[i] > 47 && expr[i] < 58) {

				var (
					valueIndex int64 = numericValue(expr, int64(i))
				)

				lx.Tag = NUMBER
				lx.Image = expr[i:valueIndex]

				*lexems = append(*lexems, lx)

				i = valueIndex - 1

			} else if (expr[i] > 64 && expr[i] < 91) || (expr[i] > 96 && expr[i] < 123) {

				var (
					variableIndex int64 = variableIndex(expr, i)
				)

				lx.Tag = VAR
				lx.Image = expr[i: variableIndex]

				*lexems = append(*lexems, lx)

				i = variableIndex - 1

			} else {

				lx.Tag = ERROR
				*lexems = append(*lexems, lx)

			}

		}
	} 
}

func E_Parse(index *int64, c int64, lexArr []Lexem, V_G *[]Vertex)  {
	T_Parse(index, c, lexArr, V_G)
	E_Comma_Parse(index,  c, lexArr, V_G)
}

func E_Comma_Parse(index *int64, c int64, lexArr []Lexem, V_G *[]Vertex)  {
	var lx Lexem
	if *index < int64(len(lexArr)) {
		lx = lexArr[*index]
	}
	if lx.Tag & PLUS != 0 {
		*index++
		T_Parse(index, c, lexArr, V_G)
		E_Comma_Parse(index, c, lexArr, V_G)

	}
	if lx.Tag & MINUS != 0 {
		*index++
		T_Parse(index, c, lexArr, V_G)
		E_Comma_Parse(index, c, lexArr, V_G)
	}
	if lx.Tag & (VAR | NUMBER | ERROR) != 0 {
		//fmt.Printf("Count %d", need)
		need++
		makeAnError()
	}
}

func T_Parse(index *int64, c int64, lexArr []Lexem, V_G *[]Vertex)  {
	F_Parse(index, c ,lexArr, V_G)
	T_Comma_Parse(index, c ,lexArr, V_G)
}

func T_Comma_Parse(index *int64, c int64, lexArr []Lexem, V_G *[]Vertex)  {
	var lx Lexem
	if *index < int64(len(lexArr)) {
		lx = lexArr[*index]
	}
	if lx.Tag & DIV != 0 {
		*index++
		F_Parse(index, c, lexArr, V_G)
		T_Comma_Parse(index, c, lexArr, V_G)

	}
	if lx.Tag & MUL != 0 {
		*index++
		F_Parse(index, c, lexArr, V_G)
		T_Comma_Parse(index, c, lexArr, V_G)
	}
	if lx.Tag & (VAR | NUMBER | ERROR) != 0 {
		//fmt.Printf("Count %d", need)
		need++
		makeAnError()
	}

}

func F_Parse(index *int64, c int64, lexArr []Lexem, V_G *[]Vertex)  {
	var lx Lexem
	if *index < int64(len(lexArr)) {
		lx = lexArr[int(*index)]
	} else {
		//fmt.Printf("Count IT WAS HERE %d", need)
		need++
		makeAnError()
	}

	if lx.Tag & (ERROR | RPAREN | PLUS | DIV | MUL | ERROR)!= 0 {
		//fmt.Printf("Count %d", need)
		need++
		makeAnError()
	}
	if lx.Tag & VAR != 0 {
		*index++
		n, ok := dict[lx.Image]
		if !ok {
			//fmt.Printf("Count %d", need)
			need++
			makeAnError()
		}

		var (
			flag bool
		)
		for _, x := range (*V_G)[n].edges {
			if x == c {
				flag = true
			}
		}
		if !flag {
			(*V_G)[n].edges = append((*V_G)[n].edges, c)
		}
	}
	if lx.Tag & NUMBER != 0 {
		*index++
	}
	if lx.Tag & MINUS != 0 {
		*index++
		F_Parse(index, c ,lexArr, V_G)
	}
	if lx.Tag & LPAREN != 0 {
		*index++
		E_Parse(index, c, lexArr, V_G)
		if *index < int64(len(lexArr)) {
			lx = lexArr[*index]
			*index++
		}
		if lx.Tag & RPAREN == 0{
			//fmt.Printf("Count %d", need)
			need++
			makeAnError()
		}
	}

}

func doScan(scan *bufio.Scanner, V_G *[]Vertex) {
	for scan.Scan() {
		var (
			formula string = scan.Text()
		)
		if strings.Compare(formula, "") == 0 {
			break
		}
		var (
			out string = formula
		)
		formula = strings.Replace(formula, " ", "", -1)
		if strings.Index(formula, "=") == -1 {
			//fmt.Printf("Count %d", need)
			need++
			makeAnError()
		}
		var (
			left string = formula[:strings.Index(formula, "=")]
			right string = formula[strings.Index(formula, "=") + 1:]
		)
		if left == "" || right == "" {
			//fmt.Printf("Count %d", need)
			need++
			makeAnError()
		}
		var (
			variable, varValue int64
		)
		for strings.Index(left, ",") != -1 || strings.Index(right, ",") != -1{
			if strings.Index(left, ",") != -1 {
				variable++
			}
			if strings.Index(right, ",") != -1 {
				varValue++
			}
			if strings.Index(left, ",") == -1 {
				break
			} else {
				var (
					tmp string = left[:strings.Index(left, ",")]
				)
				if _, ok := dict[tmp]; ok {
					//fmt.Printf("Count %d", need)
					need++
					makeAnError()
				}
				if (tmp[0] < 65 || tmp[0] > 90) && (tmp[0] < 97 || tmp[0] > 122) {
					//fmt.Printf("Count %d", need)
					need++
					makeAnError()
				}
				for i := 1; i < len(tmp); i++ {
					if (tmp[i] < 65 || tmp[i] > 90) && (tmp[i] < 97 || tmp[i] > 122) && (tmp[i] < 48 || tmp[i] > 57) {
						//fmt.Printf("Count %d", need)
						need++
						makeAnError()
					}
				}
				dict[tmp] = count
				left = left[strings.Index(left, ",") + 1:]
				right = right[strings.Index(right, ",") + 1:]
			}
		}
		if variable != varValue {
			//fmt.Printf("Count %d", need)
			need++
			makeAnError()
		} else if _, flag := dict[left]; flag {
			//fmt.Printf("Count%d", need)
			need++
			makeAnError()
		} else {
			if (left[0] < 65 || left[0] > 90) && (left[0] < 97 || left[0] > 122) {
				//fmt.Printf("Count %d", need)
				need++
				makeAnError()
			}
			for i := 1; i < len(left); i++ {
				if (left[i] < 65 || left[i] > 90) && (left[i] < 97 || left[i] > 122) && (left[i] < 48 || left[i] > 57) {
					//fmt.Printf("Count%d", need)
					need++
					makeAnError()
				}
			}
			dict[left] = count
			if variable != varValue {
				//fmt.Printf("Count %d", need)
				need++
				makeAnError()
			} else {
				left = formula[:strings.Index(formula, "=")]
				addVertex(V_G, out, formula)
				count++
			}
		}
	}
}

func DFS(V_G *[]Vertex, time *int)  {
	for i := range *V_G {
		if strings.Compare((*V_G)[i].color, "black") == 0 {
			visitVertex(V_G, int64(i), time)
		}
		if strings.Compare((*V_G)[i].color, "gray") == 0 {
			makeCycle()
		}
	}
}

func makeDFS(v_g *[]Vertex) {
	var (
		time int
	)
	DFS(v_g, &time)
}

func DFSandSort(v_g *[]Vertex) {
	makeDFS(v_g)
	sort.Slice(*v_g, func(i, j int) bool {
		return int((*v_g)[int(i)].exit) > int((*v_g)[j].exit)
	})
}

func visitVertex(V_G *[]Vertex, v int64, time *int)  {
	(*V_G)[v].color = "gray"
	for _, e := range (*V_G)[v].edges {
		if strings.Compare((*V_G)[e].color, "black") == 0 {
			visitVertex(V_G, e, time)
		}
		if strings.Compare((*V_G)[e].color, "gray") == 0 {
			makeCycle()
		}
	}
	(*V_G)[v].exit = int64(*time)
	*time++
	(*V_G)[v].color = "white"
}

var (
	count int64
	люблюпарсеры bool = true 
)

func makeAnError() {
	fmt.Println("syntax error")
	os.Exit(2)
}

func makeCycle() {
	fmt.Println("cycle")
	os.Exit(2)
}

func addVertex(v_g *[]Vertex, out string, formula string) {
	var tmp Vertex
	tmp.color = "black"
	tmp.formula = formula
	tmp.edges = make([]int64, 0)
	tmp.output = out
	*v_g = append(*v_g, tmp)
}

var (
	dict = make(map[string]int64)
)

func main()  {
	var (
		V_G []Vertex = make([]Vertex, 0)
		scan *bufio.Scanner = bufio.NewScanner(os.Stdin)
	)
	doScan(scan, &V_G)
	for i, _ := range V_G {
		tmp := V_G[i].output[strings.Index(V_G[i].output, "=") + 1:]
		for strings.Index(tmp, ",") != -1{
			var (
				lexArr = make([]Lexem, 0)
			)
			lexer(tmp[:strings.Index(tmp, ",")], &lexArr)
			var (
				index int64
			)
			E_Parse(&index, int64(i), lexArr, &V_G)
			tmp = tmp[strings.Index(tmp, ",") + 1:]
		}
		var (
			lexArr = make([]Lexem, 0)
			index int64
		)
		lexer(tmp, &lexArr)
		E_Parse(&index, int64(i), lexArr, &V_G)
	}

	DFSandSort(&V_G)
	for _, v := range V_G {
		fmt.Println(v.output)
	}

	if люблюпарсеры {
		os.Exit(3)
	}

	fmt.Println("Как")
	fmt.Println("же")
	fmt.Println("я")
	fmt.Println("люблю")
	fmt.Println("парсеры")
}