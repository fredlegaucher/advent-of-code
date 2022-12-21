package main

import (
	"testing"
)

func TestSolveBothExample1(t *testing.T){
	results := solveProblem("input_test_1.txt")	
    if results != 9 {
        t.Fatalf("Expected 9, not %v",results)
    }
}

func TestSolveBothExample2(t *testing.T){
	results := solveProblem("input_test_2.txt")	
    if results != 24 {
        t.Fatalf("Expected 24, not %v",results)
    }
}

func TestSolveFullFile(t *testing.T){
	results := solveProblem("input.txt")	
    if results != 1675 {
        t.Fatalf("Expected 1675, not %v",results)
    }
}













