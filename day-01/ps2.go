package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin) // create a scanner for each line
	re, err := regexp.Compile(`(\d+) +(\d+)`)

	list1 := make([]int, 0, 10)
	list2 := make([]int, 0, 10)

	if err != nil {
		fmt.Errorf("Regex exp is invalid", err)
	}

	for scanner.Scan() { // scan line by line
		line := scanner.Text() // get the line
		submatches := re.FindStringSubmatch(line)
		if len(submatches) != 0 {
			// convert to ints
			aToI1, _ := strconv.Atoi(submatches[1])
			aToI2, _ := strconv.Atoi(submatches[2])

			list1 = append(list1, aToI1)
			list2 = append(list2, aToI2)
		}
	}

	counter := make(map[int]int)

	for _, number := range list2 {
		if _, ok := counter[number]; ok {
			counter[number]++
		} else {
			counter[number] = 1
		}
	}

	acc := 0
	for _, number := range list1 {
		if multiplier, ok := counter[number]; ok {
			acc += number * multiplier
		}
	}

	fmt.Println(acc)
}
