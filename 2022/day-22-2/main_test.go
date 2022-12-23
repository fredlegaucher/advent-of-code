package main

import (
	"testing"
)

func TestSolveProblem(t *testing.T){
	a,b,c := solveProblem("input.txt")	
    
    if a != 197 || b != 30 || c!= 2 {
        t.Fatalf("Expected 197,30,2 not %v,%v,%v",a,b,c)
    }
}