package main

import f "fmt"

type Animal interface {
	speak() string
	describe() string
}

type Dog struct {
	speakSound string
}

func (d Dog) speak() string {
	return d.speakSound
}
func (d Dog) describe() string {
	des := f.Sprintln("I'm a Dog. My sound is " + d.speakSound)
	return des
}

type Cat struct {
	speakSound string
}

func (c Cat) speak() string {
	return c.speakSound
}
func (c Cat) describe() string {
	des := f.Sprintln("I'm a Cat. My sound is " + c.speakSound)
	return des
}
func RunInterface() {
	var dog Animal = Dog{speakSound: "Woof!"}
	var cat Animal = Cat{speakSound: "Meow!"}
	f.Println("Dog says:", dog.speak())
	f.Println("Dog description:", dog.describe())
	f.Println("Cat says:", cat.speak())
	f.Println("Cat description:", cat.describe())
}

//
type testEmptyInterface interface{}

func PrintAnyInterface(i testEmptyInterface) {
	switch v := i.(type) {
	case int:
		f.Println("Integer:", v)
	case string:
		f.Println("String:", v)
	case float64:
		f.Println("Float64:", v)
	case bool:
		f.Println("Boolean:", v)
	default:
		f.Println("Unknown type:", v)
	}
}
func RunAdvancedInterface() {
	profile := make(map[string]testEmptyInterface)
	profile["Name"] = "To Cong Quan"
	profile["Age"] = 30
	profile["Height"] = 1.72
	profile["IsStudent"] = true
	profile["Hobbies"] = []string{"Gaming", "Coding", "Traveling"}
	f.Print("Profile Information:\n")
	for key, value := range profile {
		f.Print(key, ": ", value, "\n")
	}
}
