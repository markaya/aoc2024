package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	mapset "github.com/deckarep/golang-set/v2"
)

func main() {
	dat, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer dat.Close()
	fileScanner := bufio.NewScanner(dat)

	myMap := make(map[int][]int)

	var counter int
	var counterWithOrder int

	for fileScanner.Scan() {
		txt := fileScanner.Text()

		if strings.Contains(txt, "|") {
			out := strings.Split(txt, "|")
			if len(out) != 2 {
				panic("Invalid input")
			}

			k, err := strconv.Atoi(out[0])
			if err != nil {
				panic("Not int")
			}
			v, err := strconv.Atoi(out[1])
			if err != nil {
				panic("Not int")
			}

			addToMap(k, v, myMap)

		} else {
			parts := strings.Split(strings.TrimSpace(txt), ",")
			if len(parts) > 1 {

				numbers := make([]int, len(parts))
				for i, part := range parts {
					numbers[i], err = strconv.Atoi(strings.TrimSpace(part))
					if err != nil {
						panic("error parsing number:")
					}
				}
				checkParts(numbers, myMap)
				if checkParts(numbers, myMap) {
					middle := numbers[(len(numbers)-1)/2]
					counter += middle

				} else {
					ordered := orderFailedParts(numbers, myMap)
					middle := ordered[(len(ordered)-1)/2]
					counterWithOrder += middle

				}
			}

		}

	}

	fmt.Println(counter)
	fmt.Println(counterWithOrder)

}

func addToMap(key, value int, selectedMap map[int][]int) {
	list := selectedMap[key]
	selectedMap[key] = append(list, value)

}

func orderFailedParts(numbers []int, myMap map[int][]int) []int {
	for k, v := range numbers {
		mapSet := mapset.NewSet(myMap[v]...)
		restSet := mapset.NewSet(numbers[k+1:]...)
		if !restSet.IsSubset(mapSet) {
			s := append(numbers[:k], numbers[k+1:]...)
			s = append(s, v)
			return orderFailedParts(s, myMap)
		}
	}
	return numbers
}

func checkParts(numbers []int, myMap map[int][]int) bool {
	for k, v := range numbers {
		mapSet := mapset.NewSet(myMap[v]...)
		restSet := mapset.NewSet(numbers[k+1:]...)
		if !restSet.IsSubset(mapSet) {
			return false
		}
	}
	return true
}
