package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
)

func main() {
	dat, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}
	var leftValues []int64
	var rightValues []int64

	leftList := true
	var num []byte
	for _, v := range dat {
		if string(v) == " " {
			if num == nil {
				continue
			}
			number, _ := strconv.ParseInt(string(num), 0, 64)
			leftValues = append(leftValues, number)
			leftList = !leftList
			num = nil
			continue
		} else if string(v) == ("\n") {
			number, _ := strconv.ParseInt(string(num), 0, 64)
			rightValues = append(rightValues, number)
			leftList = !leftList
			num = nil
			continue
		} else {
			num = append(num, v)
		}

	}

	if len(leftValues) != len(rightValues) {
		panic("Something not right")
	}
	slices.Sort(leftValues)
	slices.Sort(rightValues)

	var distancesSum int64
	for i := 0; i < len(leftValues); i++ {
		var distance int64
		if leftValues[i] > rightValues[i] {
			distance = leftValues[i] - rightValues[i]
		} else {
			distance = rightValues[i] - leftValues[i]
		}
		distancesSum = distancesSum + distance
	}

	var similarityScore int64

	for i := 0; i < len(leftValues); i++ {
		num := leftValues[i]
		counter := 0
		for j := 0; j < len(rightValues); j++ {
			if num == rightValues[j] {
				counter++
			}
		}
		similarityScore += num * int64(counter)
	}

	fmt.Println(similarityScore)
	fmt.Println(distancesSum)

}
