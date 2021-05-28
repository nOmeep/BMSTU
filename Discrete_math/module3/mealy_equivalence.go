package main

import (
 	"fmt"
)

var (
	time int
	keepGoing bool = true
)

func do_scan(n2 int, m2 int, Delt2 *[][]int, Phi2 *[][]string) {
	for i := 0; i < n2; increment(&i) { 
		(*Delt2)[i] = make([]int, m2)
		(*Phi2)[i] = make([]string, m2)
		global_arr1[i] = -1 
	}
	for i := 0; i < n2; increment(&i) { 
		for j := 0; j < m2; increment(&j) { 
			fmt.Scan(&(*Delt2)[i][j]) 
		} 
	}
	for i := 0; i < n2; increment(&i) { 
		for j := 0; j < m2; increment(&j) { 
			fmt.Scan(&(*Phi2)[i][j]) 
		} 
	}
}

var (
	global_arr1 []int
	global_arr2 []int
)

func main() {
	var (
		n2 int
		m2 int
		q_start2 int
		mark_Delt int
		mark_Phi int
	)

	fmt.Scan(&n2)
	fmt.Scan(&m2)
	fmt.Scan(&q_start2)

	var (
		Delt2 [][]int = make([][]int, n2) 
		Phi2 [][]string =  make([][]string, n2)
	)

	global_arr1 = make([]int, n2)
	global_arr2 = make([]int, 0, n2)

	do_scan(n2, m2, &Delt2, &Phi2)

	var (
		minDelt2 [][]int
		minPhi2 [][]string
	)

	var (
		m int = len(Delt2)
	)
	var (
		pi []int = make([]int, 0, m)
	)
	for i := range Delt2 { 
		pi = append(pi, i) 
	}
	for i := range Delt2 {
		for j := range Delt2 {
			if pi[i] != pi[j] {
				var (
					flag bool = true
				)
				for k := 0; k < len(Delt2[0]); increment(&k) {
					if Phi2[i][k] != Phi2[j][k] { //////////////////////////////////
						flag = false
						break
					}
				}
				if flag {
					for q := 0; q < len(pi); increment(&q) { 
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
		var (
			m1 int
			pi1 []int
		)
		m1, pi1 = Split(&Delt2, &Phi2, &pi, q_start2)
		if m == m1 { 
			break 
		}
		m = m1
		pi = pi1
	}

	var (
		h3 map[int]int = make(map[int]int, len(pi))
	)
	
	for i, x := range pi {
		if _, flag := h3[x]; flag { 
			pi[i] = h3[x] 
		} else { 
			var (
				length int = len(h3)
			)
			pi[i] = length
			h3[x] = length 
		}
	}

	var (
		Delt1 [][]int = make([][]int, m)
		Phi1 [][]string = make([][]string, m)
		q01 int
	)

	for i := 0; i < m; increment(&i) { 
		Delt1[i] = make([]int, len(Delt2[q_start2]))
		Phi1[i] = make([]string, len(Delt2[q_start2]))
	}
	
	for i := 0; i < len(Delt2); increment(&i) {
		var (
			tmp_q int = pi[i]
		)
		if i == q_start2 { 
			q01 = tmp_q 
		}
		for j := 0; j < len(Delt2[q_start2]); increment(&j) { 
			Delt1[tmp_q][j] = pi[Delt2[i][j]]
			Phi1[tmp_q][j] = Phi2[i][j] 
		}
	}

	DFS(&Delt1, &Phi1, q01)

	var (
		new_Delt [][]int
		new_Phi [][]string
	)

	var (
		tmp_Delt1 [][]int = make([][]int, time)
		tmp_Phi1 [][]string = make([][]string, time)
		tmp_yes bool
	)

	for i := 0; i < time; increment(&i) { 
		tmp_Delt1[i] = make([]int, len(Delt1[0])) 
		tmp_Phi1[i] = make([]string, len(Delt1[0])) 
	}
	for i := 0; i < time; increment(&i) {
		tmp_yes = true
		var (
			x int = global_arr2[i]
			curr1 int = -1
			curr2 string = "x"
		)
		for j := 0; j < len(Delt1[x]); increment(&j) {
			curr1 = Delt1[x][j]
			tmp_Delt1[i][j] = global_arr1[curr1]
		}
		for j := 0; j < len(Phi1[x]); increment(&j) {
			curr2 = Phi1[x][j]
			tmp_Phi1[i][j] = curr2
		}
	}
	new_Delt = tmp_Delt1
	new_Phi = tmp_Phi1
	
	if tmp_yes { 
		minDelt2 = new_Delt
		minPhi2 = new_Phi
	} else {
		minDelt2 = nil
		minPhi2 = nil
	}

	var (
		n3 int
		m3 int
		q_start3 int
	)

	fmt.Scan(&n3)
	fmt.Scan(&m3)
	fmt.Scan(&q_start3)

	var (
		Delt3 [][]int = make([][]int, n3) 
		Phi3 [][]string =  make([][]string, n3)
	)

	global_arr1 = make([]int, n3)
	global_arr2 = make([]int, 0, n3)
	time = 0

	do_scan(n3, m3, &Delt3, &Phi3)

	var (
		minDelt3 [][]int
		minPhi3 [][]string
	)

	m = len(Delt3)
	pi = make([]int, 0, m)
	for i := range Delt3 { 
		pi = append(pi, i) 
	}
	for i := range Delt3 {
		for j := range Delt3 {
			if pi[i] != pi[j] {
				var (
					flag bool = true
				)
				for k := 0; k < len(Delt3[0]); increment(&k) {
					if Phi3[i][k] != Phi3[j][k] {
						flag = false
						break
					}
				}
				if flag {
					for q := 0; q < len(pi); increment(&q) { 
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
		var (
			m1 int
			pi1 []int
		)
		m1, pi1 = Split(&Delt3, &Phi3, &pi, q_start3)
		if m == m1 { 
			break 
		}
		m = m1 
		pi = pi1
	}

	h3 = make(map[int]int, len(pi))

	for i, x := range pi {
		if _, flag := h3[x]; flag { 
			pi[i] = h3[x]
		} else { 
			var (
				length int = len(h3)
			)
			pi[i] = length
			h3[x] = length 
		}
	}

	Delt1 = make([][]int, m)
	Phi1 = make([][]string, m)
	q01 = 0

	for i := 0; i < m; increment(&i) { 
		Delt1[i] = make([]int, len(Delt3[q_start3]))
		Phi1[i] = make([]string, len(Delt3[q_start3]))
	}
	
	for i := 0; i < len(Delt3); increment(&i) {
		q := pi[i]
		if i == q_start3 { 
			q01 = q 
		}
		for j := 0; j < len(Delt3[q_start3]); increment(&j) { 
			Delt1[q][j] = pi[Delt3[i][j]]
			Phi1[q][j] = Phi3[i][j] 
		}
	}

	DFS(&Delt1, &Phi1, q01)


	tmp_Delt1 = make([][]int, time)
	tmp_Phi1 = make([][]string, time)
	tmp_yes = false

	for i := 0; i < time; increment(&i) { 
		tmp_Delt1[i] = make([]int, len(Delt1[0]))
		tmp_Phi1[i] = make([]string, len(Delt1[0])) 
	}
	for i := 0; i < time; increment(&i) {
		tmp_yes = true
		var (
			x int = global_arr2[i]
			curr1 int = -1
			curr2 string = "x"
		)
		for j := 0; j < len(Delt1[x]); increment(&j) {
			curr1 = Delt1[x][j]
			tmp_Delt1[i][j] = global_arr1[curr1]
		}
		for j := 0; j < len(Phi1[x]); increment(&j) {
			curr2 = Phi1[x][j]
			tmp_Phi1[i][j] = curr2
		}
	}
	new_Delt = tmp_Delt1
	new_Phi = tmp_Phi1
	
	if tmp_yes { 
		minDelt3 = new_Delt
		minPhi3 = new_Phi
	} else {
		minDelt3 = nil
		minPhi3 = nil
	}

	compare(&mark_Delt, &mark_Phi, &minDelt2, &minPhi2, &minDelt3, &minPhi3)


	if keepGoing {
		if mark_Delt == len(minDelt2)*len(minDelt2[0]) && mark_Phi == len(minPhi2)*len(minPhi2[0]) { 
			fmt.Println("EQUAL") 
		} else { 
			fmt.Println("NOT EQUAL") 
		}
	}
}

func compare(mark_Delt *int, mark_Phi *int, Delt1 *[][]int, Phi1 *[][]string, Delt2 *[][]int, Phi2 *[][]string) {
	if len(*Delt1) != len(*Delt2) || len((*Delt1)[0]) != len((*Delt2)[0]) { 
		fmt.Println("NOT EQUAL")
		keepGoing = false
		return
	}
	for i := 0; i < len(*Delt1); increment(&i) { 
		for j := 0; j < len((*Delt1)[0]); increment(&j) { 
			if (*Delt1)[i][j] == (*Delt2)[i][j] { 
				increment(mark_Delt) 
			} 
		}
	}
	for i := 0; i < len(*Phi1); increment(&i) { 
		for j := 0; j < len((*Phi1)[0]); increment(&j) { 
			if (*Phi1)[i][j] == (*Phi2)[i][j] {
				increment(mark_Phi)
			} 
		} 
	}
}

func increment(i *int) {
	*i++
}

func Union(x int, y int, z int, pi *[]int) {
	for z < int(len(*pi)) { 
		if (*pi)[z] == x { 
			(*pi)[z] = y 
		}
		increment(&z)
	}
}

func Split(Delt *[][]int, Phi *[][]string, pi *[]int, q0 int) (int, []int) {
    var (
		m_classes int = int(len(*Delt))
    	pi1 []int = make([]int, 0, m_classes)
	)
	for i, _ := range *Delt { 
		pi1 = append(pi1, int(i)) 
	}
	for i := range *Delt {
		for j := range *Delt {
			if (*pi)[i] == (*pi)[j] && pi1[i] != pi1[j] {
				eq := true
				for k := 0; k < len((*Delt)[q0]); increment(&k) {
					w1, w2 := (*Delt)[i][k], (*Delt)[j][k]
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

func do_DFS(Delt *[][]int, Phi *[][]string, q int) {
	DFS(Delt, Phi, q)
}

func DFS(Delt *[][]int, Phi *[][]string, q int) {
	if 0 > global_arr1[q] { 
			global_arr1[q] = time
			increment(&time)
			global_arr2 = append(global_arr2, q)
			for i := 0; i < len((*Delt)[q]); increment(&i) {
				do_DFS(Delt, Phi, (*Delt)[q][i]) 
			} 
        }
}