package main

import (
	"bufio"
	"os"
)

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

}
