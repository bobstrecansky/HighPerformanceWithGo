package main

import "fmt"

type User struct {
	uid   int
	name  string
	email string
	phone string
}

func (u User) updateEmail(newEmail string) {
	u.email = newEmail
}

func (u *User) updatePhone(newPhone string) {
	u.phone = newPhone
}

func (u User) printEmail() {
	fmt.Println("User email: ", u.email)
}

func main() {

	userExample := User{
		uid:   1,
		name:  "bob",
		email: "bob@example.com",
		phone: "123-456-7890",
	}

	userExample.updateEmail("bob.strecansky@example.com")
	userExample.printEmail()
	(userExample).updatePhone("000-000-0000")
	fmt.Println("Updated User Email: ", userExample.email)
	fmt.Println("Updated User Phone: ", userExample.phone)

}
