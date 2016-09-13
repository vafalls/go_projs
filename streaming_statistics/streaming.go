package main

import (
    "fmt"
    "sort"
    "os"
)

type logEntry struct {
    startTime uint64
    endTime uint64
    duration uint64
    bitRate uint16
}

type query struct {
    startTime uint64
    endTime uint64
}

type LogEntries []logEntry
var sortByEndtime bool

func (slice LogEntries) Len() int {
    return len(slice)
}
func (slice LogEntries) Less(i, j int) bool {
    if sortByEndtime {
        return slice[i].endTime < slice[j].endTime;
    } else {
        return slice[i].startTime < slice[j].startTime
    }
}
func (slice LogEntries) Swap(i, j int) {
    slice[i], slice[j] = slice[j], slice[i]
}

func scanData(logEntries LogEntries, sliceOfQueries []query) (LogEntries, []query){
    var nrOfLogEntries,nrOfQueries uint32

    fmt.Scanf("%d", &nrOfLogEntries)
    for i:=uint32(0); i<nrOfLogEntries; i++{
        tmp_entry := logEntry{}
        fmt.Scanf("%d %d %d", &tmp_entry.endTime, &tmp_entry.duration, &tmp_entry.bitRate)
        tmp_entry.startTime = tmp_entry.endTime - tmp_entry.duration
        logEntries = append(logEntries, tmp_entry)
    }

    fmt.Scanf("%d", &nrOfQueries)
    for i:=uint32(0); i<nrOfQueries; i++{
        tmp_query := query{}
        fmt.Scanf("%d %d", &tmp_query.startTime, &tmp_query.endTime)
        sliceOfQueries = append(sliceOfQueries, tmp_query)
    }
    return logEntries, sliceOfQueries
}

func calc(query query, logEntries LogEntries) (float64) {
    var intervalTime float64
    var startTime, stopTime uint64

    for i:=0; i<len(logEntries); i++{
        
        if logEntries[i].startTime >= query.startTime {
            startTime = logEntries[i].startTime
        } else {
            startTime = query.startTime
        }
        if logEntries[i].endTime <= query.endTime {
            stopTime = logEntries[i].endTime
        } else {
            stopTime = query.endTime
        }

        if startTime < stopTime {
            intervalTime += float64(float64(logEntries[i].bitRate)/float64(1000) * float64(stopTime-startTime))
        }
    }
    return intervalTime
}

func removeEarlierPart(logEntries LogEntries, queryStartTime uint64)(LogEntries) {
    low := 0
    high := len(logEntries)-1
    mid := 0

    if logEntries[low].endTime > queryStartTime {
        return logEntries
    } else if logEntries[high].endTime <= queryStartTime {
        fmt.Println("0.000")
        os.Exit(0)
    }

    for high - low > 1{
        mid = low + (high-low)/2

        //fmt.Println(low,mid,high)

        if high-low == 1 || high-low==2 {
            if logEntries[mid].endTime <= queryStartTime && logEntries[mid+1].endTime > queryStartTime{
                return logEntries[mid+1:]
            }
            if logEntries[mid+1].endTime <= queryStartTime && mid+2 < len(logEntries)-1 && logEntries[mid+2].endTime > queryStartTime{

                return logEntries[mid+2:]
            }
        }

        if queryStartTime < logEntries[mid].endTime {
            high = mid
        } else if queryStartTime > logEntries[mid].endTime {
            low = mid
        } else {
            high = high-1
        }
    }
    return logEntries
}

func removeLaterPart(logEntries LogEntries, queryEndTime uint64)(LogEntries) {
    low := 0
    high := len(logEntries)-1
    mid := 0

    if logEntries[high].startTime < queryEndTime {
        return logEntries
    } else if logEntries[low].startTime >= queryEndTime {
        fmt.Println("0.000")
        os.Exit(0)
    }

    for high - low > 1 {
        mid = low + (high-low)/2

        //fmt.Println(low,mid,high)

        if high-low == 1 || high-low==2 {
            if logEntries[mid].startTime >= queryEndTime && logEntries[mid+1].startTime > queryEndTime {
                return logEntries[0:mid-1]
            }
            if logEntries[mid + 1].startTime >= queryEndTime && mid+2 < len(logEntries)-1 && logEntries[mid+2].startTime > queryEndTime {
                return logEntries[0:mid]
            }
        }

        if queryEndTime < logEntries[mid].startTime {
            high = mid
        } else if queryEndTime > logEntries[mid].startTime {
            low = mid
        } else {
            high = high-1
        }
    }
    return logEntries
}

func getValuesToBeCalculated(logEntriesEarlier LogEntries, logEntriesLater LogEntries) (LogEntries) {
    valuesToBeCalculated := LogEntries{}
    for _,earlierElement := range logEntriesEarlier {
        for i:=0; i<len(logEntriesLater); i++ {
            //fmt.Println("index is",i)
            if earlierElement.startTime == logEntriesLater[i].startTime &&
                    earlierElement.bitRate == logEntriesLater[i].bitRate &&
                    earlierElement.duration == logEntriesLater[i].duration &&
                    earlierElement.endTime == logEntriesLater[i].endTime {
                valuesToBeCalculated = append(valuesToBeCalculated, earlierElement)
                logEntriesLater = append(logEntriesLater[:i], logEntriesLater[i+1:]...)
                //fmt.Println("breaking")
                break
            }

        }
    }
    //fmt.Println("this returns")
    //fmt.Println(valuesToBeCalculated)
    return valuesToBeCalculated
}

func main() {

    logEntriesSortedByStart := LogEntries{}
    var sliceOfQueries = make([]query, 0)

    logEntriesSortedByStart, sliceOfQueries = scanData(logEntriesSortedByStart, sliceOfQueries)

    sortByEndtime = false
    sort.Sort(logEntriesSortedByStart)

    logEntriesSortedByEnd := make(LogEntries, len(logEntriesSortedByStart))
    copy(logEntriesSortedByEnd, logEntriesSortedByStart)
    sortByEndtime = true
    sort.Sort(logEntriesSortedByEnd)

    logEntriesLater := LogEntries{}
    logEntriesEarlier := LogEntries{}
    valuesToBeCalculated := LogEntries{}


    //fmt.Println("logEntriesSortedByStart")
    //fmt.Println(logEntriesSortedByStart)
    //fmt.Println("logEntriesSortedByEnd")
    //fmt.Println(logEntriesSortedByEnd)

    for _,element := range sliceOfQueries {

        logEntriesEarlier = removeLaterPart(logEntriesSortedByStart, element.endTime)
        //fmt.Println("logEntriesEarlier")
        //fmt.Println(logEntriesEarlier)

        logEntriesLater = removeEarlierPart(logEntriesSortedByEnd, element.startTime)
        //fmt.Println("logEntriesLater")
        //fmt.Println(logEntriesLater)

        valuesToBeCalculated = getValuesToBeCalculated(logEntriesEarlier, logEntriesLater)
        //fmt.Println("valuesToBeCalculated")
        //fmt.Println(valuesToBeCalculated)
        fmt.Printf("%.3f\n",calc(element, valuesToBeCalculated))
    }
}