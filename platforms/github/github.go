package github

import (
	"encoding/json"
	"fmt"
	"sync"

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
func (g github) Check(username string, c chan platforms.CheckResult, w *sync.WaitGroup) {
	defer w.Done()

	result := platforms.CheckResult{Name: "github", Value: false}
	data := githubResponse{}
	url := fmt.Sprintf(rootURL, username)
	body, err := platforms.GetData(url, nil)
	if err != nil {
		// handle error
		fmt.Println("Error:>", err)
	}
	json.Unmarshal(body, &data)

	result.Value = data.Login == ""

	c <- result

}
