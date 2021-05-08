package main

import (
	"fmt"
	"github.com/achenet/TimeLog/data_store"
)

func main() {
	fmt.Println("TimeLog")
	fmt.Println(FirstBranchTest())
}

func FirstBranchTest() (tl data_store.TaskList) {
	tl = make(data_store.TaskList)
	tl["first"] = data_store.Task{
		Name: "first",
	}
	return tl
}
