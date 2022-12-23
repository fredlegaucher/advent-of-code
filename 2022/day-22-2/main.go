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

func main() {
	a,b,c := solveProblem("input.txt")
	fmt.Println(1000 * a + 4 * b + c)
}


func solveProblem(fileName string) (int,int,int) {

	sRow,sColumn,sFace,overallMap,instructions := setupProblem(fileName)
	fmt.Println("Starting point ", sRow,"-", sColumn, " on face ", sFace)

	facing := 0 //starting facing right
	cRow := sRow
	cColumn := sColumn
	cFace := sFace

	stepsAsString := ""
	for _,rune := range instructions {
		symbol := string(rune)

		if symbol == "L" || symbol == "R" {
			//first move by steps accumulated to date
			steps,_ := strconv.Atoi(stepsAsString)
			carryOn := true
			for i := 0 ; i < steps && carryOn ; i++ {
				cRow,cColumn,cFace,facing,carryOn = move(cRow,cColumn,cFace,facing,overallMap)
				if !carryOn {
					break
				}
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
			cRow,cColumn,cFace,facing,carryOn = move(cRow,cColumn,cFace,facing,overallMap)
			if !carryOn {
				break
			}
		}
	}

	fmt.Println("row ", cRow,"column", cColumn," face ", cFace, "facing" , facing)
	return cRow+1,cColumn+1,facing
}

func move(cRow,cColumn,cFace, facing int,overall map[int]map[int]string) (int,int,int,int,bool){
	if facing % 4 == 0 {
		symbol,ok := overall[cRow][cColumn+1]
		if ok {
			if symbol == "." {
				return cRow,cColumn+1,whichFace(cRow,cColumn+1),facing,true
			}
			return cRow,cColumn,cFace,facing,false
		} 
	}

	if facing % 4 == 1 {
		symbol,ok := overall[cRow+1][cColumn]
		if ok {
			if symbol == "." {
				return cRow+1,cColumn,whichFace(cRow+1,cColumn),facing,true
			}
			return cRow,cColumn,cFace,facing,false
		}
	}

	if facing % 4 == 2 {
		symbol,ok := overall[cRow][cColumn-1]
		if ok {
			if symbol == "." {
				return cRow,cColumn-1,whichFace(cRow,cColumn-1),facing,true
			}
			return cRow,cColumn,cFace,facing,false
		}
	}

	if facing % 4 == 3 {
		symbol,ok := overall[cRow-1][cColumn]
		if ok {
			if symbol == "." {
				return cRow-1,cColumn,whichFace(cRow-1,cColumn),facing,true
			}
			return cRow,cColumn,cFace,facing,false
		}
	}

	//we have not returned so changing face to find the next item
	outRow,outColumn,outFace,outFacing := teleport(cRow,cColumn,cFace,facing)
	symbolNextFace,okNextFace := overall[outRow][outColumn]
	if !okNextFace {
		panic("We have a bug")
	}
	if symbolNextFace == "#" { //we cannot move forward with this
		return cRow,cColumn,cFace,facing,false
	}
	return outRow,outColumn,outFace,outFacing,true
}

// Facing is 0 for right (>), 1 for down (v), 2 for left (<), and 3 for up (^)
func teleport(inRow,inColumn,inFace,inFacing int) (int,int,int,int){
	outFace, outFacing,outRow,outColumn := -1,-1,-1,-1
	switch {
	case inFace == 2 && inFacing == 3: outFace,outFacing,outRow,outColumn = 6,0,inColumn%50,0
	case inFace == 1 && inFacing == 3: outFace,outFacing,outRow,outColumn = 6,3,49,inColumn%50
	case inFace == 1 && inFacing == 0: outFace,outFacing,outRow,outColumn = 4,2,49-(inRow%50),49
	case inFace == 1 && inFacing == 1: outFace,outFacing,outRow,outColumn = 3,2,inColumn%50,49
	case inFace == 3 && inFacing == 0: outFace,outFacing,outRow,outColumn = 1,3,49,inRow%50
	case inFace == 4 && inFacing == 0: outFace,outFacing,outRow,outColumn = 1,2,49-(inRow%50),49
	case inFace == 4 && inFacing == 1: outFace,outFacing,outRow,outColumn = 6,2,inColumn%50,49	
	case inFace == 6 && inFacing == 0: outFace,outFacing,outRow,outColumn = 4,3,49,inRow%50
	case inFace == 6 && inFacing == 1: outFace,outFacing,outRow,outColumn = 1,1,0,inColumn%50
	case inFace == 6 && inFacing == 2: outFace,outFacing,outRow,outColumn = 2,1,0,inRow%50
	case inFace == 5 && inFacing == 2: outFace,outFacing,outRow,outColumn = 2,0,49-(inRow%50),0		
	case inFace == 5 && inFacing == 3: outFace,outFacing,outRow,outColumn = 3,0,inColumn%50,0
	case inFace == 3 && inFacing == 2: outFace,outFacing,outRow,outColumn = 5,1,0,inRow%50				
	case inFace == 2 && inFacing == 2: outFace,outFacing,outRow,outColumn = 5,0,49-(inRow%50),0
	}

	//convert back to overallMap (from being onto a face)
	outRow,outColumn = faceToMapCoordinate(outRow,outColumn,outFace)

	return outRow,outColumn,outFace,outFacing
}

func faceToMapCoordinate(row,column,face int) (int,int){
	rowOut,columnOut := 0,0

	switch face {
	case 1: rowOut,columnOut = row%50, column%50 + 100
	case 2: rowOut,columnOut = row%50, column%50 + 50
	case 3: rowOut,columnOut = row%50 + 50, column%50 + 50
	case 4: rowOut,columnOut = row%50 + 100, column%50 + 50
	case 5: rowOut,columnOut = row%50 + 100, column%50
	case 6: rowOut,columnOut = row%50 + 150, column%50
	}

	return rowOut,columnOut
}



func setupProblem(fileName string) (int,int,int,map[int]map[int]string,string){
	readFile, err := os.Open(fileName)
	check(err)
	defer readFile.Close()

    fileScanner := bufio.NewScanner(readFile)
    fileScanner.Split(bufio.ScanLines)

	var instructions string
	instructionsIsNext := false

	overallMap := make(map[int]map[int]string)
	row,sRow,sColumn,sFace := 0,-1,-1,-1
	faces := [6][50][50]string{}
	
	for fileScanner.Scan() {
		line := fileScanner.Text()
		overallMap[row] = make(map[int]string)

		if len(line) == 0 {
			instructionsIsNext = true
			continue
		}

		if instructionsIsNext {
			instructions = line
			break
		} else {
			
			for column,rune := range line {
				symbol := string(rune)
				
				if symbol == " " {
					continue
				}

				overallMap[row][column]=symbol
				face := whichFace(row,column)
				faces[face-1][row%50][column%50]= symbol

				if row == 0 && sRow == -1 { 
					sRow,sColumn,sFace = row,column,face
				}
			}
		}
		row++
	}

	// for i:= 0 ; i < len(faces) ; i++ {
	// 	fmt.Println("Face ",i+1)
	// 	face := faces[i]
	// 	for row := 0 ; row < len(face) ; row ++ {
	// 		line := ""
	// 		for column := 0 ; column < len(face[row]); column ++{
	// 			line += face[row][column]
	// 		}
	// 		fmt.Println(line)
	// 	} 
	// 	fmt.Println()
	// }

	return sRow,sColumn,sFace,overallMap,instructions
}

func whichFace(row,column int) int {
	switch {
		case row < 50 && column >= 100 :  return 1
		case row < 50 && column < 100 :   return 2
		case row >= 50 && row < 100  :     return 3
		case row >=100 && row < 150 && column >= 50 :  return 4
		case row >=100 && row < 150 && column < 50 :  return 5
		case row >= 150:  return 6
		default : {
			fmt.Println("Row ", row, " column ", column)
			panic ("We cannot be here")
		}
	}
}
