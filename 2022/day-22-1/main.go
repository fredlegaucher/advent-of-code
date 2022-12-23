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

type Position struct { 
	value  string 
	row,column int
	up,down,left,right *Position 
} 

func main() {
	a,b,c := solveProblem("input.txt")
	fmt.Println(1000 * a + 4 * b + c)
}

func (p *Position) print(){
	fmt.Println("Row ", p.row, " - column " , p.column)
}

func solveProblem(fileName string) (int,int,int) {

	startingPoint,instructions := setupProblem(fileName)
	
	facing := 0 //starting facing right
	position := startingPoint 
	position.print()

	stepsAsString := ""
	for _,rune := range instructions {
		symbol := string(rune)

		if symbol == "L" || symbol == "R" {
			//first move by steps accumulated to date
			steps,_ := strconv.Atoi(stepsAsString)
			carryOn := true
			for i:= 0 ; i < steps && carryOn ; i++ {
				position,carryOn = move(position,facing)
				if !carryOn {
					break
				}
				position.print()
			}
			
			stepsAsString = ""
			//turn
			if symbol == "L" {
				facing = (facing - 1 + 4) % 4
			} else {
				facing = (facing + 1) % 4
			}


		} else {
			stepsAsString += symbol
		}

	}

	if len(stepsAsString) > 0 {
		steps,_ := strconv.Atoi(stepsAsString)
		carryOn := true
		for i:= 0 ; i < steps && carryOn ; i++ {
			position,carryOn = move(position,facing)
			if !carryOn {
				break
			}
			position.print()
		}
	}

	return position.row+1,position.column+1,facing % 4
}

func move(position *Position, facing int) (*Position,bool){
	if facing % 4 == 0 && position.right != nil && position.right.value == "."{
		return position.right,true
	}
	if facing % 4 == 1  && position.down != nil && position.down.value == "."{ 
		return position.down,true
	}
	if facing % 4 == 2  && position.left != nil && position.left.value == "." {
		return position.left,true
	}
	if facing % 4 == 3  && position.up != nil && position.up.value == "." {
		return position.up,true
	}
	return position,false
}


func setupProblem(fileName string) (*Position,string){
	readFile, err := os.Open(fileName)
	check(err)
	defer readFile.Close()

    fileScanner := bufio.NewScanner(readFile)
    fileScanner.Split(bufio.ScanLines)

	var instructions string
	instructionsIsNext := false

	lines := []string{}
	maxWidth :=0
	
	for fileScanner.Scan() {
		line := fileScanner.Text()
		
		if len(line) == 0 {
			instructionsIsNext = true
			continue
		}

		if instructionsIsNext {
			instructions = line
		} else {
			lines = append(lines, line)
			if len(line) > maxWidth {
				maxWidth = len(line)
			}
		}

			
	}

	fmt.Print("width", maxWidth)

	topMostPositions := make([]int,maxWidth)
	for i := range topMostPositions {
		topMostPositions[i] = len(lines)+1
	}
	
	bottomMostPositions := make([]int,maxWidth)
	var current,previous,startingPoint *Position 
	overallMap := make(map[int]map[int]*Position)
	
	
	for row,line := range lines {
		overallMap[row] = make(map[int]*Position)
		current = nil
		previous = nil
		var firstForTheRow *Position
		
		for column,rune := range line {
			symbol := string(rune)
			
			if symbol == "." || symbol == "#" {
				current = &Position{value:symbol, row:row,column:column}
				overallMap[row][column]=current

				//maintain top values
				if row < topMostPositions[column] {
					topMostPositions[column]=row	
				}
				//maintain bottom values
				if (row > bottomMostPositions[column]){
					bottomMostPositions[column]=row
				}
				
				//chain on the left
				if previous != nil {
					current.left = previous
					previous.right = current
				}

				//chain above - apart from for top row
				up,ok := overallMap[row-1][column]
				if ok {
					current.up = up
					up.down = current
				}
			
			
				if row == 0 && startingPoint == nil {
					startingPoint = current
				}

				if firstForTheRow == nil {
					firstForTheRow = current
				}

				previous = current

				//chain to the left for the first item on the row - 2 scenarios
				//1. we are at the end of the row and on the map
				if column == len(line) - 1 {
					firstForTheRow.left = current
					current.right = firstForTheRow
					break
				}


			}

			//chain to the left for the first item on the row - 2 scenarios
			//2.  we are at the end of the row and there is blank space next
			if symbol == " " && previous != nil {
				firstForTheRow.left = previous
				previous.right = firstForTheRow
				previous = nil
				break
			}
		}
	}

	//go column by column to wrap-around vertically
	for column := 0 ; column < len(bottomMostPositions) ; column++ {
		fmt.Println(column)
		bottomPos := overallMap[bottomMostPositions[column]][column]
		topPos    := overallMap[topMostPositions[column]][column]

		bottomPos.down = topPos
		topPos.up = bottomPos
	}

	return startingPoint,instructions 	
}
