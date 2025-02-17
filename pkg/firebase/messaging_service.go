package firebase

import (
	"context"
	"firebase.google.com/go/v4/messaging"
	"log"
)

type MessagingService struct{ *Messaging }

func NewMessagingService(messaging *Messaging) *MessagingService {
	return &MessagingService{messaging}
}

func (m *MessagingService) SendMessage(data map[string]string, token, title, body string) error {
	message := &messaging.Message{
		Token: token,
		Data:  data,
		Notification: &messaging.Notification{
			Title: title,
			Body:  body,
		},
	}

	response, err := m.Client.Send(context.Background(), message)
	if err != nil {
		return err
	}

	log.Printf("successful send push notification: %v", response)

	return nil
}
