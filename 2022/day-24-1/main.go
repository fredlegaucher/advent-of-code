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
        // right, down, left, up
        nRow    = [4]int{0,1,0,-1}
        nColumn = [4]int{1,0,-1,0}
)

type Coordinate struct {
        row,column int
}

func solveProblem(fileName string) int {
	valley := setupProblem(fileName)
        maxRow := len(valley) - 1
        maxColumn := len(valley[0]) - 1
        step := 1
        currentPositions := []*Coordinate{{row:0,column:1}}

        for { // loop
                nextPositions := []*Coordinate{}
                for _,currentPos := range currentPositions { 
                        //let's see which of the 4 positions around me I can add
                        nextPositionsForOne :=  []*Coordinate{}
                        for d := 0 ; d < 4 ; d ++ {
                                nextPos := &Coordinate{row:currentPos.row+nRow[d],column:currentPos.column+nColumn[d]}
                                if nextPos.row <= maxRow &&
                                nextPos.row > 0 &&
                                nextPos.column <= maxColumn &&
                                nextPos.column >= 0 && 
                                string(valley[nextPos.row][nextPos.column]) != "#"{
                                //let's find out if there was wind coming from one of the 4 directions if we travel back into time
                                        rowNow := nextPos.row
                                        columnNow := nextPos.column
                                        //east
                                        rowThen := rowNow
                                        columnThen := ((columnNow - step - 1) + 1000 *  (maxColumn - 1)) % (maxColumn - 1) + 1
                                        if columnThen > 0 && string(valley[rowThen][columnThen]) == ">" {
                                                continue
                                        }
                                        //west
                                        rowThen = rowNow
                                        columnThen = ((columnNow + step - 1) + 1000 *  (maxColumn - 1)) % (maxColumn - 1) + 1
                                        if columnThen > 0 && string(valley[rowThen][columnThen]) == "<" {
                                                continue
                                        }
                                        //south
                                        rowThen = ((rowNow - step - 1) + 1000 *  (maxRow - 1)) % (maxRow - 1) + 1
                                        columnThen = columnNow
                                        if rowThen > 0 && string(valley[rowThen][columnThen]) == "v" {
                                                continue
                                        }
                                        //north
                                        rowThen = ((rowNow + step - 1) + 1000 *  (maxRow - 1)) % (maxRow - 1) + 1
                                        columnThen = columnNow
                                        if rowThen > 0 && string(valley[rowThen][columnThen]) == "^" {
                                                continue
                                        }


                                        //we found a place to go to
                                        nextPositionsForOne = append(nextPositionsForOne, nextPos)
                                }
                        }

                        if len(nextPositionsForOne) == 0 { // no position found so we wait from here
                                nextPositionsForOne = append(nextPositionsForOne, currentPos)
                        }

                        for _, newPos := range nextPositionsForOne {
                                duplicate := false
                                for _,existingPos := range nextPositions {
                                        if existingPos.column == newPos.column && existingPos.row == newPos.row {
                                                duplicate = true
                                        }
                                }
                                if ! duplicate {
                                        nextPositions = append(nextPositions, newPos)
                                }
                        }
                }

                currentPositions = nextPositions

                arrived:=false
                for _,pos := range currentPositions {
                        if pos.column == (maxColumn -1) && pos.row == maxRow {
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