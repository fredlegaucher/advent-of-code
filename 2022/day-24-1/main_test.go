package main

import (
	"testing"
)

func TestSolveExample(t *testing.T){
	result := solveProblem("input_test.txt")	
    
    if result != 18 {
        t.Fatalf("Expected 18 not %v",result)
    }
}

func TestSolveProblem(t *testing.T){
	result := solveProblem("input.txt")	
    
    if result != 286 {
        t.Fatalf("Expected 286 not %v",result)
    }
}


