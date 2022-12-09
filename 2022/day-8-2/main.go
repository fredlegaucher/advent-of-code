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


	rows := make(map[int][]int)
	columns := make(map[int][]int)
	
	//gather the data
	currentRowIndex  := 0;
	for fileScanner.Scan() {
        line := fileScanner.Text()
 		
		currentRow := make([]int,0)

		//populate columns
			

		//looking right 
		for currentColumnIndex,c := range line {
			currentTreeHeight := int(c - '0')

			//populate row
			currentRow = append(currentRow, currentTreeHeight)

			if currentRowIndex == 0 {
				columns[currentColumnIndex] = make([]int,0)
			}
			columns[currentColumnIndex] = append(columns[currentColumnIndex], currentTreeHeight)
		}
		rows[currentRowIndex] = currentRow
		currentRowIndex++
	}

	maxScenicScore := 0 

	for currentRowIndex,currentRow := range rows {

		for currentColumnIndex,currentTreeHeight := range currentRow {
			if currentRowIndex == 0 || currentColumnIndex == 0 || currentRowIndex == len(currentRow) - 1 || currentColumnIndex == len(columns[currentColumnIndex]) -1 {
				continue
			}

			currentTreeScenicScore := calculateScore(currentTreeHeight,currentRowIndex,currentColumnIndex,currentRow,columns[currentColumnIndex])
			if currentTreeScenicScore > maxScenicScore {
				maxScenicScore = currentTreeScenicScore
			}
		}
	}
	return maxScenicScore
}

func calculateScore(currentTreeHeight, currentRowIndex, currentColumnIndex int, currentRow, currentColumn []int) int {
	//look left
	treesOnLeft :=0
	for i := currentColumnIndex - 1 ; i >= 0 ; i-- {
		
		if currentRow[i] < currentTreeHeight {
			treesOnLeft++
			continue
		}

		treesOnLeft++
		break
	}

	//look right
	treesOnRight :=0
	for i := currentColumnIndex + 1; i < len(currentRow) ; i++ {

		if currentRow[i] < currentTreeHeight {
			treesOnRight++
			continue
		}
		treesOnRight++
		break
	} 
	//look up
	treesUp := 0
	for i := currentRowIndex -1 ; i >= 0 ; i-- {
		if currentColumn[i] < currentTreeHeight {
			treesUp++
			continue
		}
		treesUp++
		break
	}

	//look down
	treesDown := 0
	for i := currentRowIndex + 1; i < len(currentColumn) ; i++ {
		if currentColumn[i] < currentTreeHeight {
			treesDown++
			continue
		}
		treesDown++
		break
	} 

	return treesOnLeft * treesOnRight * treesUp * treesDown
}

