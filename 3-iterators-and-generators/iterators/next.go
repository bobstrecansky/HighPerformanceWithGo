package main

import "fmt"

const top = 100

type CounterStruct struct {
	err error
	max int
	cur int
}

func NewCounterIterator(top int) *CounterStruct {
	var err error
	return &CounterStruct{
		err: err,
		max: top,
		cur: 0,
	}
}

func (i *CounterStruct) Next() bool {
	if i.err != nil {
		return false
	}

	i.cur++
	return i.cur <= i.max
}

func (i *CounterStruct) Value() int {
	if i.err != nil || i.cur > i.max {
		panic("Value is not valid after iterator finished")
	}
	return i.cur
}

func nextLoop(top int) {
	nextIterator := NewCounterIterator(top)
	for nextIterator.Next() {
		fmt.Println(nextIterator.Value())
	}
}

func main() {
	nextLoop(top)
}
