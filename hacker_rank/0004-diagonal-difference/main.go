package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	stdin := bufio.NewScanner(os.Stdin)
	stdin.Scan()
	text := stdin.Text()
	matrixSize, _ := strconv.Atoi(text)
	numbers := []int{}
	for i := 0; i < matrixSize; i++ {
		stdin.Scan()
		text := stdin.Text()
		numbers = append(numbers, strToInt(text)...)
	}

	primaryDiagonalIndex := matrixSize + 1
	var primaryDiagonalSum int
	secondaryDiagonalIndex := matrixSize - 1
	var secondaryDiagonalSum int
	var countSum int

	for i := 0; i < len(numbers); i++ {
		if i == 0 {
			primaryDiagonalSum += numbers[i]
		} else if i%primaryDiagonalIndex == 0 {
			primaryDiagonalSum += numbers[i]
		}
		if i != 0 && i%secondaryDiagonalIndex == 0 && countSum < matrixSize {
			secondaryDiagonalSum += numbers[i]
			countSum++
		}

	}

	fmt.Println(math.Abs(float64(primaryDiagonalSum - secondaryDiagonalSum)))
}

func strToInt(str string) []int {
	strSlice := strings.Split(str, " ")
	numbers := []int{}
	for i := 0; i < len(strSlice); i++ {
		intNumber, _ := strconv.Atoi(strSlice[i])
		numbers = append(numbers, intNumber)
	}
	return numbers
}
