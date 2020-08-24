package dtos

// MatchRequestEvent will be the match making event published by games
type MatchRequestEvent struct {
	userID string
	gameID string
}

// GetUserID is getter for string userID
func (e *MatchRequestEvent) GetUserID() string {
	return e.userID
}

// GetGameID is getter for string gameID
func (e *MatchRequestEvent) GetGameID() string {
	return e.gameID
}