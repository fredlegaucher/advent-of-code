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
	readFile, err := os.Open("input.txt")
  
    if err != nil {
        fmt.Println(err)
    }
    fileScanner := bufio.NewScanner(readFile)
 
    fileScanner.Split(bufio.ScanLines)

	lineMap := make(map[int]string)
	lineCount := 1
    for fileScanner.Scan() {
        lineMap[lineCount] = fileScanner.Text()
		lineCount++
    }

	fmt.Println(solveProblem(lineMap))

    readFile.Close()
}

func solveProblem(lines map[int]string) int {
	score := 0

	for _,line := range lines{
		
		score += scoreRucksack(line)
	}

	return score
}

func getCompartments(rucksack string) (firstCompartment string, secondCompartment string){
	if len(rucksack) % 2 == 1 {
		panic("rucksack not of even length")
	}

	return rucksack[:len(rucksack)/2],rucksack[len(rucksack)/2:] 
}

func findCommonItemtype(firstCompartment string, secondCompartment string) string {
	for _,r := range firstCompartment { 
		if strings.ContainsRune(secondCompartment,r) {
			return string(r)
		}
	}
	panic("no common letter found")
}

func scoreCommonLetter(commonLetter string) int {
	return strings.Index("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ",commonLetter) + 1
}

func scoreRucksack(rucksack string) int {
	return scoreCommonLetter(findCommonItemtype(getCompartments(rucksack)))
}
