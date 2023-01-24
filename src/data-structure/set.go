package main

import (
	"errors"
	"fmt"
)

type Set struct {
	Elements map[string]struct{}
}

func NewSet() *Set {
	var set Set
	set.Elements = make(map[string]struct{})
	return &set
}

func (s *Set) Add(key string) {
	s.Elements[key] = struct{}{}
}

func (s *Set) Delete(key string) error {
	if _, exists := s.Elements[key]; !exists {
		return errors.New("element not present in set")
	}

	delete(s.Elements, key)
	return nil
}

func (s *Set) Contains(key string) bool {
	_, exists := s.Elements[key]
	return exists
}

func (s *Set) List() {
	for k, _ := range s.Elements {
		fmt.Println(k)
	}
}

func main() {
	mySet := NewSet()

	mySet.Add("joonseo")
	mySet.Add("joonha")
	mySet.Add("joonseo")

	fmt.Println(mySet.Contains("joonha"))
	fmt.Println(mySet.Contains("joonha1"))

	mySet.List()
}
