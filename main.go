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

    for i := 0; i < len(myArgs)-1; i++ {
        add := myArgs[i]
        for n := 0; n < len(myArgs[i]); n++ {
            digit := string(add[n])
            index, _ := strconv.ParseInt(digit, 0, 64)
            toPrint += numWords[index]
        }
        toPrint += ","
    }

    add := myArgs[len(myArgs)-1]
    for n := 0; n < len(myArgs[len(myArgs)-1]); n++ {
        digit := string(add[n])
        index, _ := strconv.ParseInt(digit, 0, 64)
        toPrint += numWords[index]
    }

    fmt.Println(toPrint)
}