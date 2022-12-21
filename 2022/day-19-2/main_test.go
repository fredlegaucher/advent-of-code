package main

import (
	"testing"
)

func TestSolveBothExample1(t *testing.T){
	results := solveProblem("input_test_1.txt")	
    if results != 56 {
        t.Fatalf("Expected 56, not %v",results)
    }
}

func TestSolveBothExample2(t *testing.T){
	results := solveProblem("input_test_2.txt")	
    if results != 62 {
        t.Fatalf("Expected 62, not %v",results)
    }
}














