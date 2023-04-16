package main

import "fmt"

// 단일 연결 리스트의 노드를 나타내는 구조체
type Node struct {
	data int   // 데이터 필드
	next *Node // 다음 노드를 가리키는 포인터
}

// 새로운 노드를 생성하는 함수
func newNode(data int) *Node {
	return &Node{data: data, next: nil}
}

// 단일 연결 리스트를 나타내는 구조체
type LinkedList struct {
	head *Node // 맨 앞 노드를 가리키는 포인터
}

// 새로운 빈 단일 연결 리스트를 생성하는 함수
func newLinkedList() *LinkedList {
	return &LinkedList{head: nil}
}

// 단일 연결 리스트의 맨 앞에 노드를 추가하는 함수
func (list *LinkedList) prepend(data int) {
	newNode := newNode(data)

	// 새로운 노드가 맨 앞 노드가 되도록 포인터 조정
	newNode.next = list.head
	list.head = newNode
}

//  (deprecated)
// 단일 연결 리스트의 중간에 노드를 추가하는 함수
// func (list *LinkedList) insertAfter(data int, prevNode *Node) {
// 	if prevNode == nil {
// 		fmt.Println("이전 노드가 nil입니다.")
// 		return
// 	}

// 	for list.head.next == nil {

// 	}
// 	newNode := newNode(data)

// 	// 이전 노드의 다음 노드를 새로운 노드의 다음 노드로 연결
// 	newNode.next = prevNode.next

// 	// 이전 노드의 다음 노드가 새로운 노드가 되도록 연결
// 	prevNode.next = newNode
// }

// 단일 연결 리스트의 맨 뒤에 노드를 추가하는 함수
func (list *LinkedList) append(data int) {
	newNode := newNode(data)

	// 빈 리스트인 경우
	if list.head == nil {
		list.head = newNode
		return
	}

	// 마지막 노드까지 이동하여 새로운 노드를 추가
	lastNode := list.head
	for lastNode.next != nil {
		lastNode = lastNode.next
	}
	lastNode.next = newNode
}

// 단일 연결 리스트의 모든 노드를 출력하는 함수
func (list *LinkedList) print() {
	currentNode := list.head
	for currentNode != nil {
		fmt.Printf("%d ", currentNode.data)
		currentNode = currentNode.next
	}
	fmt.Println()
}

// 예제 실행 함수
func main() {
	list := newLinkedList()
	list.append(1)
	list.append(3)
	list.print() // 1 3

	list.prepend(0)
	list.print() // 0 1 3

}
