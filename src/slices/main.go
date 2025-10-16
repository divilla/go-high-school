package main

import "fmt"

func main() {
	lesson03()
}

func lesson01() {
	s1 := []int{0, 1, 2}
	fmt.Printf("ptr:%p values:%v\n", s1, s1)

	s2 := s1
	fmt.Printf("ptr:%p values:%v\n", s2, s2)

	s2[0] = 3
	fmt.Printf("ptr:%p values:%v\n", s1, s1)
}

func lesson02() {
	s1 := []int{0, 1, 2}
	fmt.Printf("ptr:%p values:%v\n", s1, s1)

	s2 := make([]int, len(s1))
	copy(s2, s1)
	fmt.Printf("ptr:%p values:%v\n", s2, s2)

	s2[0] = 3
	fmt.Printf("ptr:%p values:%v\n", s1, s1)
}

func lesson03() {
	s1 := []int{0, 1, 3, 4}
	fmt.Printf("ptr:%p values:%v\n", s1, s1)

	i := 2
	v := 2
	s2 := s1[0:i]
	s2 = append(s2, v)
	fmt.Printf("ptr:%p values:%v\n", s1[i:len(s1)], s1[i:len(s1)])
	s2 = append(s2, s1[i:len(s1)]...)
	fmt.Printf("ptr:%p values:%v\n", s2, s2)
}

func lesson10() {
	s := make([]int, 10)
	for i := 0; i < len(s); i++ {
		fmt.Printf("value:%v\n", s[i])
	}
}

func lesson12() {
	s := make([]int, 0, 10)
	for i := 0; i < len(s); i++ {
		fmt.Printf("value:%v\n", s[i])
	}
}

func lesson13() {
	s := make([]int, 0)
	fmt.Printf("ptr:%p len:%d, cap:%d, values:%v\n", s, len(s), cap(s), s)
	for i := 0; i < 10; i++ {
		s = append(s, i)
		fmt.Printf("ptr:%p len:%d, cap:%d, values:%v\n", s, len(s), cap(s), s)
	}
}

func lesson14() {
	s := make([]int, 0, 10)
	fmt.Printf("ptr:%p len:%d, cap:%d, values:%v\n", s, len(s), cap(s), s)
	for i := 0; i < cap(s); i++ {
		s = append(s, i)
		fmt.Printf("ptr:%p len:%d, cap:%d, values:%v\n", s, len(s), cap(s), s)
	}
}

func lesson05() {
	s := make([]int, 10)
	fmt.Printf("ptr:%p len:%d, cap:%d, values:%v\n", s, len(s), cap(s), s)
	for i := 0; i < len(s); i++ {
		s = append(s, i)
		fmt.Printf("ptr:%p len:%d, cap:%d, values:%v\n", s, len(s), cap(s), s)
	}
}
