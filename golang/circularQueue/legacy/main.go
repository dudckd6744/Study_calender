package main

import "fmt"

var slice1 = [10]int{}

var size = len(slice1)

var front, rear int = 0, 0

func isEmpty() bool {
	return front == rear
}

func isFull() bool {

	return ((rear + 1) % size) == front

}

// 큐에서 데이터를 삽입한다.
func enqueue(data int) {
	if isFull() {
		fmt.Println("Q is full")
		return
	}
	// 원형 큐이기 때문에 SIZE보다 크면 다시 처음부터 원소가 삽입되어야하므로 모퓰러 연산이 들어간다.
	rear = (rear + 1) % size

	slice1[rear] = data
}

// 큐에서 데이터를 꺼내온다.
func dequeue() int {
	if isEmpty() {
		fmt.Println("Q is Empty")
		return 0
	}

	// 원형 큐이기 때문에 SIZE보다 크면 다시 처음으로 돌아가 원소를 꺼내와야 되므로 모퓰러 연산이 들어간다.
	front = (front + 1) % size

	return slice1[front]
}

func main() {

	for i := 0; i < 100; i++ {
		enqueue(i)
		dequeue()
	}
}

/**
* 참고 : https://reakwon.tistory.com/30
**/
