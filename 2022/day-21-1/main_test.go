package main

import (
	"testing"
)

func TestSolveExample(t *testing.T){
	result := solveProblem("input_test.txt")	
    
    if result != 152 {
        t.Fatalf("Expected 152 not %v",result)
    }
}

func TestSolveProblem(t *testing.T){
	result := solveProblem("input.txt")	
    
    if result != 168502451381566 {
        t.Fatalf("Expected 168502451381566 not %v",result)
    }
}
