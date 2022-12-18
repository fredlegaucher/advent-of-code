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

	type CC struct {
		x,y,z int
	}

	func (c *CC) String() string {
		return fmt.Sprintf("%v-%v-%v",c.x,c.y,c.z)
	}

	func solveProblem(fileName string) int {

		//parse file
		re := regexp.MustCompile(`(?P<x>\d+),(?P<y>\d+),(?P<z>\d+)`)

		readFile, err := os.Open(fileName)
		check(err)
		defer readFile.Close()

		fileScanner := bufio.NewScanner(readFile)
		fileScanner.Split(bufio.ScanLines)

		lavaCubeCentres := make(map[string]*CC) // cube also modelled as x-y-z
		maxX,maxY,maxZ,minX,minY,minZ := 0,0,0,0,0,0
		firstRow:= true
		for fileScanner.Scan() {
			s := re.FindAllStringSubmatch(fileScanner.Text(), -1)[0][1:]
			x,_:=       strconv.Atoi(s[0])
			y,_ :=      strconv.Atoi(s[1])
			z,_ :=      strconv.Atoi(s[2])
			x*=10
			y*=10
			z*=10

			newCube := &CC{x,y,z}
			lavaCubeCentres[newCube.String()]=newCube

			if firstRow {
				maxX = x
				maxY = y
				maxZ = z
				minX = x
				minY = y
				minZ = z
				firstRow = false
			} else {
				maxX = max(x,maxX)
				maxY = max(y,maxY)
				maxZ = max(z,maxZ)
				minX = min(x,minX)
				minY = min(y,minY)
				minZ = min(z,minZ)
			}

	
		}


		minX-=10
		minY-=10
		minZ-=10
		maxX+=10
		maxY+=10
		maxZ+=10

		
		//where did the water go
		waterCubes := bfs(minX,minY,minZ,maxX,maxY,maxZ,lavaCubeCentres) 

		//find the cubes that are not water
		notWaterCubes := make(map[string]*CC)
		for x:=minX ; x <= maxX ; x+=10 {
			for y:=minY ; y <= maxY ; y+=10 {
				for z:=minZ ; z <= maxZ ; z+=10 {
					key := fmt.Sprintf("%v-%v-%v",x,y,z)
					_,water := waterCubes[key]
					if !water  {
						notWaterCubes[key]=&CC{x,y,z}
					}
				}
			}
		}

		combinedFaceCentres := make(map[string]int) // faces modelled as x-y-z
		for _,cube := range notWaterCubes {
			
			x:=cube.x
			y:=cube.y
			z:=cube.z
		
			//x
			leftSide := fmt.Sprintf("%v-%v-%v",x-5,y,z)
			rightSide := fmt.Sprintf("%v-%v-%v",x+5,y,z)
			//y
			frontSide := fmt.Sprintf("%v-%v-%v",x,y-5,z)
			backSide := fmt.Sprintf("%v-%v-%v",x,y+5,z)
			//z
			topSide := fmt.Sprintf("%v-%v-%v",x,y,z+5)
			bottomSide := fmt.Sprintf("%v-%v-%v",x,y,z-5)

			combinedFaceCentres[leftSide]++
			combinedFaceCentres[rightSide]++
			combinedFaceCentres[frontSide]++
			combinedFaceCentres[backSide]++
			combinedFaceCentres[topSide]++
			combinedFaceCentres[bottomSide]++
		}
		
		part2Count := 0	
		for _,v := range combinedFaceCentres {
			if v == 1 {
				part2Count++
			}
		}

		return part2Count
	}


	func (c *CC) neighbours() []*CC{
		neighbours := make([]*CC,0)
		neighbours = append(neighbours,&CC{c.x-10,c.y,c.z})
		neighbours = append(neighbours, &CC{c.x+10,c.y,c.z})
		neighbours = append(neighbours,&CC{c.x,c.y-10,c.z})
		neighbours = append(neighbours,&CC{c.x,c.y+10,c.z})
		neighbours = append(neighbours, &CC{c.x,c.y,c.z-10})
		neighbours = append(neighbours,&CC{c.x,c.y,c.z+10})
		return neighbours
	}

	func bfs(minX,minY,minZ,maxX,maxY,maxZ int, cubeCentres map[string]*CC) map[string]bool {
		queue := make([]*CC,0)
		queue = append(queue, &CC{minX,minY,minZ})
		waterCubes := make(map[string]bool)

		for  ; len(queue) > 0 ; {
			currentCube := queue[0]
			queue = queue[1:]
			currentCubeKey := currentCube.String()

			_,currentKubeAlreadyVisited := waterCubes[currentCubeKey]
			if currentKubeAlreadyVisited {
				continue
			}
			waterCubes[currentCubeKey]=true

			for _,n := range currentCube.neighbours(){
				if n.x <= maxX && n.x >= minX && n.y >= minY && n.y <= maxY && n.z <= maxZ && n.z >= minZ {
					_,isDroplet := cubeCentres[n.String()]
					if !isDroplet{
						queue = append(queue, n)
					}
				}
			}
		}
		
		return waterCubes
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