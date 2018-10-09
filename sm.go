package main

import (
	"log"
	"net/smtp"
)

func main() {
	send("Hello 幹嘛傳給我拉")
}

func send(body string) {
	from := "jonahjun.wu@gmail.com"
	pass := "2xuiuiji"
	to := "junmein@hotmail.com"
	/*
		msg := "From: " + from + "\n" +
			"To: " + to + "\n" +
			"Subject: Hello there\n\n" +
			body
	*/
	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: 幹嘛拉\n\n" +
		body

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return
	}

	log.Print("sent, visit junmein@hotmail.com")
}
