package reddit

import (
	"fmt"
	"sync"

	"github.com/magneticz/namepick/platforms"
)

// reddit is a struct
type reddit struct{}

func init() {
	assignToMap(&reddit{})
}

func assignToMap(p platforms.Platform) {
	platforms.Platforms["reddit"] = p
}

var rootURL = "https://www.reddit.com/api/username_available.json?user=%s"

// Check checks if username exists in twitters
func (r *reddit) Check(username string, c chan platforms.CheckResult, w *sync.WaitGroup) {

	defer w.Done()
	result := platforms.CheckResult{Name: "reddit", Value: false}
	url := fmt.Sprintf(rootURL, username)
	headers := map[string]string{
		"user-agent": "Chrome/77.0.3865.90",
	}
	body, err := platforms.GetData(url, headers)
	if err != nil {
		// handle error
		fmt.Println("Error:>", err)
	}
	result.Value = string(body) == "true"
	c <- result

}
