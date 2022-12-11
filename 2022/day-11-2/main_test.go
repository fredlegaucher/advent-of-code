package main

import (
	"testing"
)

func setupTestMonkeys() []*SuperMonkey {
	monkey0 := &Monkey{ []int{79, 98},"multiply",19,23,2,3,0}
	monkey1 := &Monkey{ []int{54, 65, 75, 74},"add",6,19,2,0,0}
	monkey2 := &Monkey{ []int{79, 60, 97},"square",0,13,1,3,0}
	monkey3 := &Monkey{ []int{74},"add",3,17,0,1,0}
	family := []*Monkey{monkey0,monkey1,monkey2,monkey3}

	return []*SuperMonkey{
		monkey0.convertToSuperMonkey(family),
		monkey1.convertToSuperMonkey(family),
		monkey2.convertToSuperMonkey(family),
		monkey3.convertToSuperMonkey(family),
		}
}

func TestSolveProblemSlow(t *testing.T){
	results := solveProblem(setupTestMonkeys(),10000)	
    if results != 2713310158 {
        t.Fatalf("Expected 2713310158, not %v",results)
    }
}

func TestSolveProblemQuick(t *testing.T){
	results := solveProblem(setupTestMonkeys(),20)	
    if results != 99 * 103 {
        t.Fatalf("Expected %v, not %v",99*103,results)
    }
}










