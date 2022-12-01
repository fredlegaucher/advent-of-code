package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
	fmt.Println(solveProblem("input.txt"))
}

func solveProblem(fileName string) int {
	readFile, err := os.Open(fileName)
	defer readFile.Close()
  
    if err != nil {
        fmt.Println(err)
    }
    fileScanner := bufio.NewScanner(readFile)
 
    fileScanner.Split(bufio.ScanLines)

	score := 0
	lineCount := 1
	previousLine := ""
    
	for fileScanner.Scan() {
        line := fileScanner.Text()

		if (lineCount % 3) == 1 {
			previousLine = line
		}
		if (lineCount % 3) == 2 {
			previousLine = findCommonElements(previousLine,line)
		}
		if (lineCount % 3) == 0 {
			score += scoreCommonLetter(findCommonElements(previousLine,line))
			previousLine = ""

		}

		lineCount++
	}
    
	return score
}

func findCommonElements(lhs string, rhs string) string {
	result := ""
	for _,r := range lhs { 
		if strings.ContainsRune(rhs,r) && !strings.ContainsRune(result,r) {
			result += string(r)
		}
	}
	return result
}

func scoreCommonLetter(commonLetter string) int {
	return strings.Index("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ",commonLetter) + 1
}

