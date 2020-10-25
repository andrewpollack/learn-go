package main

import "fmt"

const _hash = 10000

type ListNode struct {
	Key  int
	Val  int
	Next *ListNode
}

type MyHashMap struct {
	table [_hash]*ListNode
}

func HashMapConstructor() MyHashMap {
	return MyHashMap{}
}

func (m *MyHashMap) Put(key int, value int) {
	var prev *ListNode

	hash := key % _hash
	curr := m.table[hash]
	node := ListNode{
		Key: key,
		Val: value,
	}

	if curr == nil {
		m.table[hash] = &node
		return
	}

	for curr != nil {
		if curr.Key == key {
			curr.Val = value
			return
		}

		prev = curr
		curr = curr.Next
	}

	prev.Next = &node
}

func (m *MyHashMap) Get(key int) int {
	hash := key % _hash
	curr := m.table[hash]

	for curr != nil {
		if curr.Key == key {
			return curr.Val
		}

		curr = curr.Next
	}

	return -1
}

func (m *MyHashMap) Remove(key int) {
	var prev *ListNode

	hash := key % _hash
	curr := m.table[hash]

	// Miss
	if curr == nil {
		return
	}

	// Remove head
	if curr.Key == key {
		m.table[hash] = curr.Next
		return
	}

	for curr != nil {
		if curr.Key == key {
			if prev == nil {
				m.table[hash] = nil
				return
			}

			prev.Next = curr.Next
			return
		}

		prev = curr
		curr = curr.Next
	}
}

func main() {
	obj := HashMapConstructor()
	key := 1
	val := 4

	obj.Put(key, val)
	getResult := obj.Get(key)
	fmt.Printf("Desired: %v, Found: %v\n", val, getResult)

	obj.Remove(key)
	getResult = obj.Get(key)
	fmt.Printf("Desired: %v, Found: %v\n", -1, getResult)
}
