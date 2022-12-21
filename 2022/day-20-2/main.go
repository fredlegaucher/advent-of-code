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

type LinkedListItem struct { 
	value  int 
	nextItem *LinkedListItem 
	previousItem *LinkedListItem 
} 

func main() {
	a,b,c := solveProblem("input.txt")
	fmt.Println(a+b+c)
}


func mixing(linkedList []*LinkedListItem) {
	for _, current := range linkedList { 
		offset := current.value % (len(linkedList) - 1) 

		if offset == 0 { 
				continue 
		} else if offset < 0 { 
				// Remove element 
				current.previousItem.nextItem = current.nextItem 
				current.nextItem.previousItem = current.previousItem 

				// Go backwards 
				insert := current 
				for dx := 0; dx < -offset; dx++ { 
						insert = insert.previousItem 
				} 

				// Insert before 
				insert.previousItem.nextItem = current 
				current.previousItem = insert.previousItem 
				current.nextItem = insert 
				insert.previousItem = current 
		} else { 
				// Remove element 
				current.previousItem.nextItem = current.nextItem 
				current.nextItem.previousItem = current.previousItem 

				// Go forwards 
				insert := current 
				for dx := 0; dx < offset; dx++ { 
						insert = insert.nextItem 
				} 

				// Insert after 
				insert.nextItem.previousItem = current 
				current.nextItem = insert.nextItem 
				insert.nextItem = current 
				current.previousItem = insert 
		} 
	} 
}

func solveProblem(fileName string) (int,int,int) {

	linkedList := setupProblem(fileName)
	
	for i:= 0 ; i < 10 ; i++{
		mixing(linkedList)
	}

	var zeroItem *LinkedListItem
	for zeroItem = linkedList[0] ; zeroItem.value != 0 ; zeroItem = zeroItem.nextItem {
	}
			
	currentItem := zeroItem
	var thousandth,twoThousandth,threeThousdands int
	for i := 0; i <= 3000; i++ { 
		if i == 1000  { 
			thousandth = currentItem.value
		} 
		if i == 2000  { 
			twoThousandth = currentItem.value
		} 
		if i == 3000  { 
			threeThousdands = currentItem.value
		} 
		currentItem = currentItem.nextItem
	} 

	return thousandth,twoThousandth,threeThousdands
}


func setupProblem(fileName string) []*LinkedListItem {
	readFile, err := os.Open(fileName)
	check(err)
	defer readFile.Close()

    fileScanner := bufio.NewScanner(readFile)
    fileScanner.Split(bufio.ScanLines)

	var current *LinkedListItem 
	result := []*LinkedListItem{}
	
	for fileScanner.Scan() {
		value,_ := strconv.Atoi(fileScanner.Text())
		nextItem := &LinkedListItem{value: value, previousItem: current} 
		result = append(result, nextItem)
		if current != nil {
			current.nextItem = nextItem 
		} 
		current = nextItem 
	}
	//loop the loop
	current.nextItem = result[0]
	result[0].previousItem = current 

	for i:=0 ; i < len(result); i++ {
		result[i].value = result[i].value * 811589153
	}
	return result 	
}
