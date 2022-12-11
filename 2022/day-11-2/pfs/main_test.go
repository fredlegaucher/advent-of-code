package main

import (
	"fmt"
	"testing"
)

func setupTestMonkeys() []*PrimeMonkey {
	fmt.Println("Setting up monkeys")	
	fmt.Printf("%v",^uint64(0))
	monkey0 := convertToPrimeMonkey(&Monkey{ []uint64{79, 98},"multiply",19,23,2,3,0})
	monkey1 := convertToPrimeMonkey(&Monkey{ []uint64{54, 65, 75, 74},"add",6,19,2,0,0})
	monkey2 := convertToPrimeMonkey(&Monkey{ []uint64{79, 60, 97},"square",0,13,1,3,0})
	monkey3 := convertToPrimeMonkey(&Monkey{ []uint64{74},"add",3,17,0,1,0})
	fmt.Println("Monkeys setup")	
	return []*PrimeMonkey{monkey0,monkey1,monkey2,monkey3}
}


func TestSolveProblem(t *testing.T){
	results := solveProblem(setupTestMonkeys())	
    if results != 2713310158 {
        t.Fatalf("Expected 2713310158, not %v",results)
    }
}










