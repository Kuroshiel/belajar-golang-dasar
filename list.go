package main

import (
	"container/list"
	"fmt"
)

func main() {
	var data *list.List = list.New()

	data.PushBack("Eko")
	data.PushBack("Kurniawan")
	data.PushBack("Khennndy")

	var head *list.Element = data.Front()
	fmt.Println(head.Value) // Eko

	next := head.Next()
	fmt.Println(next.Value) // Kurniawan

	next = head.Next()
	fmt.Println(next.Value) //Khannedy

	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

}
