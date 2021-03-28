package main

import ("fmt")

func main() {
	var (
		num int64
		res int64
	)
	fmt.Scan(&num)
	if num < 9 {
		fmt.Print(num + 1)
	} else {
		num++
		var (
			tmp int64 = (num - 9)
			i int64
		)
		for i = 1; tmp > 0; i++ {
			num = tmp;
			tmp = num - 9 * expon(i) * (i + 1)
		}
		if (num / i == 0) {
			res = expon(i - 1)
		} else {
			res = expon(i - 1) + num % i + num / i - 1
		}
		for num = i - num % i; num % i != 0 && num > 0; num-- {
			res /= 10
		}
		fmt.Println(res % 10)
	}

}

func expon(x int64) int64{
	var (
		res int64 = 1
		i int64 = 1
	)
	if x == 0 {
		return res
	}
	for i = 1; i < (x + 1); i++ {
		res *= 10
	}
	return res
}