package main

import "fmt"

var mapSize int = 50

type Node struct {
	Key   string
	Value interface{}
	Next  *Node
}
type HashMap struct {
	Arr []*Node
}

func NewHashMap() *HashMap {
	return &HashMap{Arr: make([]*Node, mapSize)}
}
func hash(key string) (hash uint32) {
	hash = 0
	for _, ch := range key {
		hash += uint32(ch)
		hash += hash << 10
		hash ^= hash >> 6
	}
	hash += hash << 3
	hash ^= hash >> 11
	hash += hash << 15
	return
}
func getIndex(key string) (index int) {
	return int(hash(key)) % mapSize
}
func main() {
	a := NewHashMap()
	a.Insert("name", "ishan")
	a.Insert("gender", "male")
	a.Insert("city", "mumbai")
	a.Insert("lastname", "khare")
	if value, ok := a.Get("name"); ok {
		fmt.Println(value)
	} else {
		fmt.Println("Value did not match!")
	}

}

//insert/update
func (h *HashMap) Insert(key string, value interface{}) {
	index := getIndex(key)
	if h.Arr[index] == nil {
		//do insert
		h.Arr[index] = &Node{Key: key, Value: value}
	} else {
		//collision
		startTraverse := h.Arr[index]
		for ; startTraverse.Next != nil; startTraverse = startTraverse.Next {
			if startTraverse.Key == key {
				//update operation
				startTraverse.Value = value
			}
		}
		startTraverse.Next = &Node{Key: key, Value: value}
	}
}
func (h *HashMap) Get(key string) (interface{}, bool) {
	index := getIndex(key)
	if h.Arr[index] != nil {
		if h.Arr[index].Key == key {
			return h.Arr[index].Value, true
		} else {
			startTraverse := h.Arr[index]
			for ; ; startTraverse = startTraverse.Next {
				if startTraverse.Key == key {
					return startTraverse.Value, true
				}
				if startTraverse.Next == nil {
					break
				}
			}
		}
	}
	//does not exist
	return nil, false
}
