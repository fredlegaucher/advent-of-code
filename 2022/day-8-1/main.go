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


func solveProblem(fileName string) int {
		
	readFile, err := os.Open(fileName)
	defer readFile.Close()
  
    if err != nil {
        fmt.Println(err)
    }
    fileScanner := bufio.NewScanner(readFile)
    fileScanner.Split(bufio.ScanLines)


	columns := make(map[int][]int)
	currentRowIndex := 0;
	visibleTreeCount := 0 
	visibleTrees := make(map[int]map[int]bool)
	
	//handle the rows
	for fileScanner.Scan() {
        line := fileScanner.Text()
 		
		biggestTreeSoFar := 0
		currentRow := make([]int,0)
		visibleTrees[currentRowIndex] = make(map[int]bool)	
		
		//looking right + gathering data
		for currentColumnIndex,c := range line {
			currentTreeHeight := int(c)

			//populate row
			currentRow = append(currentRow, currentTreeHeight)

			//populate columns
			if currentRowIndex == 0 {
				columns[currentColumnIndex] = make([]int,0)
			}
			
			columns[currentColumnIndex] = append(columns[currentColumnIndex], currentTreeHeight)

			if currentColumnIndex == 0 || currentTreeHeight > biggestTreeSoFar {
				visibleTreeCount++
				visibleTrees[currentRowIndex][currentColumnIndex] = true
				biggestTreeSoFar = currentTreeHeight
				continue
			}
		}

		//looking left
		biggestTreeSoFar = 0
		for currentColumnIndex := len(currentRow) - 1 ; currentColumnIndex >= 0 ; currentColumnIndex-- {
			currentTreeHeight:= currentRow[currentColumnIndex]

			if currentColumnIndex == len(currentRow) - 1 || currentTreeHeight > biggestTreeSoFar {
				if !visibleTrees[currentRowIndex][currentColumnIndex]{ // avoid double counting a given tree
					visibleTreeCount++
				}
				visibleTrees[currentRowIndex][currentColumnIndex] = true
				biggestTreeSoFar = currentTreeHeight
				continue
			}
		}
		currentRowIndex++
	}



	for currentColumnIndex,currentColumn := range columns {

		// now let's look down
		biggestTreeSoFar := 0
		for currentRowIndex := 0 ; currentRowIndex < len(currentColumn) ; currentRowIndex++ {
			currentTreeHeight:= currentColumn[currentRowIndex]

			if currentRowIndex == 0 || currentTreeHeight > biggestTreeSoFar {
				if !visibleTrees[currentRowIndex][currentColumnIndex]{ // avoid double counting a given tree
					visibleTreeCount++
				}
				visibleTrees[currentRowIndex][currentColumnIndex] = true
				biggestTreeSoFar = currentTreeHeight
				continue
			}
		}

		// and up
		biggestTreeSoFar = 0
		for currentRowIndex := len(currentColumn) - 1 ; currentRowIndex >= 0 ; currentRowIndex-- {
			currentTreeHeight:= currentColumn[currentRowIndex]

			if currentRowIndex == len(currentColumn) - 1 || currentTreeHeight > biggestTreeSoFar {
				if !visibleTrees[currentRowIndex][currentColumnIndex]{ // avoid double counting a given tree
					visibleTreeCount++
				}
				visibleTrees[currentRowIndex][currentColumnIndex] = true
				biggestTreeSoFar = currentTreeHeight
				continue
			}
		}
	}

	return visibleTreeCount
}

