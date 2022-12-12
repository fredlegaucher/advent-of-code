package main

import (
	"testing"
)



func TestSolveProblem(t *testing.T){
	results := solveProblem("input_test.txt")	
    if results != 29 {
        t.Fatalf("Expected 29, not %v",results)
    }
}












