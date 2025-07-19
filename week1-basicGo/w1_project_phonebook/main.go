package main

import (
	f "fmt"
)

type Person struct {
	Name  string
	Phone string
}

func (p *Person) addPerson(lst map[string]string) {
	f.Print("Type name: ")
	f.Scan(&p.Name)
	f.Print("Type phone: ")
	f.Scan(&p.Phone)
	lst[p.Name] = p.Phone
	f.Println("Added person:", p.Name, "with phone:", p.Phone)
}

func (p *Person) showPhoneBook(lst map[string]string) {
	f.Println("------Phonebook------")
	for name, phone := range lst {
		f.Printf("%s: %s\n", name, phone)
	}
}

func (p *Person) deletePerson(lst map[string]string) {
	f.Print("Type name to delete: ")
	var name string
	f.Scan(&name)
	if _, exists := lst[name]; exists {
		delete(lst, name)
		f.Println("Deleted person:", name)
	} else {
		f.Println("Person not found:", name)
	}
}

func updatePerson(lst map[string]string) {
	f.Print("Type name to update: ")
	var name string
	f.Scan(&name)
	if phone, exists := lst[name]; exists {
		f.Print("Current phone is ", phone, ".\nType new phone: ")
		var newPhone string
		f.Scan(&newPhone)
		lst[name] = newPhone
		f.Println("Updated person:", name, "to new phone:", newPhone)
	} else {
		f.Println("Person not found:", name)
	}
}
func main() {
	f.Println("✨----------------Running phonebook application----------------✨")
	lst := make(map[string]string)
	var p Person
	var choose int
	for {
		f.Println("❌ 0. Exit")
		f.Println("💡 1. Add person")
		f.Println("💡 2. Show phonebook")
		f.Println("💡 3. Delete person")
		f.Println("💡 4. Update person")

		f.Print("Choose an option: ")
		f.Scan(&choose)

		switch choose {
		case 0:
			f.Println("🔥 Exiting phonebook application.")
			return
		case 1:
			p.addPerson(lst)
		case 2:
			p.showPhoneBook(lst)
		case 3:
			p.deletePerson(lst)
		case 4:
			updatePerson(lst)
		default:
			f.Println("❌ Invalid option. Please choose a valid option.")
		}
	}
}
