package main

import "testing"

func TestCalc(t *testing.T) {
    query := query{startTime: 30, endTime: 50}
    logEntries := LogEntries{}

    logEntries = append(logEntries, logEntry{1, 60, 59, 1})
    logEntries = append(logEntries, logEntry{1, 40, 39, 1})
    logEntries = append(logEntries, logEntry{40, 60, 20, 1})
    logEntries = append(logEntries, logEntry{1, 20, 19, 1})
    logEntries = append(logEntries, logEntry{60, 70, 10, 1})
    logEntries = append(logEntries, logEntry{1, 30, 29, 1})
    logEntries = append(logEntries, logEntry{50, 70, 20, 1})
    logEntries = append(logEntries, logEntry{30, 50, 20, 1})
    logEntries = append(logEntries, logEntry{30, 40, 10, 1})
    logEntries = append(logEntries, logEntry{40, 50, 10, 1})
    logEntries = append(logEntries, logEntry{35, 45, 10, 1})
    logEntries = append(logEntries, logEntry{1, 1, 0, 1})
    logEntries = append(logEntries, logEntry{30, 30, 0, 1})
    logEntries = append(logEntries, logEntry{40, 40, 0, 1})
    logEntries = append(logEntries, logEntry{50, 50, 0, 1})
    logEntries = append(logEntries, logEntry{60, 60, 0, 1})

    result := calc(query, logEntries)
    expcted_result := 0.08999999999999998

    if result != expcted_result {
        t.Error("Expected 20, got", result)
    }
}
