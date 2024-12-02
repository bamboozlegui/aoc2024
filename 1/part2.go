package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func absint(integer int) int {
	if integer < 0 {
		return -integer
	}
	return integer
}

func main() {
	file, err := os.ReadFile("input.txt")
	check(err)

	couples := strings.Split(string(file), "\n")
	file_len := len(couples)
	couples = couples[:file_len-1]

	var values1 = make([]int, 0, file_len)
	var values2 = make([]int, 0, file_len)

	// parse input to two slices values1, values2
	for _, i := range couples {
		id := strings.Split(i, "   ")
		sliced, err := strconv.Atoi(id[0])
		values1 = append(values1, sliced)
		sliced, err = strconv.Atoi(id[1])
		values2 = append(values2, sliced)
		check(err)
	}
	// sort two lists
	slices.Sort(values1)
	slices.Sort(values2)
	// if smth wrong, abort
	if len(values1) != len(values2) {
		return
	}

	// sum up the difference
	answer := 0
	for i := 0; i < len(values1); i++ {
		answer = answer + absint(values1[i]-values2[i])
	}
	fmt.Println("answer", answer)

	sim_score := 0
	sim_score_total := 0
	for _, i := range values1 {
		for _, j := range values2 {
			if i == j {
				sim_score = sim_score + 1
			} else if i < j {
				break
			}
		}
		sim_score_total = sim_score_total + (i * sim_score)
		sim_score = 0
	}
	fmt.Println("similarity score: ", sim_score_total)
}

// define algorithm
// INPUT PARSE
// 12345 54321
// ...   ...
// 67890 09876
// TO LIST
// list 1: [12345,...,67890]
// list 2: [54321,..,09876]
// ALGORITHM
// 1. sort lists by ascending order.
// 2. subtract each members of the list.
// 3. sum the resulting list into some big int.
