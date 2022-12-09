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
	child *Knot
}

var visitedPositions map[int]map[int]int = make(map[int]map[int]int)


func solveProblem(fileName string) int {
	readFile, err := os.Open(fileName)
	defer readFile.Close()
  
    if err != nil {
        fmt.Println(err)
    }
    fileScanner := bufio.NewScanner(readFile)
    fileScanner.Split(bufio.ScanLines)
	
	visitedPositions[0] = make(map[int]int)
	visitedPositions[0][0] = 1
	
	count := 1
	
	tail9 := &Knot{0,0,nil}
	eight := &Knot{0,0,tail9}
	seven := &Knot{0,0,eight}
	six := &Knot{0,0,seven}
	five := &Knot{0,0,six}
	four := &Knot{0,0,five}
	three := &Knot{0,0,four}
	two := &Knot{0,0,three}
	one := &Knot{0,0,two}
	head := &Knot{0,0,one}
	

	for fileScanner.Scan() {
        line := strings.Split(fileScanner.Text()," ")
		direction := line[0]
		steps,_ := strconv.Atoi(line[1])

		for step := 1; step <= steps ; step++ {
			head.moveKnotByOne(direction)	
			moveDescendant(head,&count)	
		}	
	}

	return count;
	
}

func moveDescendant(head *Knot, count *int){
	if head.child == nil {
		return
	}

	next:= head.child

	xDiff := head.X - next.X
	yDiff := head.Y - next.Y	
	if math.Abs(float64(xDiff)) <=1 && math.Abs(float64(yDiff)) <= 1 {
		return
	}

	switch {
		case xDiff > 0 : next.X++
		case xDiff < 0 : next.X-- 
	}

	switch {
		case yDiff > 0 : next.Y++
		case yDiff < 0 : next.Y-- 
	}

	if next.child == nil {

		if _, ok := visitedPositions[next.X]; !ok {
			visitedPositions[next.X] = make(map[int]int)
		}

		if visitedPositions[next.X][next.Y] == 0 {
			*count++
		}

		visitedPositions[next.X][next.Y]++

	}

	moveDescendant(next,count)
}

func (knot *Knot ) moveKnotByOne(direction string){
	switch direction {
		case "U": knot.Y++
		case "D": knot.Y--
		case "L": knot.X--
		case "R": knot.X++
		}
}
