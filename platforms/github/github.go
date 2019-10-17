package github

import (
	"encoding/json"
	"fmt"

	"github.com/magneticz/namepick/platforms"
)

type github struct{}

type githubResponse struct {
	Login string `json:"login"`
}

var rootURL = "https://api.github.com/users/%s"

func init() {
	assignToMap(&github{})
}

func assignToMap(p platforms.Platform) {
	platforms.Platforms["github"] = p
}

// Check returns a pointer to a new github instance
func (g github) Check(username string) (bool, error) {
	data := githubResponse{}
	url := fmt.Sprintf(rootURL, username)
	body, err := platforms.GetData(url, nil)
	if err != nil {
		// handle error
		fmt.Println("Error:>", err)
		return false, err
	}
	json.Unmarshal(body, &data)

	if data.Login != "" {
		return false, nil
	}

	return true, nil
}
