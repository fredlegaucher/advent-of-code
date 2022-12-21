package main

import (
	"testing"
)

func TestSolveExample(t *testing.T){
	a,b,c := solveProblem("input_test.txt")	
    
    if a != 4 || b != -3 || c!= 2 {
        t.Fatalf("Expected 4,-3,2 not %v,%v,%v",a,b,c)
    }
}

func TestSolveProblem(t *testing.T){
	a,b,c := solveProblem("input.txt")	
    
    if a != 8666 || b != 4363 || c!= -8805 {
        t.Fatalf("Expected 8666,4363,-8805 not %v,%v,%v",a,b,c)
    }
}
