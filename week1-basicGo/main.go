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
	RunReadFile("../README.md")
	f.Println("\n--------------Running RunPAnic function--------------")
	RunPAnic()
}
