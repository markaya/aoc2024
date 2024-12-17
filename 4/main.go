package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

const X byte = byte('X')
const M byte = byte('M')
const A byte = byte('A')
const S byte = byte('S')

var iDim int
var jDim int

var counterChannel = make(chan int)
var counterChannel2 = make(chan int)

func main() {
	dat, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer dat.Close()
	fileScanner := bufio.NewScanner(dat)

	var matrix [][]byte
	for fileScanner.Scan() {
		bts := append([]byte{}, fileScanner.Bytes()...)
		matrix = append(matrix, bts)

	}
	iDim = len(matrix)
	jDim = len(matrix[0])
	var counter int
	var counter2 int

	go func() {
		for v := range counterChannel {
			counter += v
		}
	}()
	go func() {
		for v := range counterChannel2 {
			counter2 += v
		}
	}()

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			charVal := matrix[i][j]
			if charVal == X || charVal == S {
				counterChannel <- check(i, j, matrix)
			}
			if charVal == A {
				counterChannel2 <- check2(i, j, matrix)
			}
		}
	}

	fmt.Println(counter)
	fmt.Println(counter2)

}

var strings = []string{"MSSM", "SMMS", "SSMM", "MMSS"}

func check2(i, j int, matrix [][]byte) int {
	if i-1 < 0 || j-1 < 0 || i+1 >= iDim || j+1 >= jDim {
		return 0
	}
	var checkString = string([]byte{matrix[i-1][j-1], matrix[i-1][j+1], matrix[i+1][j+1], matrix[i+1][j-1]})

	if slices.Contains(strings, checkString) {
		return 1
	}
	return 0

}
func check(i, j int, matrix [][]byte) int {
	var dirMatrix = [][]int{}

	if j+3 < jDim {
		if i-3 >= 0 {
			dirMatrix = append(dirMatrix, []int{-1, 1})
		}

		dirMatrix = append(dirMatrix, []int{0, 1})

		if i+3 < iDim {
			dirMatrix = append(dirMatrix, []int{1, 1})
		}
	}
	if i+3 < iDim {
		dirMatrix = append(dirMatrix, []int{1, 0})
	}

	counter := 0

	for _, v := range dirMatrix {
		var arr []byte
		arr = append(arr, matrix[i][j])

		i, j := i+v[0], j+v[1]
		arr = append(arr, matrix[i][j])

		i, j = i+v[0], j+v[1]
		arr = append(arr, matrix[i][j])

		i, j = i+v[0], j+v[1]
		arr = append(arr, matrix[i][j])

		if string(arr) == "XMAS" || string(arr) == "SAMX" {
			counter++
		}
	}
	return counter
}
