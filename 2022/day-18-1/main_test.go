package main

import (
	"testing"
)

func TestSolveProblem(t *testing.T){
	results := solveProblem("input_test.txt")	
    if results != 64 {
        t.Fatalf("Expected 64, not %v",results)
    }
}

func TestSolveProblemShort(t *testing.T){
	results := solveProblem("input_test_0.txt")	
    if results != 10 {
        t.Fatalf("Expected 10, not %v",results)
    }
}












