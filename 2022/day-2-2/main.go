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
  

// X means you need to lose, Y means you need to end the round in a draw, and Z means you need to win. Good luck!"
   //               X Lose 0    Y Draw  3   Z Win 6               
//A Rock            Scissors 3  Rock 4      Papper 8
//B Paper           Rock 1      Paper 5     Scissors 9
//C Scissors        Paper 2  Scissors 6     Rock 7


	var score int = 0

	var scoreMap = map[string]int{"A X":3, "A Y": 4, "A Z": 8,
	"B X": 1, "B Y": 5, "B Z": 9,
	"C X": 2, "C Y": 6, "C Z": 7}	

    for fileScanner.Scan() {
        line := fileScanner.Text()
		score += scoreMap[line]
    }
	
	fmt.Println(score)

    readFile.Close()
}

