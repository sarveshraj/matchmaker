package model

// Message to be processed
type Message struct {
	value string // json string of the Message
	timestamp int64
}

// GetValue is getter for string value
func (e *Message) GetValue() string {
	return e.value
}

// GetTimestamp is getter for uint timestamp
func (e *Message) GetTimestamp() int64 {
	return e.timestamp
}