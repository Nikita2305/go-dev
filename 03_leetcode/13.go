package main

/*
https://leetcode.com/problems/number-of-recent-calls/?envType=study-plan-v2&envId=leetcode-75
*/

import (
	"fmt"
)

type Queue struct {
	begin int
	end int
    buffer []int
}

func (queue *Queue) Size() int {
	return (len(queue.buffer) + queue.end - queue.begin) % len(queue.buffer)
}

func (queue *Queue) Front() (int, error) {
	if (queue.begin == queue.end) {
		return 0, fmt.Errorf("Empty queue")
	}
	return queue.buffer[queue.begin], nil
}

func (queue *Queue) Push(value int) error {
	nextEnd := (queue.end + 1) % len(queue.buffer) 
	if (nextEnd == queue.begin) {
		return fmt.Errorf("Not enough buffer")
	}
	queue.buffer[queue.end] = value
	queue.end = nextEnd
	return nil
}

func (queue *Queue) Pop() (int, error) {
	if (queue.begin == queue.end) {
		return 0, fmt.Errorf("Empty queue")
	}
	value := queue.buffer[queue.begin]
	queue.begin = (queue.begin + 1) % len(queue.buffer)
	return value, nil
}

type RecentCounter struct {
	maxDiff int
	queue Queue
}

func Constructor() RecentCounter {
	maxDiff := 3000
    return RecentCounter{maxDiff, Queue{begin: 0, end: 0, buffer: make([]int, maxDiff + 10)}}
}


func (this *RecentCounter) Ping(t int) int {
	for this.queue.Size() > 0 {
		front, err := this.queue.Front()
		if err != nil {
			panic("unexpected")
		}
		if (front >= t - this.maxDiff) {
			break
		}
		this.queue.Pop()
	}
	this.queue.Push(t)
	return this.queue.Size()
}

func main() {

}