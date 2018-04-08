package main

import (
	"fmt"
	"sort"

	"github.com/golang-collections/collections/stack"
)

//If there is a set of Jobs with given start and end time, we need to find a set where
//we can perform maximum number of jobs that are nit overlapping.
//Example:
//   jobs = [{2, 8},{3, 5},{2, 4},{10, 13},{14, 17},{9, 11}]
//   Output = [{2 4} {9 11} {14 17}]

func getOptimalJobSet(intervals []interval) []interval {
	var result []interval = make([]interval, 0)
	if intervals == nil || len(intervals) <= 0 {
		fmt.Println("Invalid input")
		return nil
	}
	tempStack := stack.New()
	jobs := SortByStartTime(intervals)

	tempStack.Push(jobs[0])

	for i := 1; i < len(jobs); i++ {
		old := tempStack.Peek().(interval)
		if isOverlapping(old, jobs[i]) {
			if old.EndTime > jobs[i].EndTime {
				tempStack.Pop()
				tempStack.Push(jobs[i])
			}
		} else {
			result = append(result, tempStack.Pop().(interval))
			tempStack.Push(jobs[i])
		}
	}

	if tempStack.Len() > 0 {
		result = append(result, tempStack.Pop().(interval))
	}
	return result
}

func isOverlapping(old interval, current interval) bool {
	var res bool
	oldStartTime := old.StartTime
	oldEndTime := old.EndTime
	if current.StartTime >= oldStartTime && (current.StartTime <= oldEndTime) {
		res = true
	} else {
		res = false
	}
	return res
}

type interval struct {
	StartTime int
	EndTime   int
}

type Intervals []interval

func (intervals Intervals) Len() int      { return len(intervals) }
func (intervals Intervals) Swap(i, j int) { intervals[i], intervals[j] = intervals[j], intervals[i] }

// ByStartTime implements sort.Interface by providing Less and using the Len and
// Swap methods of the embedded Intervals value.
type ByStartTime struct{ Intervals }

func (b ByStartTime) Less(i, j int) bool { return b.Intervals[i].StartTime < b.Intervals[j].StartTime }

func SortByStartTime(intervals []interval) []interval {
	sort.Sort(ByStartTime{intervals})
	return intervals
}
