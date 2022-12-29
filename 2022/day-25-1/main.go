package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
) 

func check(e error) {
        if e != nil {
                panic(e)
        }
}

func main() {
        result := 0 
	for _,value := range convertFileFromSnafu("input.txt"){
                result += value
        }
        fmt.Println("Result in base 10: ", result)
        fmt.Println("Result in base SNAFU: ", convertFromBase10ToSnafu(result))

}

func convertFromBase10ToSnafu(base10 int) string{
       
        //find length
         lengthInSnaffu,maxValue,maxValuePreviousRound := 0,0,0

        for ;   ; {
                maxValue = maxValuePreviousRound + 2*pow(5,lengthInSnaffu)
                if maxValue >= base10 {
                        break
                }
                maxValuePreviousRound = maxValue
                lengthInSnaffu++
        }
        lengthInSnaffu++

                //bigger than 2 * 5^(lengthInSnaffu-1) -> 2 is the first character. recurse of remainder
        //or less than it but strictly more than 1 * 5^(lengthInSnaffu-1) + maxValueFromPreviousRound ->2 is still the first value, recurse of -minder and flip the result
        //or more than 1 * 5^(lengthInSnaffu-1) -> 1 is the first character. recurse of remainder
        //or less than 5^(lengthInSnaffu-1) but more than -> 1 is the first character. recurse of -remainder, flip the result
        if lengthInSnaffu == 1 && (base10 > 2 || base10 < -2){
                panic("that's not right")
        }
        if lengthInSnaffu == 1 {
                switch base10 {
                case 2,1,0: return fmt.Sprintf("%v",base10)
                case -1: return "-"
                case -2: return "="
                }
        }



        if base10 >= 2 * pow(5,lengthInSnaffu-1){
                firstCharacter := "2"
                remainder := convertFromBase10ToSnafu(base10 - 2*pow(5,lengthInSnaffu-1))
                if len(remainder) < lengthInSnaffu - 1 {
                        for i:= 0 ; i < lengthInSnaffu - 1 - len(remainder); i++{
                                firstCharacter += "0"
                        }
                }
                return firstCharacter + remainder
        } else if base10 < 2 * pow(5,lengthInSnaffu-1) && base10 > pow(5,lengthInSnaffu-1) + maxValuePreviousRound  {
                firstCharacter := "2"
                remainder := flip(convertFromBase10ToSnafu(2*pow(5,lengthInSnaffu-1) - base10))
                if len(remainder) < lengthInSnaffu - 1 {
                        for i:= 0 ; i < lengthInSnaffu - 1 - len(remainder); i++{
                                firstCharacter += "0"
                        }
                }
                return firstCharacter + remainder
        } else if base10 >= pow(5,lengthInSnaffu-1) {
                firstCharacter := "1"
                remainder := convertFromBase10ToSnafu(base10 - pow(5,lengthInSnaffu-1))
                if len(remainder) < lengthInSnaffu - 1 {
                        for i:= 0 ; i < lengthInSnaffu - 1 - len(remainder); i++{
                                firstCharacter += "0"
                        }
                }
                return firstCharacter + remainder
        } else if base10 < pow(5,lengthInSnaffu-1) && base10 > maxValuePreviousRound {
                firstCharacter := "1"
                remainder := flip(convertFromBase10ToSnafu(pow(5,lengthInSnaffu-1) - base10))
                if len(remainder) < lengthInSnaffu - 1 {
                        for i:= 0 ; i < lengthInSnaffu - 1 - len(remainder); i++{
                                firstCharacter += "0"
                        }
                }
                return firstCharacter + remainder
        } else {
                panic("Should not be here")
        }

}

func flip(input string) string {
        result := ""
        for _,char := range input {
                switch char {
                        case '2': result += "="
                        case '1': result += "-"
                        case '0': result += "0"
                        case '-': result += "1"
                        case '=': result += "2"
                }
        }
        return result
}


var values = []string{"2","1","0","-","="}

func convertFileFromSnafu(fileName string) []int {
        readFile, err := os.Open(fileName)
        check(err)
        defer readFile.Close()

        fileScanner := bufio.NewScanner(readFile)
        fileScanner.Split(bufio.ScanLines)

        results := []int{} 
        for fileScanner.Scan() {
                results = append(results, convertToBase10(fileScanner.Text()))
        }
        return results
}

func convertToBase10(input string) int {
        width := len(input)
        result := 0
        
        for i := 0 ; i < width ; i++ {
                char := string(input[width - 1 - i])
                if char == "=" {
                        result += (-2) * pow(5,i)
                        continue
                }
                if char == "-" {
                        result += (-1) * pow(5,i)
                        continue
                }
                actualInt,_ := strconv.Atoi(char)
                result += actualInt * pow(5,i)
        }
        return result
}

func pow(input int, power int)int {
        return _pow(1,input,power)
}

func _pow(accumulateResult int, input int, power int)int{
        if power == 0 {
                return accumulateResult
        }
        return _pow(input*accumulateResult,input,power-1)
}