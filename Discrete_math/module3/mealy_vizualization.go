package main

import "fmt"

func input_matrix_scan(M *[][]int64, n int64, m int64) {
	for i := 0; int64(i) < n; i++ {
		for j := 0; int64(j) < m; j++ {
			fmt.Scan(&(*M)[i][j])
		}
	}
}

func out_matrix_scan(M *[][]string, n int64, m int64) {
	for i := 0; int64(i) < n; i++ {
		for j := 0; int64(j) < m; j++ {
			fmt.Scan(&(*M)[i][j])
		}
	}
}

func draw_graph(matrix_in *[][]int64, matrix_out *[][]string, n int64, m int64, q0 int64) {
	fmt.Println("digraph {")
	fmt.Println("\trankdir = LR")
	fmt.Println("\tdummy [label = \"\", shape = none]")
	for i := 0; int64(i) < n; i++ {
		fmt.Printf("\t%d [shape = circle]\n", i)
	}
	fmt.Println("\tdummy ->", q0)
	for i := 0; int64(i) < n; i++ {
		var (
			symbol byte = 97;
		)
		for j := 0; int64(j) < m; j++ {
			fmt.Printf("\t%d -> %d [label = \"%c(%s)\"]", i, (*matrix_in)[i][j], symbol, (*matrix_out)[i][j])
			fmt.Println();
			symbol++
		}
	}
	fmt.Println("}")
}

func main(){
	var (
		n int64; // количество состояний автомата
		m int64; // размер входного алфавита
		q0 int64; // номер начального состояния 
	)
	
	fmt.Scan(&n);
	fmt.Scan(&m);
	fmt.Scan(&q0);

	var (
		matrix_in [][]int64 = make([][]int64, n)
		matrix_out [][]string = make([][]string, n);
	)

	for i := 0; int64(i) < n; i++ {
		matrix_in[i] = make([]int64, m);
		matrix_out[i] = make([]string, m);
	}
	
	input_matrix_scan(&matrix_in, n, m);
	out_matrix_scan(&matrix_out, n, m);

	draw_graph(&matrix_in, &matrix_out, n, m, q0);
}