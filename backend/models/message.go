package models

import (
	"time"
)

type Message struct {
	ID        string
    From      string
    RoomID    string
    Body      string
    Timestamp time.Time
}

// Create
func NewMessage(from, roomID, body string) Message {
    return Message {
        From:      from,
        RoomID:    roomID,
        Body:      body,
        Timestamp: time.Now(),
    }
}

var offlineMessages = make(map[string][]Message)

// Create
func SaveOfflineMessage(clientID string, msg Message) {
    offlineMessages[clientID] = append(offlineMessages[clientID], msg)
}

// Read
func GetOfflineMessages(clientID string) []Message {
    msgs := offlineMessages[clientID]
    delete(offlineMessages, clientID) // flush after delivery
    return msgs
}

// Update
func UpdateOfflineMessage(clientID string, index int, newMsg Message) {
    if msgs, ok := offlineMessages[clientID]; ok && index < len(msgs) {
        msgs[index] = newMsg
        offlineMessages[clientID] = msgs
    }
}

// Delete
func DeleteOfflineMessages(clientID string) {
    delete(offlineMessages, clientID)
}
