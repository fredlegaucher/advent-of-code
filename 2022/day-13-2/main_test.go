package main

import (
	"testing"
)

func TestSolveProblem(t *testing.T){
	results := solveProblem("input_test.txt")	
    if results != 140 {
        t.Fatalf("Expected 140, not %v",results)
    }
}