package main

import (
	"fmt"
)


type testS struct {
	a int `json:"a"`
}



func main() {

	arr1 := []int{1, 2, 3}
	arr2 := []*int{}
	a := 1
	p := &a
	for _, i := range arr1 {
		*p = i
		arr2 = append(arr2, p)
	}
	fmt.Print(arr2)
	return
}
