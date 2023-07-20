package main

import (
	"bufio"
	"fmt"
	"os"
)

func stringIntersection(a, b string) string {
	var result string
	for _, r := range a {
		for _, s := range b {
			if r == s {
				result += string(r)
			}
		}
	}
	return result
}

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

	rucksackCompartments := make([][]string, len(lines))
	for i, line := range lines {
		numItems := len(line) / 2
		rucksackCompartments[i] = make([]string, 2)
		rucksackCompartments[i][0] = line[:numItems]
		rucksackCompartments[i][1] = line[numItems:]
	}

	sumOfItemValues := 0
	for i := 0; i < len(rucksackCompartments); i++ {
		itemValue := int(stringIntersection(rucksackCompartments[i][0], rucksackCompartments[i][1])[0])
		if itemValue >= 97 {
			sumOfItemValues += itemValue - 96
		} else {
			sumOfItemValues += itemValue - 64 + 26
		}
	}
	fmt.Printf("Sum of item values in compartments: %d\n", sumOfItemValues)

	sumOfItemValues = 0
	for i := 0; i < len(lines); i += 3 {
		itemValue := int(stringIntersection(lines[i], stringIntersection(lines[i+1], lines[i+2]))[0])
		if itemValue >= 97 {
			sumOfItemValues += itemValue - 96
		} else {
			sumOfItemValues += itemValue - 64 + 26
		}
	}
	fmt.Printf("Sum of item values in groups: %d\n", sumOfItemValues)
}
