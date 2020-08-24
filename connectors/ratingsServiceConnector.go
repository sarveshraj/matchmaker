package connectors

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/sarveshraj/matchmaker/dtos"
)

const (
	ratingsBaseURL string = "https://api.getmega.com/ratings/"
)

// GetRatingOfUser fetches ratings of user
func GetRatingOfUser(userID string) dtos.UserRating {
	url := ratingsBaseURL
	url += userID

	resp, err := http.Get(url)
	if err != nil {
		// throw err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// throw err
	}

	var userRating dtos.UserRating
	json.Unmarshal([]byte(body), &userRating)

	return userRating
}
