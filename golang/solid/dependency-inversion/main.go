package main

import "fmt"

type MessageSender interface {
	SendMessage(msg string) error
}

type EmailSender struct{}

func (e EmailSender) SendMessage(msg string) error {
	fmt.Println("Sending email:", msg)
	return nil
}

type SMSsender struct{}

func (s SMSsender) SendMessage(msg string) error {
	fmt.Println("Sending SMS:", msg)
	return nil
}

type NotificationService struct {
	sender MessageSender
}

func (n NotificationService) Notify(msg string) error {
	return n.sender.SendMessage(msg)
}

func main() {
	emailSender := EmailSender{}
	notificationService := NotificationService{sender: emailSender}

	err := notificationService.Notify("Hello from Go!")
	if err != nil {
		fmt.Println("Notification failed:", err)
	}

	smsSender := SMSsender{}
	notificationService.sender = smsSender

	err = notificationService.Notify("Hello again from Go!")
	if err != nil {
		fmt.Println("Notification failed:", err)
	}
}
