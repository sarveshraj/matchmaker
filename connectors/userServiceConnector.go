package connectors

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/sarveshraj/matchmaker/dtos"
)

const (
	usersBaseURL    string = "https://api.getmega.com/users"
	status                 = "/status"
	ratingTypeKey          = "rating_type="
	minThresholdKey        = "min_threshold="
	maxThresholdKey        = "max_threshold="
)

// GetUserStatus gets user status
func GetUserStatus(userID string) dtos.UserStatus {
	url := usersBaseURL
	url += userID
	url += status

	resp, err := http.Get(url)
	if err != nil {
		// throw err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// throw err
	}

	var userStatus dtos.UserStatus
	json.Unmarshal([]byte(body), &userStatus)

	return userStatus
}

// GetAllOnlineUsers gets all online users
func GetAllOnlineUsers(gameID string, ratingType string, min float64, max float64) []string {
	url := usersBaseURL
	url += "?"
	url += ratingTypeKey
	url += ratingType
	url += "&"
	url += minThresholdKey
	url += strconv.FormatUint(uint64(min), 10)
	url += "&"
	url += maxThresholdKey
	url += strconv.FormatUint(uint64(max), 10)

	resp, err := http.Get(url)
	if err != nil {
		// throw err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// throw err
	}

	var dominatedUserIds []string
	json.Unmarshal([]byte(body), &dominatedUserIds)

	return dominatedUserIds
}
