package main

import (
	"testing"
)

func TestSolveSnafuToBase10(t *testing.T){
	results := convertFileFromSnafu("input_test.txt")	
    expected_values := []int{1747,906,198,11,201,31,1257,32,353,107,7,3,37}
    
    for i,value := range results {
        if value != expected_values[i]{
            t.Fatalf("Expected %v not %v",expected_values[i],value)
        }
    }
}

func TestSolveBase10ToSnafu(t *testing.T){
    input_values := []int{1747,906,198,11,201,31,1257,32,353,107,7,3,37}
    snafu_values := []string{"1=-0-2",
    "12111",
    "2=0=",
    "21",
    "2=01",
    "111",
    "20012",
    "112",
    "1=-1=",
    "1-12",
    "12",
    "1=",
    "122"}
    
    for i,value := range input_values {
        result := convertFromBase10ToSnafu(value)
        if result != snafu_values[i]{
            t.Fatalf("Expected %v not %v",snafu_values[i],result)
        }
    }
}

func Test1To12ToSnafu(t *testing.T){

    
    for i:=0 ; i < 1000 ; i++ {
        result := convertFromBase10ToSnafu(i)
        if convertToBase10(result) != i{
            t.Fatalf(" %v is not the right snafu for %v",result,i)
        }
    }
}

func TestSolveProblem(t *testing.T){
    intermediaryResult := 0 
	for _,value := range convertFileFromSnafu("input.txt"){
        intermediaryResult += value
    }
    result := convertFromBase10ToSnafu(intermediaryResult)
    if result != "Fred"{
        t.Fatalf("Expected %v not %v","Fred",result)
    }
    
}


