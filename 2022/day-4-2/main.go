package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
    
	for fileScanner.Scan() {
        line := fileScanner.Text()

		splittedLine := strings.Split(line,",")
		firstElfLow,_ := strconv.Atoi(strings.Split(splittedLine[0],"-")[0])
		firstElfHigh,_ := strconv.Atoi(strings.Split(splittedLine[0],"-")[1])
		secondElfLow,_ := strconv.Atoi(strings.Split(splittedLine[1],"-")[0])
		secondElfHigh,_ := strconv.Atoi(strings.Split(splittedLine[1],"-")[1])

		if !(firstElfHigh < secondElfLow || secondElfHigh < firstElfLow) {	
			score++
		}
	}
    
	return score
}