package twitter

import (
	"encoding/json"
	"fmt"

	"github.com/magneticz/namepick/platforms"
)

// twitter is a struct
type twitter struct{}

type twitterResponse struct {
	Valid  bool   `json:"valid"`
	Reason string `json:"reason"`
	Msg    string `json:"msg"`
	Desc   string `json:"desc"`
}

var rootURL = "https://twitter.com/users/username_available?username=%s"

func init() {
	assignToMap(&twitter{})
}

func assignToMap(t platforms.Platform) {
	platforms.Platforms["twitter"] = t
}

// Check checks if username exists in twitters
func (t twitter) Check(username string) (bool, error) {
	data := twitterResponse{}
	url := fmt.Sprintf(rootURL, username)
	body, err := platforms.GetData(url, nil)
	if err != nil {
		// handle error
		fmt.Println("Error:>", err)
		return false, err
	}

	json.Unmarshal(body, &data)

	return data.Valid, nil
}
