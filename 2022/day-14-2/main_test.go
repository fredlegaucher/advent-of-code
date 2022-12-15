package main

import (
	"testing"
)



func TestSolveProblem(t *testing.T){
	results := solveProblem("input_test.txt")	
    if results != 93 {
        t.Fatalf("Expected 93, not %v",results)
    }
}












