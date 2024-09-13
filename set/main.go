package main

import (
	"fmt"
	"my-go-playground/set/pkg"
)

func main() {
	empty := ""

	set := pkg.NewStringSet(0)
	set.Insert("a")
	set.Insert("b")
	set.InsertIfNotNil(&empty)
	set.InsertIfNotZeroValue("")

	fmt.Println(set.GetItems())
}
