package main

import (
	"fmt"
	"testing"
)

func TestSolveProblem(t *testing.T){
	results := solveProblem("input_test.txt")	
	
    if results[0] != "5" {
        t.Fatalf(fmt.Sprintf("Expected a score of 5 rather than %s ", results[0]))
    }

	if results[1] != "6" {
        t.Fatalf(fmt.Sprintf("Expected a score of 6 rather than %s ", results[1]))
    }

	if results[2] != "10" {
        t.Fatalf(fmt.Sprintf("Expected a score of 10 rather than %s ", results[2]))
    }

	if results[3] != "11" {
        t.Fatalf(fmt.Sprintf("Expected a score of 11 rather than %s ", results[3]))
    }

	if results[4] != "7" {
        t.Fatalf(fmt.Sprintf("Expected a score of 7 rather than %s ", results[3]))
    }
}







