package main

import "fmt"

type Sender interface {
	Send(msg string) error
}
type Email struct {
	Address string
}

func (e *Email) Send(msg string) error {
	fmt.Printf("Сообщение \"%v\" отправлено по почте на адрес %v \n", msg, e.Address)
	return nil
}

type Phone struct {
	Number  int
	Balance int
}

func (p *Phone) Send(msg string) error {
	fmt.Printf("Сообщение \"%v\" отправлено по телефону с номера %v \n", msg, p.Number)
	return nil
}

func Notify(s Sender) {
	err := s.Send("Notify message")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Success")
}
func main() {
	email := &Email{"zufirka@icloud.com"}
	Notify(email)
	phone := &Phone{Number: 88005553535, Balance: 150}
	Notify(phone)
}
