package main

import (
	"bufio"
	"fmt"
	"math"
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

type Knot struct {
	X int
	Y int
}


func solveProblem(fileName string) int {
	readFile, err := os.Open(fileName)
	defer readFile.Close()
  
    if err != nil {
        fmt.Println(err)
    }
    fileScanner := bufio.NewScanner(readFile)
    fileScanner.Split(bufio.ScanLines)


	visitedPositions := make(map[int]map[int]int)
	visitedPositions[0] = make(map[int]int)
	visitedPositions[0][0] = 1
	count := 1
	head := &Knot{0,0}
	tail := &Knot{0,0}

	for fileScanner.Scan() {
        line := strings.Split(fileScanner.Text()," ")
		direction := line[0]
		steps,_ := strconv.Atoi(line[1])

		for step := 1; step <= steps ; step++ {
			head.moveKnotByOne(direction)

			xDiff := head.X - tail.X
			yDiff := head.Y - tail.Y	
			if math.Abs(float64(xDiff)) <=1 && math.Abs(float64(yDiff)) <= 1 {
				continue
			}

			switch {
				case xDiff > 0 : tail.X++
				case xDiff < 0 : tail.X-- 
			}

			switch {
				case yDiff > 0 : tail.Y++
				case yDiff < 0 : tail.Y-- 
			}

			if _, ok := visitedPositions[tail.X]; !ok {
				visitedPositions[tail.X] = make(map[int]int)
			}

			if visitedPositions[tail.X][tail.Y] == 0 {
				count++
			}

			visitedPositions[tail.X][tail.Y]++
		}	
	}

	return count;
	
}

func (knot *Knot ) moveKnotByOne(direction string){
	switch direction {
		case "U": knot.Y++
		case "D": knot.Y--
		case "L": knot.X--
		case "R": knot.X++
		}
}
