package main

import (
	"testing"
)

func TestSolveProblemExample(t *testing.T){
	results := solveProblem("input_test.txt")	
    if results != 58 {
        t.Fatalf("Expected 58, not %v",results)
    }
}

