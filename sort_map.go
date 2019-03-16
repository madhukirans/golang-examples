package main

import (
	"fmt"

	"sort"
)

func main() {
  a := make (map[string]int)
  a["x"] =1
  a["a"] = 2
  fmt.Println(a)

  type Entry struct {
  	str string
  	val int
  }
  s := make ([]Entry, len(a))
  for k,v := range a {
  	s = append(s, Entry{str:k, val:v})
  }

  fmt.Println(s)
  sort.Slice(s, func( i,  j int) bool {
  	return s[i].val > s[j].val
  })

  fmt.Println(s)

}

