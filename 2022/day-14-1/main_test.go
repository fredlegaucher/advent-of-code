package main

import (
	"testing"
)



func TestSolveProblem(t *testing.T){
	results := solveProblem("input_test.txt")	
    if results != 24 {
        t.Fatalf("Expected 24, not %v",results)
    }
}












