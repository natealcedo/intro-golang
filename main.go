package main

import "fmt"

type CustomMap[K comparable, V any] struct {
	data map[K]V
}

func NewCustomMap[K comparable, V any]() *CustomMap[K, V] {
	return &CustomMap[K, V]{data: make(map[K]V)}
}

func (m *CustomMap[K, V]) Insert(k K, v V) error {
	m.data[k] = v
	return nil
}

func foo[T any](val T) {
	fmt.Println(val)
}

func main() {
	//m1 := NewCustomMap[string, int]()
	//m1.Insert("foo", 1)
	//m1.Insert("bar", 2)
	//
	//m2 := NewCustomMap[int, float64]()
	//m2.Insert(1, 9.7)
	//m2.Insert(5, 9.7)
	//
	//fmt.Printf("m1 := %+v\n", m1)
	//fmt.Printf("m2 := %+v\n", m2)

	foo("hello world")
}
