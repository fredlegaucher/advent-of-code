package main

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}


type Monkey struct {
	name string
	yelling bool
	resolved bool
	number int
	op string
	lhsName,rhsName string
	lhsMonkey,rhsMonkey *Monkey
}

func solveProblem(fileName string, yell int) (int,int) {
	tribe := setupProblem(fileName)
 
	root := tribe["root"]
	
	tribe["humn"].number = yell
	resolveMonkey(root)

	return root.lhsMonkey.number,root.rhsMonkey.number
}

func reset(tribe map[string]*Monkey){
	for _,monkey := range tribe {
		if !monkey.yelling {
			monkey.resolved = false
		}
	}
}

func resolveMonkey(current *Monkey)int {
	if current.yelling || current.resolved {
		return current.number
	}

	var resolvedNumber int
	switch current.op {
		case "+": resolvedNumber = resolveMonkey(current.lhsMonkey) + resolveMonkey(current.rhsMonkey) 
		case "-": resolvedNumber = resolveMonkey(current.lhsMonkey) - resolveMonkey(current.rhsMonkey) 
		case "*": resolvedNumber = resolveMonkey(current.lhsMonkey) * resolveMonkey(current.rhsMonkey) 
		case "/": resolvedNumber = resolveMonkey(current.lhsMonkey) / resolveMonkey(current.rhsMonkey) 
		default: panic ("Parsing went wrong")
	}

	current.resolved = true
	current.number = resolvedNumber	

	return resolvedNumber
}

func setupProblem(fileName string) map[string]*Monkey {
	//parse file
	operationRe := regexp.MustCompile(`(?P<name>\w{4}): (?P<lhs>\w{4}) (?P<op>\+|-|\*|/) (?P<rhs>\w{4})`)
	yellingRe := regexp.MustCompile(`(?P<name>\w{4}): (?P<number>\d+)`)

	readFile, err := os.Open(fileName)
	check(err)
	defer readFile.Close()

    fileScanner := bufio.NewScanner(readFile)
    fileScanner.Split(bufio.ScanLines)

	tribe := map[string]*Monkey{}
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if operationRe.MatchString(line){
			s := operationRe.FindAllStringSubmatch(fileScanner.Text(), -1)[0][1:]
			name := s[0]
			lhs  := s[1]
			op   := s[2]
			rhs  := s[3]
			tribe[name]=&Monkey{name:name,yelling:false,lhsName:lhs,rhsName:rhs,op:op}
		}

		if yellingRe.MatchString(line){
			s := yellingRe.FindAllStringSubmatch(fileScanner.Text(), -1)[0][1:]
			name := s[0]
			number,_  := strconv.Atoi(s[1])
			tribe[name]=&Monkey{name:name,yelling:true,number:number}
		}
	}

	for _,monkey := range tribe {
		if !monkey.yelling {
			monkey.lhsMonkey = tribe[monkey.lhsName]
			monkey.rhsMonkey = tribe[monkey.rhsName]
		}
	}

	return tribe
}
