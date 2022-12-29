package main

import (
	"testing"
)

func TestSolveExample(t *testing.T){
	result := solveProblem("input_test.txt")	
    
    if result != 54 {
        t.Fatalf("Expected 54 not %v",result)
    }
}

func TestSolveProblem(t *testing.T){
	result := solveProblem("input.txt")	
    
    if result != 820 {
        t.Fatalf("Expected 820 not %v",result)
    }
}


