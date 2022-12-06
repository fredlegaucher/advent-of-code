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
	fmt.Println(solveProblem("input.txt")[0])
}

func solveProblem(fileName string) []string {
	readFile, err := os.Open(fileName)
	defer readFile.Close()
  
    if err != nil {
        fmt.Println(err)
    }
    fileScanner := bufio.NewScanner(readFile)
    fileScanner.Split(bufio.ScanLines)

	result := []string{}

	for fileScanner.Scan() {
        line := fileScanner.Text()
		result = append(result,findMessage(line))
	}

	return result
}

func findMessage(line string) string {

	i :=0
	for ; i < len(line) ; i++{
		currentPacket := line[i:i+14]
		letterMap := make(map[byte]int)
		duplicates := false
		for j := 0 ; j < 14 ; j++ {
			_,ok := letterMap[currentPacket[j]]
			if ok {
				duplicates = true
				break
			}
			letterMap[currentPacket[j]]=1
		}

		if !duplicates {
			break
		}
	}

	return fmt.Sprint(i + 14)
}