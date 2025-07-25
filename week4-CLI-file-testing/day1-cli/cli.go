package cli_tool

import (
	fl "flag"
	f "fmt"
	o "os"
)

func cli() {
	if len(o.Args) < 2 {
		f.Println("❌ No Argument provided. Please provide a command to run.")
		return
	}
	for i, arg := range o.Args[1:] {
		f.Println("Argument:", i, "-> Value:", arg)
		if arg == "hi" {
			f.Println("👋 Hello! How can I help you today?")
		}
	}
}

func flag() {

	id := fl.String("id", "22521190", "ID of the user")
	name := fl.String("name", "To Cong Quan", "Name of the user")
	birthday := fl.String("dob", "31/01/2004", "Birthday of the user")
	fl.Parse()

	f.Println("Hi, I am", *name, "with ID", *id, "and my birthday is", *birthday)
}
func RunCLI() {
	//cli()
	flag()
	f.Println("⭐ Running CLI tool successful!\n-------------")
}
