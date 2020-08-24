package processors

// every processor should implement this interface
type processor interface {
    Process(e *Event)
}
