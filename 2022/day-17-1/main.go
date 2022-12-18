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

type Problem struct {
	sequenced []string
	b1,b2,b3,b4,b5 *B
	cave [][]string
}

type B struct { // block
	cs []*C
}

type C struct { //coordinates
	y,x int
}

func printCave(cave [][]string){
	for i := len(cave) -1 ; i >= 0 ; i -- {
		line := ""
		for _,e := range cave[i] {
			line += e
		}
		fmt.Println(line, " ",i )
	}
}

func setupProblem(fileName string) *Problem{
	//parse file
	readFile, err := os.Open(fileName)
	check(err)
	defer readFile.Close()

    fileScanner := bufio.NewScanner(readFile)
    fileScanner.Split(bufio.ScanLines)

	sequenced := []string{}
	var sequence string
	for fileScanner.Scan() {
		sequence = fileScanner.Text()
		for _,rune := range sequence {
			sequenced = append(sequenced,string(rune))
		}
		break
	}

	// 0,0 is bottom left of the block
	b1 := &B{[]*C{{0,0},{0,1},{0,2},{0,3}}} // ####
	// .#.
	// ###
	// .#.
	b2 := &B{[]*C{{0,1},{1,0},{1,1},{1,2},{2,1}}}
	// ..#
	// ..#
	// ###
	b3 := &B{[]*C{{0,0},{0,1},{0,2},{1,2},{2,2}}}
	// #
	// #
	// #
	// #
	b4 := &B{[]*C{{0,0},{1,0},{2,0},{3,0}}}
	// ##
	// ##
	b5 := &B{[]*C{{0,0},{0,1},{1,0},{1,1}}}

	cave:=addRows(make([][]string,0),20)
	return &Problem{sequenced,b1,b2,b3,b4,b5,cave}
}

func addRows(cave [][]string, additionalRows int) [][]string{
	if additionalRows == 0 {
		return cave
	}
	
	newRow := [7]string{".",".",".",".",".",".","."}
	result := append(cave,newRow[0:])
	for i := 1 ; i < additionalRows; i++ {
		newRow := [7]string{".",".",".",".",".",".","."}
		result = append(result,newRow[0:])
	}
	return result
}

func abs(s int) int {
	if s > 0 {
		return s
	}
	return -s
}

func max(a,b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a,b int) int {
	if a > b {
		return b
	}
	return a
}

func solveProblem(fileName string) int {

	p := setupProblem(fileName)
	cave := p.cave
	shapes:=[5]*B{p.b1,p.b2,p.b3,p.b4,p.b5}
	sequence := p.sequenced
	
	rested_rocks := 0
	shape_type_number := 1
	sequence_index :=0
	for ; rested_rocks < 2022 ; {
		//find the highest rock
		highestRock := -1 
		for h := len(cave) - 1 ; h >= 0 ; h--{
			found := false
			for x:= 0 ; x < 7 ; x++ {
				if cave[h][x] == "#"{
					highestRock = h
					found = true
					break
				}
			}
			if found {
				break
			}
		}

		if len(cave) <= highestRock + 9 {
			cave = addRows(cave,10)
		}

		//place shape
		shapeType := shapes[(shape_type_number - 1) % 5]	
		shape_type_number++	
		shapeCs := make([]*C,0)
		for _,coord := range shapeType.cs {
			shapeCs = append(shapeCs,&C{coord.y+highestRock+4,coord.x+2})
		}
		currentBlock := &B{shapeCs}

		rests := false
		
 		for ; !rests ; {
		
			//left/right
			moveX := -1 
			if sequence[sequence_index % len(sequence)] == ">" {
				moveX = 1
			} 
			sequence_index++
			
			movedLeftRightShapeCS := make([]*C,0)
			for _,coord := range currentBlock.cs {
				movedLeftRightShapeCS = append(movedLeftRightShapeCS,&C{coord.y,coord.x+moveX})
			}

			shapeIsValid := true
			for _,coord := range movedLeftRightShapeCS {
				if coord.x < 0 || coord.x >= 7 || cave[coord.y][coord.x] == "#" {
					shapeIsValid = false
					break
				}
			}

			if shapeIsValid {
				currentBlock = &B{movedLeftRightShapeCS}
			}

			//down 
			movedDownShapeCS := make([]*C,0)
			for _,coord := range currentBlock.cs {
				movedDownShapeCS = append(movedDownShapeCS,&C{coord.y-1,coord.x})
			}

		
			for _,coord := range movedDownShapeCS {
				if coord.y < 0 || cave[coord.y][coord.x] == "#" {
					rests = true
					break
				}
			}

			if rests {
				for _,coord := range currentBlock.cs {
					cave[coord.y][coord.x] = "#"
				}
				rested_rocks++
				//printCave(cave)
			} else {
				currentBlock = &B{movedDownShapeCS}
			}
		}

	}

	//find the highest rock
	highestRock := 0 
	for h := len(cave) - 1 ; h >= 0 ; h--{
		found := false
		for x:= 0 ; x < 7 ; x++ {
			if cave[h][x] == "#"{
				highestRock = h
				found = true
				break
			}
		}
		if found {
			break
		}
	}

	
	return highestRock + 1 

}













