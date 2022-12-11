package main

import (
	"testing"
)

func setupTestMonkeys() []*Monkey {
	monkey0 := &Monkey{ []int{79, 98},"multiply",19,23,2,3,0}
	monkey1 := &Monkey{ []int{54, 65, 75, 74},"add",6,19,2,0,0}
	monkey2 := &Monkey{ []int{79, 60, 97},"square",0,13,1,3,0}
	monkey3 := &Monkey{ []int{74},"add",3,17,0,1,0}
	
	return []*Monkey{monkey0,monkey1,monkey2,monkey3}
}



func TestSolveProblem(t *testing.T){
    
	results := solveProblem(setupTestMonkeys())	

    if results != 10605 {
        t.Fatalf("Expected 10605, not %v",results)
    }




}










