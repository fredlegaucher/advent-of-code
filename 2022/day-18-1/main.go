package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
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
	//parse file
	re := regexp.MustCompile(`(?P<x>\d+),(?P<y>\d+),(?P<z>\d+)`)

	readFile, err := os.Open(fileName)
	check(err)
	defer readFile.Close()

    fileScanner := bufio.NewScanner(readFile)
    fileScanner.Split(bufio.ScanLines)

	facesCentres := make(map[string]int) // faces modelled as x-y-zz
	for fileScanner.Scan() {
		s := re.FindAllStringSubmatch(fileScanner.Text(), -1)[0][1:]
		x,_:=       strconv.Atoi(s[0])
		y,_ :=      strconv.Atoi(s[1])
		z,_ :=      strconv.Atoi(s[2])
		x*=10
		y*=10
		z*=10

		//x
		leftSide := fmt.Sprintf("%v-%v-%v",x-5,y,z)
		rightSide := fmt.Sprintf("%v-%v-%v",x+5,y,z)
		//y
		frontSide := fmt.Sprintf("%v-%v-%v",x,y-5,z)
		backSide := fmt.Sprintf("%v-%v-%v",x,y+5,z)
		//z
		topSide := fmt.Sprintf("%v-%v-%v",x,y,z+5)
		bottomSide := fmt.Sprintf("%v-%v-%v",x,y,z-5)

		facesCentres[leftSide]++
		facesCentres[rightSide]++
		facesCentres[frontSide]++
		facesCentres[backSide]++
		facesCentres[topSide]++
		facesCentres[bottomSide]++
	}

	count := 0	
	for _,v := range facesCentres {
		if v == 1 {
			count++
		}
	}

	return count
}
