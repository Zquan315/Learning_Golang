package main

import (
	f "fmt"
)

func main() {
	f.Println("----------------Running arr_slice_map function----------------")
	arr_slice_map()
	f.Println("\n--------------Running learn_closure function--------------")
	learn_closure_main()
	f.Println("\n--------------Running RunReadFile function--------------")
	filename := "../README.md"
	RunReadFile(filename)
	f.Println("\n--------------Running RunPAnic function--------------")
	RunPAnic()
	f.Println("\n--------------Running RunBook function--------------")
	RunBook()
	f.Println("\n--------------Running RunSquare function--------------")
	RunSquare()
}
