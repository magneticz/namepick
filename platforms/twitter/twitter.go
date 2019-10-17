package twitter

import (
	"encoding/json"
	"fmt"
	"sync"

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
func (t twitter) Check(username string, c chan platforms.CheckResult, w *sync.WaitGroup) {
	defer w.Done()
	result := platforms.CheckResult{Name: "twitter", Value: false}
	data := twitterResponse{}
	url := fmt.Sprintf(rootURL, username)
	body, err := platforms.GetData(url, nil)
	if err != nil {
		// handle error
		fmt.Println("Error:>", err)
	}

	json.Unmarshal(body, &data)
	result.Value = data.Valid
	c <- result
}
