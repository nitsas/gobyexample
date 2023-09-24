package main

import "fmt"

func MapKeys[K comparable, V any](m map[K]V) []K {
	r := make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}

type List[T any] struct {
	head, tail *element[T]
	size       int
}

type element[T any] struct {
	val  T
	next *element[T]
}

func ListPush[T any](l *List[T], v T) {
	if l.size == 0 {
		l.head = &element[T]{val: v}
		l.tail = l.head
		l.size++
	} else {
		l.tail.next = &element[T]{val: v}
		l.tail = l.tail.next
		l.size++
	}
}

func ToArray[T any](l List[T]) []T {
	ary := make([]T, 0, l.size)
	for cur := l.head; cur != nil; cur = cur.next {
		ary = append(ary, cur.val)
	}
	return ary
}

func main() {
	m := map[int]string{15: "fifteen", 13: "thirteen", 11: "eleven", 17: "seventeen"}
	fmt.Println("m:", m)
	fmt.Println("MapKeys(m):", MapKeys(m))
	fmt.Println("MapKeys[int, string](m):", MapKeys[int, string](m))
	fmt.Println()

	l := List[int]{}
	fmt.Println("ToArray(l):", ToArray(l))
	ListPush(&l, 10)
	fmt.Println("ListPush(&l, 10) -> ToArray(l):", ToArray(l))
	ListPush(&l, 20)
	fmt.Println("ListPush(&l, 20) -> ToArray(l):", ToArray(l))
	ListPush(&l, 30)
	fmt.Println("ListPush(&l, 30) -> ToArray(l):", ToArray(l))
}
