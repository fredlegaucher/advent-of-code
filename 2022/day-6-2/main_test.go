package main

import (
	"fmt"
	"testing"
)

func TestSolveProblem(t *testing.T){
	results := solveProblem("input_test.txt")	
	
    if results[0] != "19" {
        t.Fatalf(fmt.Sprintf("Expected a score of 19 rather than %s ", results[0]))
    }

	if results[1] != "23" {
        t.Fatalf(fmt.Sprintf("Expected a score of 23 rather than %s ", results[1]))
    }

	if results[2] != "23" {
        t.Fatalf(fmt.Sprintf("Expected a score of 23 rather than %s ", results[2]))
    }

	if results[3] != "29" {
        t.Fatalf(fmt.Sprintf("Expected a score of 29 rather than %s ", results[3]))
    }

	if results[4] != "26" {
        t.Fatalf(fmt.Sprintf("Expected a score of 26 rather than %s ", results[3]))
    }
}







