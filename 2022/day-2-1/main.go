package main

import (
	"bufio"
	"fmt"
	"os"
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
  



	var score int = 0

	var scoreMap = map[string]int{"A X": 4, "A Y": 8, "A Z": 3,
	"B X": 1, "B Y": 5, "B Z": 9,
	"C X": 7, "C Y": 2, "C Z": 6}	

    for fileScanner.Scan() {
        line := fileScanner.Text()
		score += scoreMap[line]
    }
	
	fmt.Println(score)

    readFile.Close()
}

