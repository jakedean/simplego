package rotator

import (
	"fmt"
	"math/rand"
)

func Begin() {
	fmt.Printf("Welcome to the GO slice rotator, ")
	fmt.Printf("how big would you like to make the slice matrix?\n")
	var matrixSize int
	_, err:= fmt.Scanf("%v", &matrixSize)
	if err != nil {
		fmt.Println("There was an error fetching you matrix.")
	} else {
		fmt.Printf("Ok we will play with a %v by %v matrix.\n", matrixSize, matrixSize)
	}

	twoD := GenerateMatrix(matrixSize)
	fmt.Println("This is the 2d array we are going to rotate:")
	PrintSlice(twoD)

	RotateSlice(twoD)

}

func PrintSlice(sliceToPrint [][]int) {
	for _, val := range sliceToPrint {
		fmt.Printf("%v\n", val)
	}
}

func GenerateMatrix(matrixSize int) [][]int {
	matrix := make([][]int, matrixSize)
	for index, _ := range matrix {
		for i := 0; i < matrixSize; i++ {
			matrix[index] = append(matrix[index], rand.Intn(9))
		}
	}
	return matrix
}

func RotateSlice(twoD [][]int) {
	transformedTwoD := make([][]int, len(twoD))
	for _, val := range twoD {
		for i := 0; i < len(twoD); i++ {
			transformedTwoD[len(transformedTwoD)-1-i] = append(transformedTwoD[len(transformedTwoD)-1-i], val[i])
		}
	}

	fmt.Println("This original array rotated 90 degrees counter clockwise:")
	PrintSlice(transformedTwoD)
}