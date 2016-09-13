package main

import (
    "fmt"
    //"sort"
)

type query struct {
    firstNumber int
    secondNumber int
}

//sorts by startTime
type LogEntries []query
var sortFlag bool

func (slice LogEntries) Len() int {
    return len(slice)
}
func (slice LogEntries) Less(i, j int) bool {
    if sortFlag {
        return slice[i].secondNumber < slice[j].secondNumber;
    } else {
        return slice[i].firstNumber < slice[j].firstNumber;
    }
}
func (slice LogEntries) Swap(i, j int) {
    slice[i], slice[j] = slice[j], slice[i]
}

func test(mystring *string){
    *mystring = "is now changed"
}

func main() {
    mystring := "this is my string"
    fmt.Println(mystring)
    test(&mystring)
    fmt.Println(mystring)
}
