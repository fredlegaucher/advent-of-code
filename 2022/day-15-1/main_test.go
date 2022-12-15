package main

import (
	"testing"
)



func TestSolveProblem(t *testing.T){
	results := solveProblem("input_test.txt",10)	
    if results != 26 {
        t.Fatalf("Expected 26, not %v",results)
    }
}












