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

        intervalTime += float64(float64(logEntries[i].bitRate)/float64(1000) * float64(stopTime-startTime))
    }
    return intervalTime
}

func removeEarlierPart(logEntries LogEntries, queryStartTime uint64)(LogEntries) {
    sortByEndtime = true
    sort.Sort(logEntries)
    low := 0
    high := len(logEntries)-1
    mid := 0

    if logEntries[low].endTime > queryStartTime {
        return logEntries
    }

    for high - low > 1{
        mid = low + (high-low)/2

        //fmt.Println(low,mid,high)

        if high-low == 1 || high-low==2 {
            if logEntries[mid].endTime >= queryStartTime {
                return logEntries[mid:]
            }
            if logEntries[mid+1].endTime >= queryStartTime {
                return logEntries[mid+1:]
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

    //fmt.Println("removeLater")
    //fmt.Println(logEntries)

    if logEntries[high].startTime < queryEndTime {
        return logEntries
    }

    for high - low > 1 {
        mid = low + (high-low)/2

        //fmt.Println(low,mid,high)

        if high-low == 1 || high-low==2 {
            if logEntries[mid + 1].startTime >= queryEndTime {
                return logEntries[0:mid + 1]
            }
            if logEntries[mid].startTime >= queryEndTime {
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

func main() {

    logEntries := LogEntries{}
    var sliceOfQueries = make([]query, 0)

    logEntries, sliceOfQueries = scanData(logEntries, sliceOfQueries)

    sortByEndtime = false
    sort.Sort(logEntries)
    logEntriesToCalc := LogEntries{}

    for _,element := range sliceOfQueries {
        logEntriesToCalc = removeLaterPart(logEntries, element.endTime)
        //fmt.Println("after removeLater")
        //fmt.Println(logEntriesToCalc)

        logEntriesToCalc = removeEarlierPart(logEntriesToCalc, element.startTime)
        //fmt.Println("after removeEarlier")
        //fmt.Println(logEntriesToCalc)

        fmt.Printf("%.3f\n",calc(element, logEntriesToCalc))
    }
}