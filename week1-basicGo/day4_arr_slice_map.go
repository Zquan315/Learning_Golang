package main

import "fmt"

func arr_slice_map() {
	save := make(map[string]int)
	for {
		var w string
		fmt.Print("Enter a word (or 'esc' to quit): ")
		fmt.Scanln(&w)
		if w == "esc" {
			break
		} else {
			save[w]++
		}
	}
	fmt.Println("Word counts:")
	for word, count := range save {
		fmt.Printf("%s: %d\n", word, count)
	}
}
