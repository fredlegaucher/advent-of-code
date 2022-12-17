package main

import (
	"testing"
)



func TestSolveProblem(t *testing.T){
	results := solveProblem("input_test.txt")	
    if results != 1707 {
        t.Fatalf("Expected 1707, not %v",results)
    }
}












