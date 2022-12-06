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
		result = append(result,findStarterPacket(line))
	}

	return result
}

func findStarterPacket(line string) string {

	i :=0
	for ; i < len(line) ; i++{
		currentPacket := line[i:i+4]
		if  currentPacket[0] == currentPacket[1] ||
			currentPacket[0] == currentPacket[2] ||
			currentPacket[0] == currentPacket[3] ||
			currentPacket[1] == currentPacket[2] || 
			currentPacket[1] == currentPacket[3] || 
			currentPacket[2] == currentPacket[3]{
				continue
			}
		break
	}

	return fmt.Sprint(i + 4)
}