package main

import (
	"fmt"

	pq "github.com/jupp0r/go-priority-queue"
)

func main() {
	priorityQueue := pq.New()
	priorityQueue.Insert("java", 1)
	priorityQueue.Insert("golang", 1)
	priorityQueue.Insert("php", 2)
	priorityQueue.UpdatePriority("java", 3)
	for priorityQueue.Len() > 0 {
		val, err := priorityQueue.Pop()
		if err != nil {
			panic(err)
		}
		fmt.Println(val)
	}

}
