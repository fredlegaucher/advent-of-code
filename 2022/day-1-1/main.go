package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
  
	var maxCalories int = 0
	var currentElfCalories int = 0

    for fileScanner.Scan() {
        line := fileScanner.Text()
		if len(line) == 0 {
			if currentElfCalories > maxCalories {
				maxCalories = currentElfCalories
			}
			currentElfCalories = 0
		} else {
			i, err := strconv.Atoi(line)
			check(err)
			currentElfCalories += i
		}
    }
	
	fmt.Println(maxCalories)

    readFile.Close()
}

