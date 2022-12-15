package main

import (
	"testing"
)



func TestSolveProblem(t *testing.T){
	results := solveProblem("input_test.txt",20)	
    if results != 56000011 {
        t.Fatalf("Expected 56000011, not %v",results)
    }
}












