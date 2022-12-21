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

type Blueprint struct { 
	id int
	OreCost      int 
	ClayCost     int 
	ObsidianCost ObsidianCost 
	GeodeCost    GeodeCost 
	maxOreCost,maxClayCost,maxObsCost int
} 

type ObsidianCost struct { 
	OreCost  int 
	ClayCost int 
} 

type GeodeCost struct { 
	OreCost      int 
	ObsidianCost int 
} 




type GameState struct { 
	time, ores, clays, obsidians, oreRobots, clayRobots, obsidianRobots, geodeRobots int 
} 

func (d GameState) String() string { 
	return fmt.Sprintf("%d-res=[%d, %d, %d]-rob=[%d, %d, %d, %d]", d.time, d.ores, d.clays, d.obsidians, d.oreRobots, d.clayRobots, d.obsidianRobots, d.geodeRobots) 
} 

func main() {
	fmt.Println(solveProblem("input.txt"))
}

func solveProblem(fileName string) int {

	bps := setupProblem(fileName)
	overallBpQuality := 0 
	for _, b := range bps { 
		fmt.Println("Started bp:",b.id)
		gameStateCache := make(map[GameState]int)
		overallBpQuality += b.id * calculateMaxNumberOfGeodes(*b, GameState{24, 0, 0, 0, 1, 0, 0, 0}, gameStateCache,false,false,false,false) 
	} 
	return overallBpQuality
}


func setupProblem(fileName string) []*Blueprint {
	//parse file
	re := regexp.MustCompile(`Blueprint (?P<bpid>\d+): Each ore robot costs (?P<oreRobInOre>\d+) ore. Each clay robot costs (?P<clayRobInOre>\d+) ore. Each obsidian robot costs (?P<obsRobInOre>\d+) ore and (?P<obsRobInClay>\d+) clay. Each geode robot costs (?P<geodeRobInOre>\d+) ore and (?P<geodeRobInObs>\d+) obsidian.`)

	readFile, err := os.Open(fileName)
	check(err)
	defer readFile.Close()

    fileScanner := bufio.NewScanner(readFile)
    fileScanner.Split(bufio.ScanLines)

	blueprints := make([]*Blueprint,0)

	for fileScanner.Scan() {
		s := re.FindAllStringSubmatch(fileScanner.Text(), -1)[0][1:]
		id,_ := strconv.Atoi(s[0])
		oreRobInOre,_:=    strconv.Atoi(s[1])
		clayRobInOre,_:=    strconv.Atoi(s[2])
		obsRobInOre,_:=    strconv.Atoi(s[3])
		obsRobInClay,_:=    strconv.Atoi(s[4])
		geodeRobInOre,_:=    strconv.Atoi(s[5])
		geodeRobInObs,_:=    strconv.Atoi(s[6])

		maxOreCost := max(max(max(oreRobInOre,clayRobInOre),obsRobInOre),geodeRobInOre)
		maxClayCost := obsRobInClay
		maxObsCost := geodeRobInObs

		blueprints  = append(blueprints, &Blueprint{
			id:id, 
			OreCost:  oreRobInOre, 
			ClayCost: clayRobInOre, 
			ObsidianCost: ObsidianCost{ 
					OreCost:  obsRobInOre, 
					ClayCost: obsRobInClay, 
			}, 
			GeodeCost: GeodeCost{ 
					OreCost:      geodeRobInOre, 
					ObsidianCost: geodeRobInObs, 
			},maxOreCost:maxOreCost,maxClayCost:maxClayCost,maxObsCost:maxObsCost,
		})
	}

	return blueprints
}




func calculateMaxNumberOfGeodes(b Blueprint, k GameState, gameStateCache map[GameState]int,couldHaveBuildOreRob,couldHaveBuildClayRob,couldHaveBuildObsRob,couldHaveBuildGeodeRob bool) int { 
	if val, ok := gameStateCache[k]; ok { 
			return val 
	} 
	if k.time == 0 { 
			return 0 
	} 
	
	if k.time == 1 { 
		return k.geodeRobots
	} 

	maxGeodeRobots := k.geodeRobots 

	if k.ores >= b.GeodeCost.OreCost && k.obsidians >= b.GeodeCost.ObsidianCost && !couldHaveBuildGeodeRob { 
	//geode rob
		maxGeodeRobots = max(maxGeodeRobots, k.geodeRobots+calculateMaxNumberOfGeodes(b, GameState{ 
				k.time - 1, 
				k.ores + k.oreRobots - b.GeodeCost.OreCost, k.clays + k.clayRobots, k.obsidians + k.obsidianRobots - b.GeodeCost.ObsidianCost, 
				k.oreRobots, k.clayRobots, k.obsidianRobots, k.geodeRobots + 1, 
		},gameStateCache,false,false,false,false)) 
	} 
	
	if k.ores >= b.ObsidianCost.OreCost && 
		k.clays >= b.ObsidianCost.ClayCost &&
	    k.obsidianRobots < b.maxObsCost &&
		k.time * b.maxObsCost >= k.obsidianRobots &&
		!couldHaveBuildObsRob {
		
	//obsidian rob
		maxGeodeRobots = max(maxGeodeRobots, k.geodeRobots+calculateMaxNumberOfGeodes(b, GameState{ 
				k.time - 1, 
				k.ores + k.oreRobots - b.ObsidianCost.OreCost, k.clays + k.clayRobots - b.ObsidianCost.ClayCost, k.obsidians + k.obsidianRobots, 
				k.oreRobots, k.clayRobots, k.obsidianRobots + 1, k.geodeRobots, 
		},gameStateCache,false,false,false,false)) 

	} 

	if k.ores >= b.ClayCost && k.clayRobots < b.maxClayCost && k.time * b.maxClayCost >= k.clayRobots && !couldHaveBuildClayRob{ 
	//clay rob
		maxGeodeRobots = max(maxGeodeRobots, k.geodeRobots+calculateMaxNumberOfGeodes(b, GameState{ 
				k.time - 1, 
				k.ores + k.oreRobots - b.ClayCost, k.clays + k.clayRobots, k.obsidians + k.obsidianRobots, 
				k.oreRobots, k.clayRobots + 1, k.obsidianRobots, k.geodeRobots, 
		},gameStateCache,false,false,false,false)) 
	} 

	// Option: make ore robot 
	if k.ores >= b.OreCost && k.oreRobots < b.maxOreCost && k.time * b.maxOreCost >= k.oreRobots && !couldHaveBuildOreRob { 
			maxGeodeRobots = max(maxGeodeRobots, k.geodeRobots+calculateMaxNumberOfGeodes(b, GameState{ 
					k.time - 1, 
					k.ores + k.oreRobots - b.OreCost, k.clays + k.clayRobots, k.obsidians + k.obsidianRobots, 
					k.oreRobots + 1, k.clayRobots, k.obsidianRobots, k.geodeRobots, 
			},gameStateCache,false,false,false,false)) 
	} 

	//do nothing
	maxOres := max(max(max(b.OreCost, b.ClayCost), b.ObsidianCost.OreCost), b.GeodeCost.OreCost)
	if k.ores < maxOres || k.clays < b.ClayCost || k.obsidians < b.GeodeCost.ObsidianCost { 
			couldHaveBuildOreRob = k.ores >= b.OreCost
			couldHaveBuildClayRob = k.ores >= b.ClayCost
			couldHaveBuildObsRob = k.ores >= b.ObsidianCost.OreCost && k.clays >= b.ObsidianCost.ClayCost
			couldHaveBuildGeodeRob = k.ores >= b.GeodeCost.OreCost && k.obsidians > b.GeodeCost.ObsidianCost



			maxGeodeRobots = max(maxGeodeRobots, k.geodeRobots+calculateMaxNumberOfGeodes(b, GameState{ 
					k.time - 1, 
					k.ores + k.oreRobots, k.clays + k.clayRobots, k.obsidians + k.obsidianRobots, 
					k.oreRobots, k.clayRobots, k.obsidianRobots, k.geodeRobots, 
			},gameStateCache,couldHaveBuildOreRob,couldHaveBuildClayRob,couldHaveBuildObsRob,couldHaveBuildGeodeRob)) 
	} 
	

	gameStateCache[k] = maxGeodeRobots 
	return maxGeodeRobots 
} 

func max(a,b int) int {
	if a > b {
		return a
	}
	return b
}