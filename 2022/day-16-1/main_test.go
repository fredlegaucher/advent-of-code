package main

import (
	"testing"
)



func TestSolveProblem(t *testing.T){
	results := solveProblem("input_test.txt")	
    if results != 1651 {
        t.Fatalf("Expected 1651, not %v",results)
    }
}












