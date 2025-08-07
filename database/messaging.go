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
	_, err := DB.Exec("INSERT INTO messages (sender_id, receiver_id, message_content, timestamp, status) VALUES ($1, $2, $3, NOW(), $4)",
		data.SenderId, data.ReceiverId, data.Content, data.Status)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	return nil
}

func GetPendingMessage(receiver_id string) ([]DBMessage, error) {
	var QueriedMessage []DBMessage
	rows, err := DB.Query("SELECT * FROM messages WHERE receiver_id=$1 AND status='pending'", receiver_id)

	if err != nil {
		return nil, fmt.Errorf("error when trying to get message: %s", err.Error())
	}

	defer rows.Close()

	for rows.Next() {
		eachMessage := DBMessage{}
		err = rows.Scan(&eachMessage.Id, &eachMessage.SenderId, &eachMessage.ReceiverId, &eachMessage.Content, &eachMessage.Timestamp, &eachMessage.Status)
		if err != nil {
			return nil, fmt.Errorf(err.Error())
		}
		QueriedMessage = append(QueriedMessage, eachMessage)
	}
	if len(QueriedMessage) == 0 {
		return nil, fmt.Errorf("no messages found")
	}
	return QueriedMessage, nil
}

func UpdateMessageStatus(message_id string) error {
	_, err := DB.Exec("UPDATE messages SET status='delivered' WHERE id=$1", message_id)
	if err != nil {
		return fmt.Errorf("error when updating message status: %s", err.Error())
	}
	return nil
}

func GetMessageByUser(sender_id, receiver_id string) ([]DBMessage, error) {
	var messages []DBMessage
	rows, err := DB.Query("SELECT * FROM messages WHERE (sender_id=$1 AND receiver_id=$2) OR (sender_id=$2 AND receiver_id=$1 ORDER BY timestamp ASC)", sender_id, receiver_id)
	if err != nil {
		return nil, fmt.Errorf("error when get message by user: %s", err.Error())
	}
	for rows.Next() {
		var message DBMessage

		err = rows.Scan(&message.Id, &message.SenderId, &message.ReceiverId, &message.Content, &message.Timestamp, &message.Status)
		if err != nil {
			return nil, fmt.Errorf("error when scanning message: %s", err.Error())
		}
		messages = append(messages, message)
	}

	if len(messages) == 0 {
		return messages, nil
	}
	return messages, nil

}
