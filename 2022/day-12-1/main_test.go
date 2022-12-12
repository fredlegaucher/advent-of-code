package main

import (
	"testing"
)



func TestSolveProblem(t *testing.T){
	results := solveProblem("input_test.txt")	
    if results != 31 {
        t.Fatalf("Expected 31, not %v",results)
    }
}












