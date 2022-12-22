package main

import (
	"testing"
)

func TestSolveExample(t *testing.T){
	a,b := solveProblem("input_test.txt",301)	
    
    if a != b {
        t.Fatalf("Expected 301 to work")
    }
}

func TestSolveProblem(t *testing.T){
	a,b := solveProblem("input.txt",3343167719435)	
    
    if a != b {
        t.Fatalf("Expected 3343167719435 to work")
    }
}


