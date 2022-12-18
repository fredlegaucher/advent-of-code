package main

import (
	"testing"
)



func TestSolveProblem(t *testing.T){
	results := solveProblem("input_test.txt")	
    if results != 3068 {
        t.Fatalf("Expected 3068, not %v",results)
    }
}












