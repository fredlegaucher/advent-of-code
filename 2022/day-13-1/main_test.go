package main

import (
	"testing"
)



func TestSolveProblem(t *testing.T){
	results := solveProblem("input_test.txt")	
    if results != 13 {
        t.Fatalf("Expected 13, not %v",results)
    }
}

func TestSolveProblemPairByPair(t *testing.T){

    if isPairCorrectlyOrdered("[7,7,7,7]","[7,7,7]")	 {
        t.Fatalf("Expected false")
    }

    if !isPairCorrectlyOrdered("[]","[3]")	 {
        t.Fatalf("Expected true")
    }


    if !isPairCorrectlyOrdered("[1,1,3,1,1]","[1,1,5,1,1]") {
        t.Fatalf("Expected true")
    }


    if !isPairCorrectlyOrdered("[[1],[2,3,4]]","[[1],4]")	 {
        t.Fatalf("Expected true")
    }

  

    if isPairCorrectlyOrdered("[9]","[[8,7,6]]")	 {
        t.Fatalf("Expected false")
    }

    if !isPairCorrectlyOrdered("[[4,4],4,4]","[[4,4],4,4,4]")	 {
        t.Fatalf("Expected true")
    }





    if isPairCorrectlyOrdered("[[[]]]","[[]]")	 {
        t.Fatalf("Expected false")
    }

    if isPairCorrectlyOrdered("[1,[2,[3,[4,[5,6,7]]]],8,9]","[1,[2,[3,[4,[5,6,0]]]],8,9]") {
        t.Fatalf("Expected false")
    }

}
