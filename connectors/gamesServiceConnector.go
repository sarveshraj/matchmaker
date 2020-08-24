package connectors

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const (
	gamesBaseURL string = "https://api.getmega.com/games/"
)

// GetMinPlayersRequired fetches min num of user
func GetMinPlayersRequired(gameID string) int {
	url := gamesBaseURL
	url += gameID

	resp, err := http.Get(url)
	if err != nil {
		// throw err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// throw err
	}

	var minPlayers int
	json.Unmarshal([]byte(body), &minPlayers)

	return minPlayers
}

func NotifyGame(userID string, gameID string, userIDs []string) {
	url := gamesBaseURL
	url += gameID
	url += "/"
	url += userID

	// resp, err := http.Post(url, json.Marshal(userIDs))
	// if err != nil {
	// 	// throw err
	// }
}