package lib

import (
	"fmt"
	"sort"
)

// Range encapsulates the min and max values for an interval and allows the user
// to specify some sort of metadata.
type Range struct {
	Min      int
	Max      int
	Metadata interface{}
}

// Contains will establish that the given query falls in the bounds of the
// Range, inclusive of min/max.
func (r *Range) Contains(query int) bool {
	return query >= r.Min && query <= r.Max
}

// Valid will return true when the range Min val is strictly less than the range
// Max val.
func (r *Range) Valid() bool {
	return r.Min < r.Max
}

// IntervalTreeNode stores a set of intervals in a binary tree for simplified retrieval.
type IntervalTreeNode struct {
	CenterPoint          int
	CenterIntervalsByMin []*Range
	CenterIntervalsByMax []*Range
	Left                 *IntervalTreeNode
	Right                *IntervalTreeNode
}

// NewIntervalTree returns the root of an interval tree containing all
// specified intervals, provided all intervals are valid.
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

// Find returns a slice of intervals containing the given point, inclusive of bounds.
func (i *IntervalTreeNode) Find(query int) []*Range {
	matches := make([]*Range, 0)
	return i.findInNode(query, matches)
}

func (i *IntervalTreeNode) findInNode(query int, matches []*Range) []*Range {
	if query == i.CenterPoint {
		for _, interval := range i.CenterIntervalsByMin {
			matches = append(matches, interval)
		}

		if i.Left != nil {
			matches = i.Left.findInNode(query, matches)
		}
		if i.Right != nil {
			matches = i.Right.findInNode(query, matches)
		}
		return matches
	}

	if query < i.CenterPoint {
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

	if query > i.CenterPoint {
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
