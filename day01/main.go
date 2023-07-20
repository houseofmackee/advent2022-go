package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ./main <file>")
		return
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Open file error: ", err)
		return
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	var elves = make(map[int]int)
	elfNumber := 1
	for _, line := range lines {
		if line == "" {
			elfNumber++
			continue
		}
		value, err := strconv.Atoi(strings.TrimSpace(line))
		if err != nil {
			fmt.Println("Convert error: ", err)
			return
		}
		elves[elfNumber] += value
	}

	if len(elves) == 0 {
		fmt.Println("No elves found")
		return
	}

	elfNumber = 1
	elfCalories := elves[elfNumber]
	var elfTopThree []int
	for k, v := range elves {
		elfTopThree = append(elfTopThree, v)
		if v > elfCalories {
			elfNumber = k
			elfCalories = v
		}
	}
	fmt.Printf("Elf %d has the most calories: %d\n", elfNumber, elfCalories)

	sort.Slice(elfTopThree, func(i, j int) bool {
		return elfTopThree[i] > elfTopThree[j]
	})

	topThreeCalories := elfTopThree[0] + elfTopThree[1] + elfTopThree[2]
	fmt.Printf("Top three calories: %d\n", topThreeCalories)
}
