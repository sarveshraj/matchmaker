package processors

import (
	"encoding/json"
	"math"
	"time"

	"github.com/sarveshraj/matchmaker/connectors"
	"github.com/sarveshraj/matchmaker/dtos"
	"github.com/sarveshraj/matchmaker/model"
)

const (
	online  int = 0
	playing     = 1
	waiting     = 2
)

const (
	gameRating        string = "GAME_RATING"
	similarGameRating string = "SIMILAR_GAME_RATING"
	overallRating     string = "OVERALL_RATING"
)

// Process consumes message of type MatchRequestEvent
func Process(m model.Message) error {
	messageValue := m.GetValue()
	if messageValue == "" {
		// throw error
	}

	var event dtos.MatchRequestEvent
	json.Unmarshal([]byte(messageValue), &event)

	timestamp := m.GetTimestamp()

	// call user service and check user status
	userStatus := connectors.GetUserStatus(event.GetUserID())

	// if user is not waiting for opponents return
	if userStatus.Status != waiting || userStatus.GameID != event.GetGameID() {
		// throw err
	}

	// call games service and get min number of players required to start game
	minPlayers := connectors.GetMinPlayersRequired(event.GetGameID())

	// call ratings service and get ratings of user
	userRating := connectors.GetRatingOfUser(event.GetUserID())

	// calculate all thresholds of user
	userThresholds := calculateThresholds(userRating, timestamp)

	// dominated users list
	var userIDsInDominance []string

	// call user service and get list of all online users which have game rating in range +- game threshold
	var minThreshold float64 = 0
	if userRating.GameRating-userThresholds.GameThreshold > 0 {
		minThreshold = userRating.GameRating - userThresholds.GameThreshold
	}

	matchCandidates := connectors.GetAllOnlineUsers(event.GetGameID(), gameRating, minThreshold, userRating.GameRating+userThresholds.GameThreshold)

	// add the list to dominating users list
	userIDsInDominance = append(userIDsInDominance, matchCandidates...)

	// call user service and get list of all online users which have similar game rating in range +- similar game threshold
	minThreshold = 0
	if userRating.SimilarGameRating-userThresholds.SimilarGameThreshold > 0 {
		minThreshold = userRating.SimilarGameRating-userThresholds.SimilarGameThreshold
	}

	matchCandidates = connectors.GetAllOnlineUsers(event.GetGameID(), similarGameRating, minThreshold, userRating.SimilarGameRating+userThresholds.SimilarGameThreshold)

	// add the list to dominating users list
	userIDsInDominance = append(userIDsInDominance, matchCandidates...)

	// call user service and get list of all online users which have overall rating in range +- overall threshold
	minThreshold = 0
	if userRating.OverallRating-userThresholds.OverallThreshold > 0 {
		minThreshold = userRating.OverallRating-userThresholds.OverallThreshold
	}

	matchCandidates = connectors.GetAllOnlineUsers(event.GetGameID(), overallRating, minThreshold, userRating.OverallRating+userThresholds.OverallThreshold)
	// add the list to dominating users list
	userIDsInDominance = append(userIDsInDominance, matchCandidates...)

	// if min num of players required to start game < list of dominated users
	if minPlayers < len(userIDsInDominance) {
		// TODO: post to retry topic
		return nil
	}

	// call game webhook with the userId and list of dominated users
	connectors.NotifyGame(event.GetUserID(), event.GetGameID(), userIDsInDominance)
	
	return nil
}

func calculateThresholds(userRating dtos.UserRating, timestamp int64) dtos.Thresholds {
	var thresholds dtos.Thresholds

	thresholds.OverallThreshold = exponentiallyScaleThreshold(userRating.OverallRating, timestamp)
	thresholds.GameThreshold = exponentiallyScaleThreshold(userRating.GameRating, timestamp)
	thresholds.SimilarGameThreshold = exponentiallyScaleThreshold(userRating.SimilarGameRating, timestamp)

	return thresholds
}

func exponentiallyScaleThreshold(rating float64, timestamp int64) float64 {
	currentTimeInMS := time.Now().UnixNano() / (int64(time.Millisecond) / int64(time.Nanosecond))
	growthFactor := float64(currentTimeInMS-timestamp) / 30000
	return math.Round(100 * math.Pow(rating/100, growthFactor))
}
