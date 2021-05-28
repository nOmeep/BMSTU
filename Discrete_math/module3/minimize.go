package main

import (
	"fmt"
)

var (
	time int64
)

func main() {
        var (
			n int64
			m int64 
			q_start int64
		)
		fmt.Scan(&n)
		fmt.Scan(&m)
		fmt.Scan(&q_start)

		var (
			delta [][]int64 = make([][]int64, n)
			phi [][]string = make([][]string, n)
		)

		global_arr1 = make([]int64, n)
		global_arr2 = make([]int64, 0, n)

        for i := 0; int64(i) < n; i++ { 
			delta[i] = make([]int64, m)
			phi[i] = make([]string, m)
			global_arr1[i] = -1
		}
        for i := 0; int64(i) < n; i++ { 
			for j := 0; int64(j) < m; j++ { 
				fmt.Scan(&delta[i][j]) 
			} 
		}
		for i := 0; int64(i) < n; i++ { 
			for j := 0; int64(j) < m; j++ { 
				fmt.Scan(&phi[i][j]) 
			} 
		}

        var (
			pi []int64 = make([]int64, 0, m)
		)

        for i := range delta { 
			pi = append(pi, int64(i)) 
		}

		for i:= range delta {
			for j := range delta {
				if pi[i] != pi[j] {
					var (
						is_equal bool = true
					)
					for k := 0; k < len(delta[0]); k++ {
						if phi[i][k] != phi[j][k] {
							is_equal = false
							break
						}
					}
				if is_equal { 
					for q := 0; q < len(pi); q++ { 
						if pi[q] == pi[i] { 
							pi[q] = pi[j] 
						} 
					}
					m-- 
				}
			}
		}
	}

	for {
        tmp_m, tmp_pi := Split(&delta, &phi, &pi, q_start)
		if m == tmp_m { 
			break 
		}
		m = tmp_m
		pi = tmp_pi
	}

	var (
		h3 map[int64]int64 = make(map[int64]int64, len(pi))
	)
	
	for i, x := range pi {
		_, tmp_flag := h3[x]
		if tmp_flag { 
			pi[i] = h3[x]
		} else { 
			pi[i] = int64(len(h3))
			h3[x] = int64(len(h3)) 
		} 
	}

	var (
		tmpD [][]int64 = make([][]int64, m)
		tmpF [][]string = make([][]string, m)
		q01 int64
	)

	for i := 0; int64(i) < m; i++ { 
		tmpD[i] = make([]int64, len(delta[q_start]))
		tmpF[i] = make([]string, len(delta[q_start]))
	}

	for i := 0; i < len(delta); i++ {
		var (
			q int64 = pi[i]
		)
		if int64(i) == q_start { 
			q01 = q 
		}
		for j := 0; j < len(delta[q_start]); j++ { 
			tmpD[q][j] = pi[delta[i][j]]
			tmpF[q][j] = phi[i][j] 
		}
	}
	
    do_DFS(&tmpD, &tmpF, q01)

	var (
		resultD [][]int64 = make([][]int64, time)
		resultF [][]string = make([][]string, time)
		flag bool
	)

	for i := 0; int64(i) < time; i++ { 
		resultD[i] = make([]int64, len(tmpD[0]))
		resultF[i] = make([]string, len(tmpD[0]))
	}
	for i := 0; int64(i) < time; i++ {
        flag = true
		
		var (
			x int64 = global_arr2[i]
			curr1 int64 = -1
			curr2 string = "x"
		)
		for j := 0; j < len(tmpD[x]); j++ {
			curr1 = tmpD[x][j]
			resultD[i][j] = global_arr1[curr1]
		}
		for j := 0; j < len(tmpF[x]); j++ {
			curr2 = tmpF[x][j]
			resultF[i][j] = curr2
		}
	}

	if flag == true { 
		fmt.Println("digraph {")
		fmt.Println("\trankdir = LR")
		fmt.Println("\tdummy [label = \"\", shape = none]")
		for i := 0; i < len(resultD); i++ { 
			fmt.Print("\t", i, " [shape = circle]")
			fmt.Println()
		}
		fmt.Println("\tdummy ->", "0")
		for i := 0; i < len(resultD); i++ { 
			for j := 0; j < len(resultF[0]); j++ { 
					var tmp_value = 'a'
					fmt.Print("\t", i, " -> ", resultD[i][j], " [label = \"")
					fmt.Printf("%c", int(tmp_value) + j)
					fmt.Print("(", resultF[i][j], ")\"]")
					fmt.Println()
			} 
		}
		fmt.Printf("}") 
	}
}

var (
	global_arr1 []int64
	global_arr2 []int64
)

func increment(i *int64) {
	*i++
}

func Union(x int64, y int64, z int64, pi *[]int64) {
	for z < int64(len(*pi)) { 
		if (*pi)[z] == x { 
			(*pi)[z] = y 
		}
		z++ 
	}
}

func Split(delta *[][]int64, phi *[][]string, pi *[]int64, q0 int64) (int64, []int64) {
    var (
		m_classes int64 = int64(len(*delta))
    	pi1 []int64 = make([]int64, 0, m_classes)
	)
	for i, _ := range *delta { 
		pi1 = append(pi1, int64(i)) 
	}
	for i := range *delta {
		for j := range *delta {
			if (*pi)[i] == (*pi)[j] && pi1[i] != pi1[j] {
				eq := true
				for k := 0; k < len((*delta)[q0]); k++ {
					w1, w2 := (*delta)[i][k], (*delta)[j][k]
					if (*pi)[w1] != (*pi)[w2] {
						eq = false
						break
					}
				}
				if eq {
					Union(pi1[i], pi1[j], 0, &pi1)
					m_classes--
				}
			}
		}
	}
	return m_classes, pi1
}

func do_DFS(delta *[][]int64, phi *[][]string, q int64) {
	DFS(delta, phi, q)
}

func DFS(delta *[][]int64, phi *[][]string, q int64) {
	if 0 > global_arr1[q] { 
			global_arr1[q] = time
			increment(&time)
			global_arr2 = append(global_arr2, q)
			for i := 0; i < len((*delta)[q]); i++ {
				do_DFS(delta, phi, (*delta)[q][i]) 
			} 
        }
}