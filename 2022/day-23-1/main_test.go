package main

import (
	"testing"
)

func TestSolveExample(t *testing.T){
	result := solveProblem("input_test.txt")	
    
    if result != 110 {
        t.Fatalf("Expected 110 not %v",result)
    }
}

func TestSolveProblem(t *testing.T){
	result := solveProblem("input.txt")	
    
    if result != 3970 {
        t.Fatalf("Expected 3970 not %v",result)
    }
}


