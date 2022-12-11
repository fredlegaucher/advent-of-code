package main

import (
	"fmt"
	"fredlegaucher/day-11-2/pfs"
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
	items []uint64
	operation string
	operand uint64
	divideByForTest uint64
	throwToMonkeyIfTrue uint64
	throwToMonkeyIfFalse uint64
	inspectionTimes uint64
}

type PrimeMonkey struct {
	itemsAsPrime []map[uint64]uint64
	operation string
	operand uint64
	operandAsPrime map[uint64]uint64
	divideByForTestAsPrime map[uint64]uint64
	throwToMonkeyIfTrue uint64
	throwToMonkeyIfFalse uint64
	inspectionTimes uint64
}

func (m *PrimeMonkey) inspectItem(family []*PrimeMonkey){
	m.inspectionTimes++

	var itemBefore map[uint64]uint64
	itemBefore,m.itemsAsPrime = m.itemsAsPrime[0],m.itemsAsPrime[1:]

	var itemAfter map[uint64]uint64
	switch m.operation{
		case "multiply": itemAfter = pfs.PfsMultiplication(itemBefore,m.operandAsPrime)
		case "add": itemAfter = pfs.AddInttoPFS(itemBefore,m.operand)
		case "square": itemAfter = pfs.PfsMultiplication(itemBefore,itemBefore)
	default: panic("Unsupported operation")
	}

	if pfs.CanPFSBeDivided(itemAfter,m.divideByForTestAsPrime) {
		family[m.throwToMonkeyIfTrue].itemsAsPrime = append(family[m.throwToMonkeyIfTrue].itemsAsPrime, itemAfter)
	} else {
		family[m.throwToMonkeyIfFalse].itemsAsPrime = append(family[m.throwToMonkeyIfFalse].itemsAsPrime, itemAfter)
	}
}

func setupMonkeys() []*PrimeMonkey {
	monkey0 := convertToPrimeMonkey(&Monkey{ []uint64{50, 70, 54, 83, 52, 78},"multiply",3,11,2,7,0})
	monkey1 := convertToPrimeMonkey(&Monkey{ []uint64{71, 52, 58, 60, 71},"square",0,7,0,2,0})
	monkey2 := convertToPrimeMonkey(&Monkey{ []uint64{66, 56, 56, 94, 60, 86, 73},"add",1,3,7,5,0})
	monkey3 := convertToPrimeMonkey(&Monkey{ []uint64{83, 99},"add",8,5,6,4,0})
	monkey4 := convertToPrimeMonkey(&Monkey{ []uint64{98, 98, 79},"add",3,17,1,0,0})
	monkey5 := convertToPrimeMonkey(&Monkey{ []uint64{76},"add",4,13,6,3,0})
	monkey6 := convertToPrimeMonkey(&Monkey{ []uint64{52, 51, 84, 54},"multiply",17,19,4,1,0})
	monkey7 := convertToPrimeMonkey(&Monkey{ []uint64{82, 86, 91, 79, 94, 92, 59, 94},"add",7,2,5,3,0})

	fmt.Println("Monkeys setup")
	return []*PrimeMonkey{monkey0,monkey1,monkey2,monkey3,monkey4,monkey5,monkey6,monkey7}
}

func convertToPrimeMonkey(m *Monkey) *PrimeMonkey {
	primedItem := make([]map[uint64]uint64,len(m.items))

	for i,k := range m.items {
		primedItem[i] = pfs.PrimeFactors(k)
	}

	return &PrimeMonkey{
		itemsAsPrime: primedItem,
		operation: m.operation,
		operand: m.operand,
		operandAsPrime: pfs.PrimeFactors(m.operand),
		divideByForTestAsPrime: pfs.PrimeFactors(m.divideByForTest),
		throwToMonkeyIfTrue: m.throwToMonkeyIfTrue,
		throwToMonkeyIfFalse: m.throwToMonkeyIfFalse,
		inspectionTimes:0,
	}
}


func solveProblem(monkeys []*PrimeMonkey) uint64 {
	

	//play the game
	for round := 0 ; round < 20 ; round ++{
		fmt.Println("Round: ",round+1)	
		roundOfGame(monkeys)
		
		//if round+1 == 1 || round+1 == 20 || round+1 == 1000 || round+1 == 2000 || round+1 == 10000 {
		
			for i,m := range monkeys {
				fmt.Println("Monkey ",i,":", m.inspectionTimes)
			}
			fmt.Println()
		//}
	}

	//find the monkeys with the highest inspection time
	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].inspectionTimes > monkeys[j].inspectionTimes
	  })

	  
	return monkeys[0].inspectionTimes * monkeys[1].inspectionTimes
	
}



func roundOfGame(monkeys []*PrimeMonkey){
	for _,monkey := range monkeys {
		monkey.turn(monkeys)
	}
}

func (m *PrimeMonkey) turn(family []*PrimeMonkey){
	for len(m.itemsAsPrime) > 0 {
		m.inspectItem(family)
	}
}





