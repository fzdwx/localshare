package server

import "encoding/json"

type Message struct {
	Type   string `json:"type"`
	UserID string `json:"sender"`

	Text string `json:"text"`

	FileName    string `json:"fileName"`
	FileContent string `json:"fileContent"`
	FileType    string `json:"fileType"`

	raw []byte
}

const MessageTypeIdentify = "identify"
const MessageTypeText = "text"
const MessageTypeFile = "file"

func parseMessage(message []byte) (*Message, error) {
	var msg Message
	if err := json.Unmarshal(message, &msg); err != nil {
		return nil, err
	}

	msg.raw = message
	return &msg, nil
}
