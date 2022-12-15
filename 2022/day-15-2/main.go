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
	fmt.Println(solveProblem("input.txt",4000000))
}

type Problem struct {
	sensors []*Sensor
	coverage map[int]map[int]string
	maxX,maxY,minX,minY int
}

type Sensor struct {
	Sx,Sy,Bx,By,distance int
}

func solveProblem(fileName string, maxXorY int) int {

	p := setupProblem(fileName)
	
	currentY:=0
	currentX := 0
	found := false
	for ; currentY<= maxXorY ; currentY++ {
		fmt.Println("Starting ", currentY)
		
		for currentX = 0 ; currentX <= maxXorY; {
			nextX := currentX
			for _,s := range p.sensors {
				if abs(s.Sx - currentX) + abs(s.Sy - currentY) <= s.distance {
					nextX =s.Sx + s.distance - abs(s.Sy - currentY) + 1 //jumpy jump
					break
				}
			}
			
			if currentX == nextX {
				found = true
				break
			}
			
			currentX = nextX
		}  

	
		if found {
			break
		}
		
	}


	return  currentX * 4000000 + currentY
}


func (p *Problem) print(){
	f, err := os.Create("output.txt")
	check(err)

	for y, row := range p.coverage {
		result := fmt.Sprintf("y=%v ",y)
		for x,item := range row {
			result += fmt.Sprintf(" x=%v %v",x,item)
		}
		_, err := f.WriteString(result+"\n")
		check(err)
		
	}
	f.Sync()
}



func setupProblem(fileName string) *Problem{
	sensors := make([]*Sensor,0)
	coverage := make(map[int]map[int]string)
	var maxX,maxY,minX,minY int
	

	//parse file
	re := regexp.MustCompile(`Sensor at x=(?P<Sx>-?\d+), y=(?P<Sy>-?\d+): closest beacon is at x=(?P<Bx>-?\d+), y=(?P<By>-?\d+)`)

	readFile, err := os.Open(fileName)
	check(err)
	defer readFile.Close()

    fileScanner := bufio.NewScanner(readFile)
    fileScanner.Split(bufio.ScanLines)


	//Sensor at x=2, y=18: closest beacon is at x=-2, y=15
	firstLine := true
	for fileScanner.Scan() {
		s := re.FindAllStringSubmatch(fileScanner.Text(), -1)[0][1:]
		Sx,_ :=      strconv.Atoi(s[0])
		Sy,_ :=      strconv.Atoi(s[1])
		Bx,_ :=      strconv.Atoi(s[2])
		By,_ :=      strconv.Atoi(s[3])
		distance := abs(Sx-Bx) + abs(Sy-By)
		sensors = append(sensors,&Sensor{Sx,Sy,Bx,By,distance})	
		if firstLine || maxX < max(Sx,Bx) {
			maxX = max(Sx,Bx)
		}
		if firstLine || maxY < max(Sy,By) {
			maxY = max(Sy,By)
		}
		if firstLine || minX > min(Sx,Bx) {
			minX = min(Sx,Bx)
		}
		if firstLine || minY > min(Sy,By) {
			minY = min(Sy,By)
		}
		firstLine = false
	}

	for y :=minY; y <= maxY ; y++ {
		coverage[y]=make(map[int]string)
	} 

	for _,s := range sensors {
		coverage[s.Sy][s.Sx]="S"
		coverage[s.By][s.Bx]="B"
	}
	
	return &Problem{sensors,coverage,maxX,maxY,minX,minY}
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











