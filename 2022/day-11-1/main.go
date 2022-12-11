package main

import (
	"fmt"
	"sort"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
	monkeys := setupMonkeys()
	fmt.Println(solveProblem(monkeys))
}


type Monkey struct {
	items []int
	operation string
	operand int
	divideByForTest int
	throwToMonkeyIfTrue int
	throwToMonkeyIfFalse int
	inspectionTimes int
}

func (m *Monkey) inspectItem(family []*Monkey){
	m.inspectionTimes++

	var item int
	item,m.items = m.items[0],m.items[1:]

	switch m.operation{
		case "multiply": item *= m.operand
		case "add": item += m.operand
		case "square": item *= item
	default: panic("Unsupported operation")
	}

	//relief
	item /= 3

	if item % m.divideByForTest == 0 {
		family[m.throwToMonkeyIfTrue].items = append(family[m.throwToMonkeyIfTrue].items, item)
	} else {
		family[m.throwToMonkeyIfFalse].items = append(family[m.throwToMonkeyIfFalse].items, item)
	}
}

func setupMonkeys() []*Monkey {
	monkey0 := &Monkey{ []int{50, 70, 54, 83, 52, 78},"multiply",3,11,2,7,0}
	monkey1 := &Monkey{ []int{71, 52, 58, 60, 71},"square",0,7,0,2,0}
	monkey2 := &Monkey{ []int{66, 56, 56, 94, 60, 86, 73},"add",1,3,7,5,0}
	monkey3 := &Monkey{ []int{83, 99},"add",8,5,6,4,0}
	monkey4 := &Monkey{ []int{98, 98, 79},"add",3,17,1,0,0}
	monkey5 := &Monkey{ []int{76},"add",4,13,6,3,0}
	monkey6 := &Monkey{ []int{52, 51, 84, 54},"multiply",17,19,4,1,0}
	monkey7 := &Monkey{ []int{82, 86, 91, 79, 94, 92, 59, 94},"add",7,2,5,3,0}

	return []*Monkey{monkey0,monkey1,monkey2,monkey3,monkey4,monkey5,monkey6,monkey7}
}


func solveProblem(monkeys []*Monkey) int {
	//play the game
	for round := 0 ; round < 20 ; round ++{
		fmt.Println("Round: ",round)
		roundOfGame(monkeys)
		
		for i,m := range monkeys {
			fmt.Println("Monkey ",i,":",m.items," and times ", m.inspectionTimes)
		}
	}

	//find the monkeys with the highest inspection time
	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].inspectionTimes > monkeys[j].inspectionTimes
	  })

	return monkeys[0].inspectionTimes * monkeys[1].inspectionTimes
	
}

func roundOfGame(monkeys []*Monkey){
	for _,monkey := range monkeys {
		monkey.turn(monkeys)
	}
}

func (m *Monkey) turn(family []*Monkey){
	for len(m.items) > 0 {
		m.inspectItem(family)
	}
}





