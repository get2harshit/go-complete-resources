package main

import (
	"errors"
	"fmt"
	"structPointers/utils"
	"time"
)

type UserId int
type ProductId int64
type Balance float64

type User struct {
	firstName string
	lastName  string
	createdAt time.Time
}

type BankAccount struct {
	accNumber string
	accountHolder string
	Balance float64
}

// struct embedding
type AccountHolder struct{
	user User
	bankAccount BankAccount
}

var myBankAccount BankAccount = BankAccount{"123", "surya", 123}

func (u User) printUserDetails() {
	fmt.Println(u.firstName, u.lastName, u.createdAt)
}

func (u *User) clearUserDetails(){
	u.firstName = ""
	u.lastName = ""
}

func (u *User) updateUserDetails(firstName, lastName string){
	u.firstName = firstName
	u.lastName = lastName
}
// constructor function
// no receiver attached
func NewUser(firstName, lastName string) (*User, error){

	if firstName == "" {
		return nil, errors.New("FirstName and Last name is required")
	}

	// success
	return &User{
		firstName: firstName,
		lastName: lastName,
		createdAt: time.Now(),
	}, nil
}
// Task: create function updateUserDetails
// eg: Rohit -> Mr. Rohit
// Modi -> Kumar Modi
func main() {
	fmt.Println("Learning Structs and Custom types")
	var u UserId = 10
	fmt.Println(u)

	// firstName:= "a"
	// lastName:="b"

	var user User = User{
		"c",
		"d",
		time.Now(),
	}

	fmt.Println(user.firstName, user.lastName, time.Now().Format("15:04:06"))

	polaris := utils.Location{Lat: 2.0, Long: 2.0}
	fmt.Println(polaris)

	hostel := utils.Location{Lat: 2, Long: 3}
	fmt.Println(hostel)

	turf := utils.Location{}
	fmt.Println(turf)
	

	// Today's class
	rohit:= User{"rohit", "modi", time.Now()}
	// printUserDetails(rohit)
	rohit.printUserDetails()
	rohit.clearUserDetails()
	rohit.printUserDetails()


	var age int = 42
	var pointer *int
	pointer = &age
	fmt.Println(age)
	fmt.Println(pointer)
	fmt.Println(*pointer)
	*pointer = 20
	fmt.Println(age)


	// pointer to struct

	rohit.updateUserDetails("Mr. Rohit", "Kumar Modi")
	rohit.printUserDetails()

	virat, err := NewUser("Virat", "Kohli")

	if err!= nil {
		fmt.Println(err)
		return 
	}


	fmt.Println(virat)
	// virat:= CreateUser("Virat", "Kohli")


	surya, err:= NewUser("surya", "kumar")

	if err!= nil{
		fmt.Println(err)
		return
	}

	fmt.Println(surya)
	fmt.Printf("%p", surya)

	suryasAccount:= AccountHolder{*surya, myBankAccount}
	fmt.Println(suryasAccount)
	
	


}
