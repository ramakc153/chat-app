package database

import (
	"fmt"
	"time"
)

type DBMessage struct {
	Id         string    `json:"id"`
	SenderId   string    `json:"sender_id"`
	ReceiverId string    `json:"receiver_id"`
	Content    string    `json:"content"`
	Timestamp  time.Time `json:"timestamp"`
	Status     string    `json:"status"`
}

func InsertMessage(data DBMessage) error {
	_, err := DB.Exec("INSERT INTO messages (id, sender_id, receiver_id, message_content, timestamp, status) VALUES (gen_random_uuid(), $1, $2, $3, NOW(), $4)",
		data.SenderId, data.ReceiverId, data.Content, data.Status)
	if err != nil {
		return fmt.Errorf("users duplicate")
	}
	return nil
}
