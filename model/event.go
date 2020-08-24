package model

// Event to be processed
type Event struct {
	value string // json string of the event
	timestamp uint
}

// GetValue is getter for string value
func (e *Event) GetValue() string {
	return e.value
}

// GetTimestamp is getter for uint timestamp
func (e *Event) GetTimestamp() uint {
	return e.timestamp
}