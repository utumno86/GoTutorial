package main

import (
	"fmt"
	"sort"
)

type people []string

func (p people) Len() int           { return len(p) }
func (p people) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p people) Less(i, j int) bool { return p[i] > p[j] }

func main() {
	studyGroup := people{"Zeno", "John", "Al", "Jenny"}
	fmt.Println(studyGroup)
	sort.Sort(studyGroup)
	fmt.Println(studyGroup)
	fmt.Println("*************")

	s := []string{"Zeno", "John", "Al", "Jenny"}
	fmt.Println(s)
	sort.StringSlice(s).Sort()
	fmt.Println(s)

	fmt.Println("************")
	sort.Sort(sort.Reverse(sort.StringSlice(s)))
	fmt.Println(s)
}
