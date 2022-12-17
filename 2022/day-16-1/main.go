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
	fmt.Println(solveProblem("input.txt"))
}

type Problem struct {
	valves map[string]*Valve //indexed on name
	nonZeroValves map[string]*Valve
	startingValve string
}

type Valve struct {
	name string
	rate int
	neighbours []string 
	distanceToOtherValves map[string]int
}

func solveProblem(fileName string) int {

	p := setupProblem(fileName)

	
	//pre-compute distances between NZ valves
	for _,nzv := range p.nonZeroValves {
		bfs(nzv,p.valves)
	}


	bfs(p.valves[p.startingValve],p.valves)
	return calculateMaxScore(p.valves[p.startingValve],p.nonZeroValves,p,0)	
}

func calculateMaxScore(previousValve *Valve, incomingValves map[string]*Valve,p *Problem, timeUsed int) int {
	score := 0
	for name,valve := range incomingValves {
		time := timeUsed + previousValve.distanceToOtherValves[name] + 1
		timeLeft := max(0,30 - time)

		outgoingValves := make(map[string]*Valve)
		for k,v := range incomingValves {
			if k != name {
				outgoingValves[k]=v
			}
		}

		if (time <= 30){
			localScore := timeLeft * valve.rate 
			if len(outgoingValves) > 0 {
				localScore += calculateMaxScore(valve,outgoingValves,p,time)
			}
			if localScore > score {
				score = localScore
			} 
		}
	}

	return score
}


type NameDistance struct {
	name string
	distance int
}

func bfs(nzv *Valve, valves map[string]*Valve) {
	queue := make([]*NameDistance,0)
	queue = append(queue,&NameDistance{nzv.name,1})
	visitedValves := make(map[string]bool)

	for  ; len(queue) > 0 ; {
		currentItem := queue[0]
		queue = queue[1:]

		for _,neighbour := range valves[currentItem.name].neighbours {
			_,ok := visitedValves[neighbour]
			
			if ok { //already visited
				continue
			}
			
			visitedValves[neighbour]=true
			
			if valves[neighbour].rate > 0 {
				nzv.distanceToOtherValves[neighbour] = currentItem.distance
			}
			
			queue = append(queue,&NameDistance{neighbour,currentItem.distance + 1})
		}
	}

	
}

func (v *Valve) print() {
	
	fmt.Println(fmt.Sprintf("%v with rate %v and neighbours %v",v.name,v.rate,v.neighbours))
	for k,v := range v.distanceToOtherValves {
		fmt.Println(fmt.Sprintf("Distance to %v is %v",k,v))
	}
	
}


func (p *Problem) printInit(){
	fmt.Println("Valves")
	for _,v := range p.valves {
		if len(v.distanceToOtherValves) > 0 {
			v.print()
		}
	}

	fmt.Println("Starting valve",p.startingValve)
}

func setupProblem(fileName string) *Problem{
	//parse file
	re := regexp.MustCompile(`Valve (?P<name>.{2}) has flow rate=(?P<rate>\d*); tunnels{0,1} leads{0,1} to valves{0,1} (?P<valvesAsCsv>.*)`)

	readFile, err := os.Open(fileName)
	check(err)
	defer readFile.Close()

    fileScanner := bufio.NewScanner(readFile)
    fileScanner.Split(bufio.ScanLines)

	nonZeroValves := make(map[string]*Valve)
	valves := make(map[string]*Valve)
	p := &Problem{nonZeroValves: nonZeroValves, valves:valves, startingValve: "AA"}


	
	for fileScanner.Scan() {
		s := re.FindAllStringSubmatch(fileScanner.Text(), -1)[0][1:]
		name:=    s[0]
		rate,_ :=      strconv.Atoi(s[1])
		valvesAsCsv :=      s[2]

		distanceToOtherValves := make(map[string]int)
		p.valves[name] = &Valve{name,rate,strings.Split(valvesAsCsv,", "),distanceToOtherValves}

		if rate > 0 {
			p.nonZeroValves[name]=p.valves[name]
		}
	}

	return p
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











