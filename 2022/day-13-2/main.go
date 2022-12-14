package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)


func main() {
	fmt.Println(solveProblem("input.txt"))
}

func mergeSort(items []string) []string {
    if len(items) < 2 {
        return items
    }
    first := mergeSort(items[:len(items)/2])
    second := mergeSort(items[len(items)/2:])
    return merge(first, second)
}

func merge(a []string, b []string) []string {
    final := []string{}
    i := 0
    j := 0
    for i < len(a) && j < len(b) {
        if isLhsSmallerThanRhs(a[i],b[j]) {
            final = append(final, a[i])
            i++
        } else {
            final = append(final, b[j])
            j++
        }
    }
    for ; i < len(a); i++ {
        final = append(final, a[i])
    }
    for ; j < len(b); j++ {
        final = append(final, b[j])
    }
    return final
}

func solveProblem(fileName string) int {
	readFile, err := os.Open(fileName)
	defer readFile.Close()
  
    if err != nil {
        fmt.Println(err)
    }
    fileScanner := bufio.NewScanner(readFile)
    fileScanner.Split(bufio.ScanLines)

	
	lines := make([]string,0)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if len(line) == 0 {
			continue
		}
		lines = append(lines, line)
	}

	sortedLines := mergeSort(lines)

	indexFor2 := 0
	indexFor6 := 0
	for i,line := range sortedLines {
		if line == "[[2]]" {
			indexFor2 = i + 1
		}
		if line == "[[6]]"{
			indexFor6 = i + 1
		}
		
	}
	
	return indexFor2 * indexFor6
}

func isLhsSmallerThanRhs(leftLine, rightLine string) bool {
	leftElements := make([]*Element,0)
	leftPaquet := &Paquet{elements:leftElements,level:0}
	parseLine(leftLine[1:len(leftLine)-1],leftPaquet)
	fmt.Println(leftPaquet)
	
	rightElements := make([]*Element,0)
	rightPaquet := &Paquet{elements:rightElements,level:0}
	parseLine(rightLine[1:len(rightLine)-1],rightPaquet)
	fmt.Println(rightPaquet)

	return compareElements(leftPaquet.elements, rightPaquet.elements) == -1
}

type Paquet struct {
	elements []*Element
	level    int
}

func (p *Paquet) String() string{
	result := "["
	for i,e := range p.elements {
		if i == 0 {
			result += e.String()  
			continue
		}
		result += "," + e.String()  
	}
	result += "]"
	return result
}

type Element struct {
	value int // using -1 to encode nothing being set. Booh.
	paquet *Paquet
}

func (e *Element) String() string{
	if e.value >= 0 {
		return fmt.Sprintf("%v", e.value)
	}

	return e.paquet.String()
}


func parseLine(line string, parentPaquet *Paquet) string {
	remainingCharacters := line
	for ; len(remainingCharacters) > 0; {
		currentCharacter := string(remainingCharacters[0])

		if currentCharacter == "[" {
			elements := make([]*Element,0)
			newPaquet:= &Paquet{elements:elements, level: parentPaquet.level + 1}
			parentPaquet.elements = append(parentPaquet.elements, &Element{paquet:newPaquet,value:-1} )
			remainingCharacters = parseLine(remainingCharacters[1:],newPaquet)
			continue
		}

		if currentCharacter == "]" {
			return remainingCharacters[1:]	
		}

		if currentCharacter == "," {
			remainingCharacters = remainingCharacters[1:]	
			continue
		}

		//integer - althought there are some 10 so be careful!
		lengthOfInteger := 1
		if len(remainingCharacters) > 1 {
			for nextCharacter := string(remainingCharacters[lengthOfInteger]) ; nextCharacter != "]" && nextCharacter != ","; nextCharacter = string(remainingCharacters[lengthOfInteger]) {
				lengthOfInteger++
			}
		}
		value,_ := strconv.Atoi(string(remainingCharacters[0:lengthOfInteger]))
		parentPaquet.elements = append(parentPaquet.elements, &Element{value:value} )
		remainingCharacters = remainingCharacters[lengthOfInteger:]	
		}
		
		return ""
}


func compareElements(leftElements, rightElements []*Element) int {
	
	if len(leftElements) == 0 && len(rightElements) > 0 {
		return -1
	}

	if len(leftElements) > 0 && len(rightElements) == 0 {
		return 1
	}

	if len(leftElements) == 0 && len(rightElements) == 0 {
		return 0
	}
	//both have an element at least
	if leftElements[0].value > -1 && rightElements[0].value > -1 {
		comparison := compareInt(leftElements[0].value,rightElements[0].value)
		if comparison == 0 {
			return compareElements(leftElements[1:],rightElements[1:])
		}
		return comparison
	}

	if leftElements[0].value > -1 && rightElements[0].value == -1 {
		wrapupValueIntoPaquet(leftElements[0])
		return compareElements(leftElements,rightElements)
	}
	if leftElements[0].value == -1 && rightElements[0].value > -1 {
		wrapupValueIntoPaquet(rightElements[0])
		return compareElements(leftElements,rightElements)
	}


	if leftElements[0].value == -1 && rightElements[0].value == -1 {
		comparison := compareElements(leftElements[0].paquet.elements,rightElements[0].paquet.elements)
		if comparison == 0 {
			return compareElements(leftElements[1:],rightElements[1:])
		}
		return comparison
	}
	
	panic("We should not be here")

}

func compareInt(lhs, rhs int) int {
	if lhs < rhs {
		return -1
	}
	if lhs > rhs {
		return 1
	}
	return 0
}

func wrapupValueIntoPaquet(element *Element){
	newPaquet := &Paquet{}
	newElement := &Element{value:element.value}
	newPaquet.elements = []*Element{newElement}
	element.value = -1
	element.paquet = newPaquet
}







