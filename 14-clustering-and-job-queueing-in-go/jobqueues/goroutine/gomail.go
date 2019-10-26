package main

import (
	"log"
	"time"

	"gopkg.in/gomail.v2"
)

func main() {

	log.Printf("Doing Work")
	log.Printf("Sending Emails!")
	go sendMail()
	time.Sleep(time.Second)
	log.Printf("Done Sending Emails!")
}

func sendMail() {
	var sender = "USERNAME@gmail.com"
	var recipient = "RECIPIENT@gmail.com"
	var username = "USERNAME@gmail.com"
	var password = "PASSWORD"
	var host = "smtp.gmail.com"
	var port = 587

	email := gomail.NewMessage()
	email.SetHeader("From", sender)
	email.SetHeader("To", recipient)
	email.SetHeader("Subject", "Test Email From Goroutine")
	email.SetBody("text/plain", "This email is being sent from a Goroutine!")

	dialer := gomail.NewDialer(host, port, username, password)
	err := dialer.DialAndSend(email)
	if err != nil {
		log.Println("Could not send email")
		panic(err)
	}
}
