package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"slices"
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

	slices.Sort(list1)
	slices.Sort(list2)

	// fmt.Println(list1)
	// fmt.Println(list2)
	acc := 0

	for index, i := range list1 {
		j := list2[index]
		// fmt.Println(i, j, i-j)

		acc += int(math.Abs(float64(i - j)))
	}
	fmt.Println(acc)
}
