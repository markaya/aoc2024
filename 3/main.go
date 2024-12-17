package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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

	r, err := regexp.Compile(`mul\([0-9]{1,3},[0-9]{1,3}\)|do\(\)|don't\(\)`)
	if err != nil {
		panic(err)
	}

	flag := true
	counter := 0
	for fileScanner.Scan() {

		res := r.FindAll(fileScanner.Bytes(), -1)

		for _, v := range res {
			stringValue := string(v)
			if stringValue == "do()" {
				flag = true
			} else if stringValue == "don't()" {
				flag = false
			}
			if flag && stringValue != "do()" && stringValue != "don't()" {
				fmt.Println(stringValue)
				withouthPrefix, isFound := strings.CutPrefix(stringValue, "mul(")
				if !isFound {
					panic("No prefix mul(")
				}
				withoutSufix, isFound := strings.CutSuffix(withouthPrefix, ")")
				if !isFound {
					panic("No sufix )")
				}

				arr := strings.Split(withoutSufix, ",")
				if len(arr) != 2 {
					panic("long array")
				}

				int1, err := strconv.Atoi(arr[0])
				if err != nil {
					panic("Not good number 1")
				}
				int2, err := strconv.Atoi(arr[1])
				if err != nil {
					panic("Not good number 2")
				}

				res := int1 * int2

				counter += res
			}
		}
	}

	fmt.Println(counter)

	fmt.Println("====ENDS HERE====")
}
