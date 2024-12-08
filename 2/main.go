package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	dat, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer dat.Close()
	fileScanner := bufio.NewScanner(dat)

	var counter = 0
	var counterDamper = 0
	var linesNumber = 0

	for fileScanner.Scan() {
		linesNumber++
		var levels []int64
		for _, v := range strings.Split(fileScanner.Text(), " ") {
			level, _ := strconv.ParseInt(v, 10, 64)
			levels = append(levels, level)
		}
		if isSafe(levels) {
			counter++
		}
	}

	_, err = dat.Seek(0, 0)
	if err != nil {
		panic(err)
	}
	fileScanner = bufio.NewScanner(dat)

	// Part 2
	for fileScanner.Scan() {
		var levels []int64
		linesNumber++
		for _, v := range strings.Split(fileScanner.Text(), " ") {
			level, _ := strconv.ParseInt(v, 10, 64)
			levels = append(levels, level)
		}
		if isSafeDamper(levels) {
			counterDamper++
		}
	}

	fmt.Println("From:", linesNumber)
	fmt.Println("Safe:", counter)
	fmt.Println("Damper Safe:", counterDamper)

}

func isSafe(levels []int64) bool {
	first := levels[0]
	second := levels[1]
	isAscending := true

	if first == second {
		return false
	} else if first > second {
		isAscending = false
	}
	for _, el := range levels[1:] {

		if isAscending {
			if el <= first || el-first > 3 {
				return false
			}
		} else {
			if first <= el || first-el > 3 {
				return false
			}
		}
		first = el
	}
	return true
}

func isSafeDamper(levels []int64) bool {
	first := levels[0]
	second := levels[1]
	isAscending := true

	if first == second {
		return tryDamper(levels, 0)
	} else if first > second {
		isAscending = false
	}
	for i, el := range levels[1:] {
		if isAscending {
			if el <= first || el-first > 3 {
				return tryDamper(levels, i)
			}
		} else {
			if first <= el || first-el > 3 {
				return tryDamper(levels, i)
			}
		}
		first = el
	}
	return true
}

func tryDamper(levels []int64, firstId int) bool {
	var slice []int64
	if firstId == 1 && isSafe(levels[1:]) {
		return true
	} else {
		slice = append(slice, levels[:firstId]...)
		slice = append(slice, levels[firstId+1:]...)

		if isSafe(slice) {
			return true
		} else {
			var ns []int64
			if firstId+2 >= len(levels) {
				ns = append(ns, levels[:firstId+1]...)
			} else {
				ns = append(ns, levels[:firstId+1]...)
				ns = append(ns, levels[firstId+2:]...)
			}
			return isSafe(ns)
		}
	}
}
