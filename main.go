// this solution has the same process as the javascript solution but assigns variables where needed
package main

import (
    "fmt"
    "os"
    "strconv"
)

func main() {
    myArgs := os.Args[1:]
    numWords := [10]string{"Zero", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine"}
    toPrint := ""

    // loop through array of arguments appending the correct string values to the string we will print
    for i := 0; i < len(myArgs)-1; i++ {
        add := myArgs[i]
        for n := 0; n < len(myArgs[i]); n++ {
            digit := string(add[n])
            index, _ := strconv.ParseInt(digit, 0, 64)
            toPrint += numWords[index]
        }
        toPrint += ","
    }

    // similar to the javascript solution this part handles not ending the string we want to print with a comma
    // another way to deal with this is just use one for loop and delete the last comma after getting the string toPrint
    add := myArgs[len(myArgs)-1]
    for n := 0; n < len(myArgs[len(myArgs)-1]); n++ {
        digit := string(add[n])
        index, _ := strconv.ParseInt(digit, 0, 64)
        toPrint += numWords[index]
    }

    fmt.Println(toPrint)
}