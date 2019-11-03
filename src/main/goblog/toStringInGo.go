package main

import "fmt"

type Student struct {
	Name string
	Age  int
}

func (s Student) String() string {
	return fmt.Sprintf("%v (%v years)", s.Name, s.Age)
}

func main() {
	s := Student{"Boyn", 21}
	fmt.Println(s)
}
