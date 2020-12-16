package lib

import (
	"fmt"
	"sort"
)

type Range struct {
	Min      int
	Max      int
	Metadata interface{}
}

func (r *Range) Contains(query int) bool {
	return query >= r.Min && query <= r.Max
}

func (r *Range) Valid() bool {
	return r.Min < r.Max
}

type IntervalTreeNode struct {
	CenterPoint          int
	CenterIntervalsByMin []*Range
	CenterIntervalsByMax []*Range
	Left                 *IntervalTreeNode
	Right                *IntervalTreeNode
}

func NewIntervalTree(input []*Range) (*IntervalTreeNode, error) {
	for _, interval := range input {
		if !interval.Valid() {
			return nil, fmt.Errorf("Invalid range %v", interval)
		}
	}
	return newIntervalTreeNode(input), nil
}

func newIntervalTreeNode(input []*Range) *IntervalTreeNode {
	if len(input) == 0 {
		return nil
	}
	minVal, maxVal := input[0].Min, input[0].Max
	for _, interval := range input {
		if interval.Min < minVal {
			minVal = interval.Min
		}
		if interval.Max > maxVal {
			maxVal = interval.Max
		}
	}

	centerpoint := (minVal + maxVal) / 2

	intervalsByMin := make([]*Range, 0)
	leftIntervals := make([]*Range, 0)
	rightIntervals := make([]*Range, 0)
	for _, interval := range input {
		if interval.Contains(centerpoint) {
			intervalsByMin = append(intervalsByMin, interval)
		} else if interval.Max < centerpoint {
			leftIntervals = append(leftIntervals, interval)
		} else { // must be to the right if not overlapping or to the left {
			rightIntervals = append(rightIntervals, interval)
		}
	}

	intervalsByMax := make([]*Range, len(intervalsByMin))
	copy(intervalsByMax, intervalsByMin)
	sort.Slice(intervalsByMin, func(i, j int) bool {
		return intervalsByMin[i].Min < intervalsByMin[j].Min
	})
	sort.Slice(intervalsByMax, func(i, j int) bool {
		return intervalsByMax[i].Max > intervalsByMax[j].Max
	})

	return &IntervalTreeNode{
		CenterPoint:          centerpoint,
		CenterIntervalsByMin: intervalsByMin,
		CenterIntervalsByMax: intervalsByMax,
		Left:                 newIntervalTreeNode(leftIntervals),
		Right:                newIntervalTreeNode(rightIntervals),
	}
}

func (i *IntervalTreeNode) Find(query int) []*Range {
	matches := make([]*Range, 0)
	return i.findInNode(query, matches)
}

func (i *IntervalTreeNode) findInNode(query int, matches []*Range) []*Range {
	if query <= i.CenterPoint {
		for _, interval := range i.CenterIntervalsByMin {
			if interval.Min > query {
				break
			}
			matches = append(matches, interval)
		}

		if i.Left != nil {
			return i.Left.findInNode(query, matches)
		}
	}

	if query >= i.CenterPoint {
		for _, interval := range i.CenterIntervalsByMax {
			if interval.Max < query {
				break
			}
			matches = append(matches, interval)
		}

		if i.Right != nil {
			return i.Right.findInNode(query, matches)
		}
	}

	return matches
}
