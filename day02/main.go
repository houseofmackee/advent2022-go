package main

import (
	"bufio"
	"fmt"
	"os"
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

	totalScoreA := 0
	totalScoreB := 0
	for _, line := range lines {
		switch line {
		case "":
			continue

		case "A X":
			totalScoreA += (1 + 3)
			totalScoreB += (3 + 0)
		case "A Y":
			totalScoreA += (2 + 6)
			totalScoreB += (1 + 3)
		case "A Z":
			totalScoreA += (3 + 0)
			totalScoreB += (2 + 6)

		case "B X":
			totalScoreA += (1 + 0)
			totalScoreB += (1 + 0)
		case "B Y":
			totalScoreA += (2 + 3)
			totalScoreB += (2 + 3)
		case "B Z":
			totalScoreA += (3 + 6)
			totalScoreB += (3 + 6)

		case "C X":
			totalScoreA += (1 + 6)
			totalScoreB += (2 + 0)
		case "C Y":
			totalScoreA += (2 + 0)
			totalScoreB += (3 + 3)
		case "C Z":
			totalScoreA += (3 + 3)
			totalScoreB += (1 + 6)
		}
	}

	fmt.Printf("Total score round 1: %d\n", totalScoreA)
	fmt.Printf("Total score round 2: %d\n", totalScoreB)
}
