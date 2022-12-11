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
	fmt.Println(solveProblem(monkeys,10000))
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

type ModuloedItem struct{
	item int
	modulos map[int]int
}


type SuperMonkey struct {
	moduloedItems []*ModuloedItem
	operation string
	operand int
	divideByForTest int
	throwToMonkeyIfTrue int
	throwToMonkeyIfFalse int
	inspectionTimes int
}

func (m Monkey) convertToSuperMonkey(family []*Monkey) *SuperMonkey{	
	superItems := make([]*ModuloedItem,0)
	for _,item := range m.items {
		modulosForItem := make(map[int]int)
		for _,monkey := range family {
			modulosForItem[int(monkey.divideByForTest)] = item % monkey.divideByForTest
		}

		superItems = append(superItems, &ModuloedItem{item: item, modulos: modulosForItem})
	}

	return &SuperMonkey{
		moduloedItems: superItems,
		operation: m.operation,
		operand: m.operand,
		divideByForTest: m.divideByForTest,
		throwToMonkeyIfTrue: m.throwToMonkeyIfTrue,
		throwToMonkeyIfFalse: m.throwToMonkeyIfFalse,
		inspectionTimes: 0,
	}
}

func setupMonkeys() []*SuperMonkey {
	monkey0 := &Monkey{ []int{50, 70, 54, 83, 52, 78},"multiply",3,11,2,7,0}
	monkey1 := &Monkey{ []int{71, 52, 58, 60, 71},"square",0,7,0,2,0}
	monkey2 := &Monkey{ []int{66, 56, 56, 94, 60, 86, 73},"add",1,3,7,5,0}
	monkey3 := &Monkey{ []int{83, 99},"add",8,5,6,4,0}
	monkey4 := &Monkey{ []int{98, 98, 79},"add",3,17,1,0,0}
	monkey5 := &Monkey{ []int{76},"add",4,13,6,3,0}
	monkey6 := &Monkey{ []int{52, 51, 84, 54},"multiply",17,19,4,1,0}
	monkey7 := &Monkey{ []int{82, 86, 91, 79, 94, 92, 59, 94},"add",7,2,5,3,0}

	family := []*Monkey{monkey0,monkey1,monkey2,monkey3,monkey4,monkey5,monkey6,monkey7}

	return []*SuperMonkey{
		monkey0.convertToSuperMonkey(family),
		monkey1.convertToSuperMonkey(family),
		monkey2.convertToSuperMonkey(family),
		monkey3.convertToSuperMonkey(family),
		monkey4.convertToSuperMonkey(family),
		monkey5.convertToSuperMonkey(family),
		monkey6.convertToSuperMonkey(family),
		monkey7.convertToSuperMonkey(family)}
}

func solveProblem(monkeys []*SuperMonkey, roundCount int) int {
	//play the game
	for round := 0 ; round < roundCount ; round ++{
		roundOfGame(monkeys)
	}

	//find the monkeys with the highest inspection time
	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].inspectionTimes > monkeys[j].inspectionTimes
	})

	return monkeys[0].inspectionTimes * monkeys[1].inspectionTimes
	
}

func roundOfGame(monkeys []*SuperMonkey){
	for _,monkey := range monkeys {
		monkey.turn(monkeys)
	}
}

func (m *SuperMonkey) turn(family []*SuperMonkey){
	for len(m.moduloedItems) > 0 {
		m.inspectItem(family)
	}
}

func (m *SuperMonkey) inspectItem(family []*SuperMonkey){
	m.inspectionTimes++

	var currentItem *ModuloedItem
	currentItem,m.moduloedItems = m.moduloedItems[0],m.moduloedItems[1:]

	switch m.operation{
		case "multiply": multiplicationByInt(currentItem,m.operand)
		case "add": additionOfInt(currentItem,m.operand)
		case "square": square(currentItem)
		default: panic("Unsupported operation")
	}

	if currentItem.modulos[m.divideByForTest] == 0 {
		family[m.throwToMonkeyIfTrue].moduloedItems = append(family[m.throwToMonkeyIfTrue].moduloedItems, currentItem)
	} else {
		family[m.throwToMonkeyIfFalse].moduloedItems = append(family[m.throwToMonkeyIfFalse].moduloedItems, currentItem)
	}
}

func multiplicationByInt(moduloedItem *ModuloedItem, multiplier int) {
	for k,v := range moduloedItem.modulos {
		moduloedItem.modulos[k] = (v * multiplier) % k 
	}
}

func additionOfInt(moduloedItem *ModuloedItem, valueToBeAdded int) {
	for k,v := range moduloedItem.modulos {
		moduloedItem.modulos[k] = (v + (valueToBeAdded % k)) % k 
	}
}

func square(moduloedItem *ModuloedItem) {
	for k,v := range moduloedItem.modulos {
		moduloedItem.modulos[k] = (v * v) % k 
	}
}





