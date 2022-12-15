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

type Problem struct {
	carte [1000][1000]string 
	Sr,Sc int
	lowestRockRow int
}

type Position struct {
	column,row int
}

func solveProblem(fileName string) int {
	//set up map 
	p := setupProblem(fileName)
	p.print(fileName)
	
	unitsOfSand := 0
	for {
		grainOfSand := &Position{column:500, row: 0}
		grainOfSand.fall(p)

		//check where that unit of sand ended
		if grainOfSand.row == p.lowestRockRow {
			break
		} 
		unitsOfSand++
	}  
	
	return unitsOfSand
}

func (sand *Position) fall(p *Problem){

	currentRow := sand.row
	currentColumn := sand.column
	previousRow := sand.row
	for {
		if currentRow == p.lowestRockRow {
			sand.row = currentRow
			break
		}
		if previousRow > 0 && previousRow == currentRow { // we are resting
			p.carte[currentRow][currentColumn] = "o"
			sand.column = currentColumn
			sand.row = currentRow
			break
		}

		if currentRow > p.lowestRockRow || currentRow+1 >= 1000{ //we are falling in the abyss
			break
		}

		if p.carte[currentRow+1][currentColumn] == "." { //down
			previousRow = currentRow
			currentRow++  
			continue
		}

		if currentColumn-1 >= 0 && p.carte[currentRow+1][currentColumn-1] == "." { //left
			previousRow = currentRow
			currentRow++ 
			currentColumn-- 
			continue
		}

		if currentRow+1 < 1000 && p.carte[currentRow+1][currentColumn+1] == "." { //right
			previousRow = currentRow
			currentRow++ 
			currentColumn++ 
			continue
		}

		previousRow = currentRow
	} 
}

func (p *Problem) print(inputFile string) {
	f, err := os.Create("output_"+inputFile)
	check(err)

	for _, row := range p.carte {
		result := ""
		for _,item := range row {
			result += item
		}
		_, err := f.WriteString(result+"\n")
		check(err)
		
	}
	f.Sync()

	fmt.Println("Start is in position ",p.Sc," ", p.Sr)
	fmt.Println("Lowest rock is on row ",p.lowestRockRow)
}



func setupProblem(fileName string) *Problem{
	result := &Problem{Sc:500,Sr:0,lowestRockRow:-1}
	lowestRockRow:=0
	//initialise map
	carte := [1000][1000]string{}
	for i,row := range carte {
		for j := range row {
			carte[i][j] = "."
		}
	}

	//parse file
	readFile, err := os.Open(fileName)
	check(err)
	defer readFile.Close()
  
    fileScanner := bufio.NewScanner(readFile)
    fileScanner.Split(bufio.ScanLines)


	//498,4 -> 498,6 -> 496,6
	for fileScanner.Scan() {
		line := fileScanner.Text()
		fmt.Println("line: ", line)
		pairs := strings.Split(line," -> ")
		fmt.Println("pairs: ", pairs)
		for i := 0 ; i < len(pairs) - 1 ; i++ {
			leftPos := createPosition(pairs[i])
			rightPos := createPosition(pairs[i+1])

			currentRow := leftPos.row
			currentColumn := leftPos.column
			carte[rightPos.row][rightPos.column] = "#"
			for ; currentRow != rightPos.row || currentColumn != rightPos.column; {
				carte[currentRow][currentColumn] = "#" // NOTE column indices comes second

				if(currentRow > lowestRockRow){
					lowestRockRow = currentRow
				}

				diffRow := rightPos.row - currentRow
				diffCol := rightPos.column - currentColumn

				if(diffRow > 0){
					currentRow++
					continue
				}
				if(diffRow < 0){
					currentRow--
					continue
				}
				if(diffCol > 0){
					currentColumn++
					continue
				}
				if(diffCol < 0){
					currentColumn--
					continue
				}
			}
			
		}
	}
	
	result.carte = carte
	result.lowestRockRow = lowestRockRow
	return result
}

func createPosition(coordinateAsString string) *Position {
	coords := strings.Split(coordinateAsString,",")
	column,_ := strconv.Atoi(coords[0])
	row,_ := strconv.Atoi(coords[1])
	return &Position{column:column,row:row}
}







