package linkedlist

import "fmt"

/*
node
  data
  next -> node
            data
            next -> node
                      data
                      next -> nil
*/

type Node struct {
	Next *Node
	Data string
}

type LinkedList struct {
	head  *Node
	length int
	tail *Node
}

func New() *LinkedList {
	ll := &LinkedList{}

	return ll
}

func (ll *LinkedList) Push(num string) {
	
	// если длина равна нулю, создаем новую ноду
	// кладем ее указатель в поле ноудс
	// увеличиваем длину списка на один и выходим

	next := &Node{
		Next: nil, // это можно не указывать, так как это(nil) будет значение по умолчанию
		Data: num,
	}

	defer func() {
		ll.length++
	}()
	
	if ll.length == 0 {
		ll.head = next
		ll.tail = next
		return
	}

	ll.tail.Next = next
	ll.tail = next
}

func (ll *LinkedList) Length() int {
	return ll.length
}

func (ll *LinkedList) Print() {
	if ll.length == 0 {
		return
	}

	var save *Node = ll.head
	for save != nil {
		fmt.Println(save.Data)
		save = save.Next
	}
}
// list := {"["}
func (ll *LinkedList) Pop() (string, bool) {
	if ll.length == 0 {
		return "", false
	}

	defer func() {
		ll.length--
	}()

	if ll.length == 1 {
		num := ll.tail.Data
		ll.head = nil
		ll.tail = nil
		return num, true
	}

	var data string = ll.tail.Data
	var save *Node = ll.head
	for save.Next.Next != nil {
		save = save.Next
	}

	save.Next = nil
	ll.tail = save

	return data, true
	
	// взять значение последнего элемента
	// дойти до предпоследнего элемента
	// у него указатель в next установить в nil
	// указатель на предпоследний элемент положить в tail
	// сделать декремент длины списка
}

func (ll *LinkedList) PushHead(num string) () {
	next := &Node{
		Next: ll.head,
		Data: num,
	}
	
	ll.head = next

	if ll.length == 0 {
		ll.tail = next
	}
	
	ll.length++

	// в текущую ноду next положить первую ноду списка
	// после этого в head положить текущую ноду
}