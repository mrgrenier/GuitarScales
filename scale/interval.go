package scale

import "fmt"

type Interval struct {
	offset map[string]int
}

func NewInterval() *Interval {

	i := &Interval{}
	// Everything is off the major scale
	i.offset = make(map[string]int)
	i.offset["1"] = 0
	i.offset["b2"] = 1
	i.offset["2"] = 2
	i.offset["#2"] = 3
	i.offset["b3"] = 3
	i.offset["3"] = 4
	i.offset["4"] = 5
	i.offset["#4"] = 6
	i.offset["b5"] = 6
	i.offset["5"] = 7
	i.offset["#5"] = 8
	i.offset["b6"] = 8
	i.offset["6"] = 9
	i.offset["#6"] = 10
	i.offset["b7"] = 10
	i.offset["7"] = 11

	return i
}

func (i *Interval) IntervalToOffset(interval string) (int, error) {
	if inter, ok := i.offset[interval]; ok {
		return inter, nil
	}
	return 0, fmt.Errorf("%s, not a valid interval", interval)
}
