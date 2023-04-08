package main

import "fmt"

type CirQueue struct {
	array       [10]int
	front, rear int
}

func (c *CirQueue) isEmpty() bool {
	return c.front == c.rear
}

func (c *CirQueue) isFull() bool {

	return ((c.rear + 1) % cap(c.array)) == c.front

}

// 큐에서 데이터를 삽입한다.
func (c *CirQueue) enqueue(data int) {
	if c.isFull() {
		fmt.Println("Q is full")
		return
	}

	fmt.Println(c.rear, data)
	// 원형 큐이기 때문에 SIZE보다 크면 다시 처음부터 원소가 삽입되어야하므로 모퓰러 연산이 들어간다.
	c.rear = (c.rear + 1) % cap(c.array)

	c.array[c.rear] = data
}

// 큐에서 데이터를 꺼내온다.
func (c *CirQueue) dequeue() int {
	if c.isEmpty() {
		fmt.Println("Q is Empty")
		return 0
	}

	// 원형 큐이기 때문에 SIZE보다 크면 다시 처음으로 돌아가 원소를 꺼내와야 되므로 모퓰러 연산이 들어간다.
	c.front = (c.front + 1) % cap(c.array)

	return c.array[c.front]
}

func main() {

	var q CirQueue

	q.front, q.rear = 9, 9

	for i := 1; i < 12; i++ {
		q.enqueue(i)
	}

}

/**
* 참고 : https://reakwon.tistory.com/30
**/
