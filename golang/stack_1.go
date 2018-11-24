package main

import (
	"fmt"
	"time"
)
const ( 
	stime = time.Second * 5
	sizeInit = 10000
	text = "text"
)

type Stack []interface{}

func (s *Stack) Push(element interface{}) {
	*s = append(*s, element)
}

func (s *Stack) Len() int {
	return len(*s)
}

func (s Stack) Top() interface{} {
	if len(s) > 0 {
		return s[len(s)-1]
	}
	return nil
}

func (s *Stack) Pop() interface{} {
	tmp := *s
	if len(tmp) == 0 {
		return nil
	}
	*s = tmp[:len(tmp)-1]
	return tmp[len(tmp)-1]
}

func add1(s Stack) {
	for i := 0; i < sizeInit; i++ {
		s.Push(fmt.Sprintf("ala__1_%s_%d", text, i))
	}
}

func add2(s *Stack) {
	for i := 0; i < sizeInit; i++ {
		s.Push(fmt.Sprintf("ala__2_%s_%d", text, i))
	}
}

func main() {
	p := fmt.Printf
	p("\nstart")
	
	f1()
	p("\nsleep")
	time.Sleep(stime)
	p("\nend")
}

func f2() {
	p := fmt.Printf	
	s := &Stack{}
	p("\nLen: %d", s.Len())
	
		//add1(*s)
		add2(s)
	
	
	p("\nLen: %d", s.Len())
	p("\nLen: %d, top=%v", s.Len(), s.Top())
	p("\nLen: %d, pop=%v, len: %d", s.Len(), s.Pop(), s.Len())
	p("\nLen: %d, pop=%v, len: %d", s.Len(), s.Pop(), s.Len())
}

func f1() {
	p := fmt.Printf	
	s := Stack{}
	p("\nLen: %d", s.Len())
	s.Push("ala")
	s.Push("ma")
	p("\nLen: %d, top=%v", s.Len(), s.Top())
	p("\nLen: %d, pop=%v, len: %d", s.Len(), s.Pop(), s.Len())
	p("\nLen: %d, pop=%v, len: %d", s.Len(), s.Pop(), s.Len())
	

}

