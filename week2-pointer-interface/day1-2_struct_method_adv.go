package main

import f "fmt"

type Student struct {
	id   string
	name string
	age  int
}

func (s Student) addStudent(list *[]Student) {
	f.Println("Enter student ID, Name, and Age:")
	var id, name string
	var age int
	f.Scan(&id, &name, &age)
	s.id = id
	s.name = name
	s.age = age
	*list = append(*list, s)
	f.Printf("Added %s student!\n", s.name)
}

func (s Student) displayStudents(list *[]Student) {
	f.Println("List of Students:")
	for _, s := range *list {
		f.Printf("ID: %s, Name: %s, Age: %d\n", s.id, s.name, s.age)
	}
}

func (s *Student) updateAge(addAge int, list *[]Student) {
	var findID string
	f.Print("Enter student ID to update age: ")
	f.Scan(&findID)
	flag := false
	var index int
	for i, st := range *list {
		if st.id == findID {
			s = &st
			index = i
			flag = true
			break
		}
	}
	if !flag {
		f.Println("Student not found!")
		return
	}
	s.age += addAge
	(*list)[index].age = s.age
	f.Printf("Updated age of %s having id %s to %d\n", s.name, s.id, s.age)
}

func RunStudentManagement() {
	var n int
	f.Print("Enter number of students: ")
	f.Scan(&n)
	var s Student
	list := make([]Student, 0)
	for i := 0; i < n; i++ {
		f.Println("Adding student", i+1)
		s.addStudent(&list)
	}
	f.Println("\n--Displaying Students--")
	f.Println("Total students:", len(list))
	s.displayStudents(&list)
	f.Println("\n--Updating Age--")
	var addAge int
	f.Print("Enter age to add: ")
	f.Scan(&addAge)
	s.updateAge(addAge, &list)
	s.displayStudents(&list)
}

type BankAccount struct {
	accountNumber string
	balance       float64
}

func (b *BankAccount) RechargeAccount(accounts *[]BankAccount, amount float64, index int) {
	for i, acc := range *accounts {
		if i == index {
			b := &acc
			b.balance += amount
			(*accounts)[i] = *b
			break
		}
	}
}

func (b BankAccount) WithdrawFromAccount(accounts *[]BankAccount, amount float64, index int) {
	for i, acc := range *accounts {
		if i == index {
			b := &acc
			b.balance -= amount
			(*accounts)[i] = *b
			break
		}
	}
}
func RunBanking() {
	var n int
	f.Print("Enter number of bank accounts: ")
	f.Scan(&n)
	var account BankAccount
	accounts := make([]BankAccount, 0)
	for i := 0; i < n; i++ {
		f.Printf("Enter account number and initial balance for each account number %d: ", i)
		var id string
		var balance float64
		f.Scan(&id, &balance)
		account.accountNumber = id
		account.balance = balance
		accounts = append(accounts, account)
	}
	var choice int
	for {
		f.Println("❌ 0. Exit")
		f.Println("💡 1. Recharge Account")
		f.Println("💡 2. Withdraw from Account")
		f.Println("💡 3. Display Accounts")
		f.Print("✔️ Enter your choice: ")
		f.Scan(&choice)

		switch choice {
		case 1:
			check := false
			f.Print("Enter account number to recharge: ")
			var accNum string
			f.Scan(&accNum)
			for i, acc := range accounts {
				if acc.accountNumber == accNum {
					var rechargeAmount float64
					f.Print("Enter amount to recharge: ")
					f.Scan(&rechargeAmount)
					account.RechargeAccount(&accounts, rechargeAmount, i)
					check = true
					break
				}
			}
			if !check {
				f.Printf("Account number %s not found!\n", accNum)
			}
		case 2:
			check := false
			f.Print("Enter account number to withdraw: ")
			var accNum string
			f.Scan(&accNum)
			for i, acc := range accounts {
				if acc.accountNumber == accNum {
					var withdrawAmount float64
					f.Print("Enter amount to withdraw: ")
					f.Scan(&withdrawAmount)
					account.WithdrawFromAccount(&accounts, withdrawAmount, i)
					check = true
					break

				}
			}
			if !check {
				f.Printf("Account number %s not found!\n", accNum)
			}
		case 3:
			f.Println("❤️ List of Bank Accounts:")
			for _, acc := range accounts {
				f.Printf("Account Number: %s, Balance: %.2f\n", acc.accountNumber, acc.balance)
			}
		case 0:
			f.Println("Exiting Banking System.")
			return
		default:
			f.Println("Invalid choice! Please try again.")
		}
	}
}
