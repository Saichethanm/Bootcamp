package main

import "fmt"

type Student struct {
	name   string
	grades []int
	age    int
}

func (s Student) getAge() int {
	return s.age
}

func (s *Student) setAge(age int) {
	s.age = age
}

func (s *Student) getAverageGrade() float32 {
	sum := 0
	for _, x := range s.grades {
		sum += x
	}
	return float32(sum) / float32(len(s.grades)*1.0)
}
func main() {
	s1 := Student{"student1", []int{70, 89, 90, 85}, 19}
	fmt.Println(s1.getAge())
	s1.setAge(26)
	fmt.Println(s1.getAge())
	fmt.Println(s1.getAverageGrade())
}
