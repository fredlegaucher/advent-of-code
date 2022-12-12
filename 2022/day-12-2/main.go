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
	//set up map and track where S and E are
	p := setupProblem(fileName)
	fmt.Printf("%v",p)
	//traverse the map from S bfs
	for ; p.solution == nil;{
		currentPositionToBeVisited := p.toBeVisited[0]
		p.toBeVisited = p.toBeVisited[1:len(p.toBeVisited)]
	 	p.travel(currentPositionToBeVisited)
	}

	return p.solution.movesToGetHere
}



func (p *Problem) travel(position *Position) {
	Pr := position.row
	Pc := position.column
	moves := position.movesToGetHere + 1

	if p.carte[Pr][Pc] == "a" { // we have arrived!
		p.solution = position
	}

	alreadyVisited := p.visitPosition(Pr,Pc) 
		
	if !alreadyVisited {

		if Pr > 0 && p.carte[Pr][Pc][0] -1 <= p.carte[Pr-1][Pc][0] { //can go up
			p.toBeVisited = append(p.toBeVisited,&Position{row:Pr-1,column:Pc,movesToGetHere: moves})
		}

		if Pr < p.height - 1 && p.carte[Pr][Pc][0] - 1 <= p.carte[Pr+1][Pc][0]  { //can go down
			p.toBeVisited = append(p.toBeVisited,&Position{row:Pr+1,column:Pc,movesToGetHere: moves})
		}

		if Pc > 0 && p.carte[Pr][Pc][0]  -1 <= p.carte[Pr][Pc-1][0] { //can go left
			p.toBeVisited = append(p.toBeVisited,&Position{row:Pr,column:Pc-1,movesToGetHere: moves})
		}

		if Pc < p.width - 1 && p.carte[Pr][Pc][0] -1 <= p.carte[Pr][Pc+1][0] { //can go right
			p.toBeVisited = append(p.toBeVisited,&Position{row:Pr,column:Pc+1,movesToGetHere: moves})
		}
	}
}

type Problem struct {
	carte [][]string 
	Sr,Sc,Er,Ec int
	height int
	width int
	visitedLocations []bool
	toBeVisited []*Position
	solution *Position
}

type Position struct {
	row, column, movesToGetHere int
}

func (p *Problem) visitPosition(row,column int) (alreadyVisited bool) {
	if p.visitedLocations[row * p.width + column] {
		return true
	}

	p.visitedLocations[row * p.width + column] = true
	return false
}

func (p *Problem) String() string{
	result := ""
	for _, row := range p.carte {
		for _,item := range row {
			result += item
		}
		result += "\n"
	}

	result += fmt.Sprintf("Start is in position [column:%v,row:%v] \n",p.Sc,p.Sr)
	result += fmt.Sprintf("End   is in position [column:%v,row:%v] \n",p.Ec,p.Er)
	result += fmt.Sprintf("Map height is %v \n",p.height)
	result += fmt.Sprintf("Map width  is %v \n",p.width)
	return result
}

func setupProblem(fileName string) *Problem{
	result := &Problem{}

	readFile, err := os.Open(fileName)
	defer readFile.Close()
  
    if err != nil {
        fmt.Println(err)
    }
    fileScanner := bufio.NewScanner(readFile)
    fileScanner.Split(bufio.ScanLines)


	r := 0
	carte := make([][]string,0)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		
		newRow := make([]string,0)
		for c,symbol := range line {
			letter := string(symbol)
			
			if string(symbol) == "E" {
				result.Er = r
				result.Ec = c
				letter = "z"
			}

			if string(symbol) == "S" {
				result.Sr = r
				result.Sc = c
				letter = "a"
			}

			newRow = append(newRow,letter)
		}
		carte = append(carte,newRow)
		r++
	}
	
	result.carte = carte
	result.height = len(carte)
	result.width = len(carte[0])
	result.visitedLocations = make([]bool,result.height * result.width)
	result.toBeVisited = append(result.toBeVisited, &Position{row:result.Er,column:result.Ec,movesToGetHere:0})
	return result
}







