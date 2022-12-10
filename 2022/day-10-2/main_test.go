package main

import (
	"testing"
)

func TestSolveProblem(t *testing.T){
	results := solveProblem("input_test.txt")	


	
    if results[0] != "##..##..##..##..##..##..##..##..##..##.." {
        t.Fatalf("Expected a different output")
    }

	if results[1] != "###...###...###...###...###...###...###." {
        t.Fatalf("Expected a different output")
    }

	if results[2] != "####....####....####....####....####...." {
        t.Fatalf("Expected a different output")
    }

	if results[3] != "#####.....#####.....#####.....#####....." {
        t.Fatalf("Expected a different output")
    }

	if results[4] != "######......######......######......####" {
        t.Fatalf("Expected a different output")
    }

	if results[5] != "#######.......#######.......#######....." {
        t.Fatalf("Expected a different output")
    }


}










