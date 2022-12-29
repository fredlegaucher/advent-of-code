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
	fmt.Println(solveProblem("input.txt"))
}

var (
        // right, down, left, up, stay put
        nRow    = [5]int{0,1,0,-1,0}
        nColumn = [5]int{1,0,-1,0,0}
)

type Coordinate struct {
        row,column int
}

func (c *Coordinate) String() string {
        return fmt.Sprintf("row:%v-column%v",c.row,c.column)
}

func solveProblem(fileName string) int {
	valley := setupProblem(fileName)        
        start := &Coordinate{row:0,column:1}
        end   := &Coordinate{row:len(valley) - 1,column: len(valley[0]) - 2} 
        
        step1 := stepsFromStartToTarget(valley,start,end,1)
        fmt.Println("step 1 ", step1)
        
        step2 := stepsFromStartToTarget(valley,end,start,step1+1)
        

        return  stepsFromStartToTarget(valley,start,end,step2+1)
}

func stepsFromStartToTarget(valley []string, start,target *Coordinate, stepFrom int) int {
        maxRow := len(valley) - 1
        maxColumn := len(valley[0]) - 1
        step := stepFrom
        currentPositions := []*Coordinate{start}

        for { // loop
                nextPositions := []*Coordinate{}
                for _,currentPos := range currentPositions { 
                        //let's see which of the 4 positions around me I can move to/whether I can stay put
                        nextPositionsForOne :=  []*Coordinate{}
                        for d := 0 ; d < 5 ; d ++ {
                                nextPos := &Coordinate{row:currentPos.row+nRow[d],column:currentPos.column+nColumn[d]}
                                if nextPos.row <= maxRow &&
                                nextPos.row >= 0 &&
                                nextPos.column <= maxColumn &&
                                nextPos.column >= 0 && 
                                string(valley[nextPos.row][nextPos.column]) != "#" && 
                                canMoveToThisPosition(nextPos.row,nextPos.column,step,maxColumn,maxRow,valley){
                                        //we found a place to go to
                                        nextPositionsForOne = append(nextPositionsForOne, nextPos)
                                }
                        }

                        for _, newPos := range nextPositionsForOne {
                                duplicate := false
                                for _,existingPos := range nextPositions {
                                        if existingPos.column == newPos.column && existingPos.row == newPos.row {
                                                duplicate = true
                                                break
                                        }
                                }
                                if !duplicate {
                                        nextPositions = append(nextPositions, newPos)
                                }
                        }
                }

                currentPositions = nextPositions

                arrived:=false
                for _,pos := range currentPositions {
                        if pos.column == target.column && pos.row == target.row {
                                arrived=true
                                break
                        }

                }

                if arrived {
                        break
                }

                step++  
        }

        return step
}

func canMoveToThisPosition(rowNow,columnNow,step,maxColumn,maxRow int, valley []string) bool{
        //let's find out if there was wind coming from one of the 4 directions if we travel back into time

        //east
        rowThen := rowNow
        columnThen := ((columnNow - step - 1) + 1000 *  (maxColumn - 1)) % (maxColumn - 1) +1
        if columnThen <= 0 {
                panic("Should not happen")
        }
        if columnThen > 0 && string(valley[rowThen][columnThen]) == ">" {
                return false
        }
        //west
        rowThen = rowNow
        columnThen = ((columnNow + step - 1) + 1000 *  (maxColumn - 1)) % (maxColumn - 1) +1
        if columnThen <= 0 {
                panic("Should not happen")
        }
        if columnThen > 0 && string(valley[rowThen][columnThen]) == "<" {
        return false
        }
        //south
        rowThen = ((rowNow - step - 1) + 1000 * (maxRow - 1)) % (maxRow - 1) +1
        columnThen = columnNow
        if rowThen <= 0 {
                panic("Should not happen")
        }
        if rowThen > 0 && string(valley[rowThen][columnThen]) == "v" {
        return false
        }
        //north
        rowThen = ((rowNow + step - 1) + 1000 *  (maxRow - 1)) % (maxRow - 1) +1
        columnThen = columnNow
        if rowThen <= 0 {
                panic("Should not happen")
        }
        if rowThen > 0 && string(valley[rowThen][columnThen]) == "^" {
        return false
        }

        //hurrah!
        return true
}

func setupProblem(fileName string) []string{
        readFile, err := os.Open(fileName)
        check(err)
        defer readFile.Close()

        fileScanner := bufio.NewScanner(readFile)
        fileScanner.Split(bufio.ScanLines)

        valley := []string{} 
        for fileScanner.Scan() {
                valley = append(valley, fileScanner.Text()) 
        }
        return valley
}