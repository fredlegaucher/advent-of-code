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

type Elf struct { 
	row,column,proposedRow,proposedColumn int 
        hasProposed bool
} 

func (e *Elf) getElfCoordinates() string {
	return fmt.Sprintf("%d-%d", e.row , e.column)
}

func main() {
	fmt.Println(solveProblem("input.txt"))
}


func solveProblem(fileName string) int {
	mapOfElves := setupProblem(fileName)

        round := 0 
        for { 
                // Propose 
                for _, e := range mapOfElves {                         
                        e.hasProposed = false //reset here for good measyre
                        if e.areThereNoElvesAround(mapOfElves) { 
                                continue 
                        } else {
                        }

                        directionToConsider := round % 4 
                        first := true 

                        for ; directionToConsider != round%4 || first ; { 
                                first = false 
                                if e.potentialMove(directionToConsider,mapOfElves){ 
                                        break // we have proposed       
                                } 
                                directionToConsider = (directionToConsider + 1) % 4
                        } 
                } 

                // check proposals 
                for i, e1 := range mapOfElves { 
                        same := false 
                        if !e1.hasProposed { 
                                continue 
                        } 
                        for j, e2 := range mapOfElves { 
                                if !e2.hasProposed || i == j { 
                                        continue 
                                } 
 
                                if e1.proposedColumn == e2.proposedColumn && e1.proposedRow == e2.proposedRow { 
                                        e2.hasProposed = false 
                                        same = true 
                                } 
                        } 
                        if same { 
                                e1.hasProposed = false
                        } 
                } 
 
                updatedMapOfElves := map[string]*Elf{} 
 
                // execute proposal 
                moved := false 
                for _, e := range mapOfElves { 
                        if e.hasProposed { 
                                moved = true 
                                e.column = e.proposedColumn 
                                e.row = e.proposedRow 
                        } 
                        e.hasProposed = false
                        if _, ok := updatedMapOfElves[e.getElfCoordinates()]; ok { 
                                panic("we have a bug") 
                        } 
                        updatedMapOfElves[e.getElfCoordinates()] = e 
                } 
                mapOfElves = updatedMapOfElves 
                if !moved { 
                        break 
                } else {
                        round++
                }
        } 

        return round + 1
} 

func (e *Elf) potentialMove(directionToConsider int ,mapOfElves map[string]*Elf) bool{

        isThereAnotherElfInThisDirection := false
        switch directionToConsider { 
                case 0: isThereAnotherElfInThisDirection = e.isThereAnotherElfNorth(mapOfElves) 
                case 1: isThereAnotherElfInThisDirection = e.isThereAnotherElfSouth(mapOfElves) 
                case 2: isThereAnotherElfInThisDirection = e.isThereAnotherElfWest(mapOfElves) 
                case 3: isThereAnotherElfInThisDirection = e.isThereAnotherElfEast(mapOfElves) 
        }
        
        if !isThereAnotherElfInThisDirection{
                switch directionToConsider { 
                        case 0: e.proposedRow,e.proposedColumn = e.row-1,e.column
                        case 1: e.proposedRow,e.proposedColumn = e.row+1,e.column 
                        case 2: e.proposedRow,e.proposedColumn = e.row,e.column-1
                        case 3: e.proposedRow,e.proposedColumn = e.row,e.column+1
                }
                e.hasProposed = true
        }

        return e.hasProposed
} 


func setupProblem(fileName string) map[string]*Elf{
        readFile, err := os.Open(fileName)
        check(err)
        defer readFile.Close()

        fileScanner := bufio.NewScanner(readFile)
        fileScanner.Split(bufio.ScanLines)

        mapOfElves := map[string]*Elf{} 
        row := 0
        for fileScanner.Scan() {
                line := fileScanner.Text()
                if line == "" { 
                        break 
                } 

                for column, r := range line { 
                        if string(r) == "#" { 
                                        e := &Elf{ row:row,column:column } 
                                        mapOfElves[e.getElfCoordinates()] = e 
                        } 
                        
                }
                
                row++
        }

	return mapOfElves
}

func (e *Elf) areThereNoElvesAround(elves map[string]*Elf) bool { 
        _, N := elves[fmt.Sprintf("%d-%d", e.row - 1, e.column )] 
        _, NE := elves[fmt.Sprintf("%d-%d", e.row - 1, e.column +1)] 
        _, NW := elves[fmt.Sprintf("%d-%d", e.row - 1, e.column -1)] 
        _, S := elves[fmt.Sprintf("%d-%d", e.row + 1, e.column)] 
        _, SE := elves[fmt.Sprintf("%d-%d", e.row + 1, e.column+1)] 
        _, SW := elves[fmt.Sprintf("%d-%d", e.row + 1, e.column-1)]
        _, E := elves[fmt.Sprintf("%d-%d", e.row ,   e.column + 1 )] 
        _, W := elves[fmt.Sprintf("%d-%d", e.row , e.column - 1 )] 
        return !N && !NE && !NW && !S && !SE && !SW && !E && !W
} 

 func (e *Elf) isThereAnotherElfNorth(elves map[string]*Elf) bool { 
        _, N := elves[fmt.Sprintf("%d-%d", e.row - 1, e.column )] 
        _, NE := elves[fmt.Sprintf("%d-%d", e.row - 1, e.column +1)] 
        _, NW := elves[fmt.Sprintf("%d-%d", e.row - 1, e.column -1)] 
        return N || NE || NW 
} 

func (e *Elf) isThereAnotherElfSouth(elves map[string]*Elf) bool {   
        _, S := elves[fmt.Sprintf("%d-%d", e.row + 1, e.column)] 
        _, SE := elves[fmt.Sprintf("%d-%d", e.row + 1, e.column+1)] 
        _, SW := elves[fmt.Sprintf("%d-%d", e.row + 1, e.column-1)] 
        return S || SW || SE 
} 

 func (e *Elf) isThereAnotherElfEast(elves map[string]*Elf) bool { 
        _, E := elves[fmt.Sprintf("%d-%d", e.row ,   e.column + 1 )] 
        _, SE := elves[fmt.Sprintf("%d-%d", e.row +1, e.column + 1 )] 
        _, NE := elves[fmt.Sprintf("%d-%d", e.row -1, e.column + 1 )] 
        return E || NE || SE 
} 

func (e *Elf) isThereAnotherElfWest(elves map[string]*Elf) bool { 
        _, W := elves[fmt.Sprintf("%d-%d", e.row , e.column - 1 )] 
        _, NW := elves[fmt.Sprintf("%d-%d", e.row +1, e.column - 1 )] 
        _, SW := elves[fmt.Sprintf("%d-%d", e.row -1, e.column - 1 )] 
        return W || SW || NW 
}
