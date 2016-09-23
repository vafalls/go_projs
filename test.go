package main

import (
    "fmt"
    //"sort"
    "sort"
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
    list := [5]int{4,3,2,1,0}
    asd := []int{}
    fmt.Println(list)
    asd = getSubList(list)
    fmt.Println(asd)
    sort.Ints(asd)
    fmt.Println(asd)
    fmt.Println(list)

}
func getSubList(list [5]int) []int{
    fmt.Println(list)
    return list[1:3]
}
