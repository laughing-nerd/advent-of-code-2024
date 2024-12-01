package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./input")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	var (
		list1 []int
		list2 []int
	)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data := strings.Split(scanner.Text(), "   ") // The lists are separated by 3 spaces
		list1 = append(list1, str2int(data[0]))
		list2 = append(list2, str2int(data[1]))
	}
	part1(list1, list2)
	part2(list1, list2)

}

func part1(list1, list2 []int) {
	sum := 0
	sort.Ints(list1)
	sort.Ints(list2)
	for i := 0; i < len(list1); i++ {
		sum += int(math.Abs(float64(list1[i] - list2[i])))
	}
	fmt.Println("Part 1: ", sum)
}

func part2(list1, list2 []int) {
	var (
		similaritySum int
		frequencyMap  map[int]int = make(map[int]int)
	)

	for _, num := range list2 {
		frequencyMap[num]++
	}

	for _, num := range list1 {
		similaritySum += num * frequencyMap[num]
	}
	fmt.Println("Part 2: ", similaritySum)
}

// Helper func ...
func str2int(s string) int {
	intVal, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return intVal
}
