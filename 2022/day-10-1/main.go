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
	
	registerValuesDuringCyle := make(map[int]int)
	cycleCounter := 1
	registerValue := 1	
	combinedSignalStrength := 0;
	for fileScanner.Scan() {
        line := strings.Split(fileScanner.Text()," ")
		ops := line[0]
		
		if ops == "noop" {
			recordRegisterValueDuringCycle(registerValuesDuringCyle, registerValue, cycleCounter, &combinedSignalStrength)
			cycleCounter++
		}

		if ops == "addx" {
			recordRegisterValueDuringCycle(registerValuesDuringCyle, registerValue, cycleCounter, &combinedSignalStrength)
			cycleCounter++

			recordRegisterValueDuringCycle(registerValuesDuringCyle, registerValue, cycleCounter, &combinedSignalStrength)
			cycleCounter++

			value,_ := strconv.Atoi(line[1])
			registerValue += value
		}
			
	}

	return combinedSignalStrength;
}

func recordRegisterValueDuringCycle(registerHistory map[int]int, registerValue int, duringCycle int, combinedSignalStrength *int){
	registerHistory[duringCycle] = registerValue

	if duringCycle == 20 || duringCycle == 60 || duringCycle == 100 || duringCycle == 140 || duringCycle == 180 || duringCycle == 220 {
		*combinedSignalStrength += duringCycle * registerValue
	}
}
