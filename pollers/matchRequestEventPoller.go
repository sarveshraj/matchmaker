package pollers

import (
	"github.com/sarveshraj/matchmaker/model"
	"github.com/sarveshraj/matchmaker/processors"
)

// StartPolling starts polling for events
func StartPolling() {
	for {
		// TODO: logic to poll the queue
		// store the message in some value
		var message model.Message

		// process event
		go func() {
			processors.Process(message)
		}()
	}
}
