package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
	indexToStackMap := findCrateStartingArrangement("input.txt")
	fmt.Println(solveProblem("input.txt", indexToStackMap))
}

func solveProblem(fileName string, indexToStackMap map[int][]string) string {
	r := regexp.MustCompile(`move (?P<numberOfCrates>\d+) from (?P<sourceStack>\d+) to (?P<targetStack>\d+)`)
	const NUMBER_OF_CRATES string = "numberOfCrates"
	const SOURCE_STACK string = "sourceStack"
	const TARGET_STACK string = "targetStack"

	readFile, err := os.Open(fileName)
	defer readFile.Close()
  
    if err != nil {
        fmt.Println(err)
    }
    fileScanner := bufio.NewScanner(readFile)
 
    fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
        line := fileScanner.Text()

		if !strings.HasPrefix(line,"move") || len(line) == 0 {
			continue
		}

		matches := r.FindStringSubmatch(line)
		numberOfCratesToMove,_ := strconv.Atoi(matches[r.SubexpIndex(NUMBER_OF_CRATES)])
		indexOftackToMoveFrom,_ := strconv.Atoi(matches[r.SubexpIndex(SOURCE_STACK)])
		indexOfstackToMoveTo,_ := strconv.Atoi(matches[r.SubexpIndex(TARGET_STACK)])

		stackToMoveFrom := indexToStackMap[indexOftackToMoveFrom]
		stackToMoveTo := indexToStackMap[indexOfstackToMoveTo]
		
		stackToMoveTo = append(stackToMoveTo, stackToMoveFrom[len(stackToMoveFrom) - numberOfCratesToMove:]...)
		stackToMoveFrom = stackToMoveFrom[:len(stackToMoveFrom) - numberOfCratesToMove]
	
		indexToStackMap[indexOftackToMoveFrom] = stackToMoveFrom
		indexToStackMap[indexOfstackToMoveTo] = stackToMoveTo
	}

	result := ""
	for i := 1; i <= len(indexToStackMap); i++ {
		result += indexToStackMap[i][len(indexToStackMap[i])-1]
	}

	return result
}

func findCrateStartingArrangement(fileName string) map[int][]string {
	//TODO parse here
	firstStack := []string{"D","B","J","V"}
	secondStack := []string{"P","V","B","W","R","D","F"}
	thirdStack := []string{"R","G","F","L","D","C","W","Q"}
	fourthStack := []string{"W","J","P","M","L","N","D","B"}
	fifthStack := []string{"H","N","B","P","C","S","Q"}
	sixthStack := []string{"R","D","B","S","N","G"}
	seventhStack := []string{"Z","B","P","M","Q","F","S","H"}
	eightStack := []string{"W","L","F"}
	ninethStack := []string{"S","V","F","M","R"}

	indexToStackMap := make(map[int][]string)
	indexToStackMap[1] = firstStack
	indexToStackMap[2] = secondStack
	indexToStackMap[3] = thirdStack
	indexToStackMap[4] = fourthStack
	indexToStackMap[5] = fifthStack
	indexToStackMap[6] = sixthStack
	indexToStackMap[7] = seventhStack
	indexToStackMap[8] = eightStack
	indexToStackMap[9] = ninethStack

	return indexToStackMap
}

