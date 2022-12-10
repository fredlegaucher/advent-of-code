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
	for _,text := range solveProblem("input.txt") {
		fmt.Println(text)
	}
}

var render [6]string 
var spriteCentralPosition int = 1	
var crtPosition int = 0
var currentCrtLine int = 0

func solveProblem(fileName string) [6]string {
	readFile, err := os.Open(fileName)
	defer readFile.Close()
  
    if err != nil {
        fmt.Println(err)
    }
    fileScanner := bufio.NewScanner(readFile)
    fileScanner.Split(bufio.ScanLines)
	
	
	registerHistory := make(map[int]int)
	cycleCounter := 1
	
	
	for fileScanner.Scan() {
        line := strings.Split(fileScanner.Text()," ")
		ops := line[0]
		
		if ops == "noop" {
			registerHistory[cycleCounter] = spriteCentralPosition
			renderThisCycle(spriteCentralPosition)
			cycleCounter++
		}

		if ops == "addx" {
			registerHistory[cycleCounter] = spriteCentralPosition
			renderThisCycle(spriteCentralPosition)
			cycleCounter++

			registerHistory[cycleCounter] = spriteCentralPosition
			renderThisCycle(spriteCentralPosition)
			cycleCounter++

			value,_ := strconv.Atoi(line[1])
			spriteCentralPosition += value
		}
			
	}

	return render;
}

func renderThisCycle(spriteCentralPosition int){
	if spriteCentralPosition-1 <= crtPosition && crtPosition <= spriteCentralPosition +1 {
		render[currentCrtLine] += "#"
	} else {
		render[currentCrtLine] += "."
	}

	crtPosition++

	if crtPosition == 40 {
		crtPosition = 0
		currentCrtLine++
	}
}

