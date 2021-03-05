package main

import (
	"fmt"
	"math/big"
)

func main() {
	var nel int64
	fmt.Scan(&nel)

	if 1 == nel || 2 == nel {
		fmt.Println(1)
	} else {
		var matrix1 [2][2]*big.Int = setMatrix(1) 
		var matrix2 [2][2]*big.Int = setMatrix(2)
		fastPower(nel - 1, matrix1, matrix2)
	}
}

func setMatrix(index int) (resultMatrix [2][2]*big.Int) {
	if 1 == index {
		resultMatrix[0][0] = big.NewInt(1)
		resultMatrix[0][1] = big.NewInt(1)
		resultMatrix[1][0] = big.NewInt(1)
		resultMatrix[1][1] = big.NewInt(0)
	} else if 2 == index {
		resultMatrix[0][0] = big.NewInt(1)
		resultMatrix[0][1] = big.NewInt(0)
		resultMatrix[1][0] = big.NewInt(0)
		resultMatrix[1][1] = big.NewInt(1)
	} else {
		resultMatrix[0][0] = big.NewInt(0)
		resultMatrix[0][1] = big.NewInt(0)
		resultMatrix[1][0] = big.NewInt(0)
		resultMatrix[1][1] = big.NewInt(0)
	}
	return
}

func matrixMulty(M1 [2][2]*big.Int, M2 [2][2]*big.Int) (result [2][2]*big.Int) {
	result = setMatrix(0)
	tmp := big.NewInt(0)
	result[0][0].Add(result[0][0], tmp.Mul(M1[0][0], M2[0][0]))
	tmp = big.NewInt(0)
	result[0][0].Add(result[0][0], tmp.Mul(M1[0][1], M2[1][0]))
	tmp = big.NewInt(0)
	result[0][1].Add(result[0][1], tmp.Mul(M1[0][0], M2[0][1]))
	tmp = big.NewInt(0)
	result[0][1].Add(result[0][1], tmp.Mul(M1[0][1], M2[1][1]))
	tmp = big.NewInt(0)
	result[1][0].Add(result[1][0], tmp.Mul(M1[1][0], M2[0][0]))
	tmp = big.NewInt(0)
	result[1][0].Add(result[1][0], tmp.Mul(M1[1][1], M2[1][0]))
	tmp = big.NewInt(0)
	result[1][1].Add(result[1][1], tmp.Mul(M1[1][0], M2[0][1]))
	tmp = big.NewInt(0)
	result[1][1].Add(result[1][1], tmp.Mul(M1[1][1], M2[1][1]))
	return
}

func fastPower(n int64, M1 [2][2]*big.Int, M2 [2][2]*big.Int) {
	for n > 0 {
		if n&1 == 1 {
			M2 = matrixMulty(M2, M1)
		}
		M1 = matrixMulty(M1, M1)
		n >>= 1
	}

	result := big.NewInt(0)
	fmt.Println(result.Add(M2[1][0], M2[1][1]))
}
