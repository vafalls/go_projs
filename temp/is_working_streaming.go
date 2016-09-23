package main

import (
    "fmt"
    "sort"
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
var sliceOfQueries = make([]query, 0)

func (slice LogEntries) Len() int {
    return len(slice)
}

func (slice LogEntries) Less(i, j int) bool {
    return slice[i].startTime < slice[j].startTime;
}

func (slice LogEntries) Swap(i, j int) {
    slice[i], slice[j] = slice[j], slice[i]
}

func scanData(logEntries LogEntries) (LogEntries){
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
    return logEntries
}

func calc(query query, logEntries LogEntries) (float64) {
    var intervalTime float64
    var startTime, stopTime uint64
    for i:=0; i<len(logEntries); i++{
        if logEntries[i].startTime >= query.endTime {
            break
        }
        if logEntries[i].endTime <= query.startTime ||
            logEntries[i].startTime >= query.endTime {
            continue
        }

        if logEntries[i].startTime <= query.startTime {
            startTime = query.startTime
        } else {
            startTime = logEntries[i].startTime
        }
        if logEntries[i].endTime <= query.endTime {
            stopTime = logEntries[i].endTime
        } else {
            stopTime = query.endTime
        }

        fmt.Println("adding",logEntries[i])
        intervalTime += float64(float64(logEntries[i].bitRate)/float64(1000) * float64(stopTime-startTime))
    }
    return intervalTime
}

func main() {
    logEntries := LogEntries{}
    logEntries = scanData(logEntries)

    //fmt.Println(logEntries)
    sort.Sort(logEntries)
    //fmt.Println(logEntries)

    for _,element := range sliceOfQueries {
        fmt.Printf("%.3f\n",calc(element, logEntries))
    }
}