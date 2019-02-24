package main

import "sort"

func main() {

}

type Interval struct {
	Start int
	End   int
}

/**
 * Definition for an interval.
 * type Interval struct {
 *	   Start int
 *	   End   int
 * }
 */
func merge(intervals []Interval) []Interval {
	res := make([]Interval, 0)
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i].Start == intervals[j].Start {
			return intervals[i].End < intervals[j].End
		}
		return intervals[i].Start < intervals[j].Start
	})
	pre := false
	tmp := Interval{}
	for _, v := range intervals {
		if !pre {
			tmp = v
			pre = true
		} else {
			if v.Start > tmp.End {
				res = append(res, tmp)
				tmp = v
			} else {
				end := tmp.End
				if v.End > end {
					end = v.End
				}
				tmp.End = end
			}
		}
	}
	if pre{
		res = append(res,tmp)
	}
	return res
}
