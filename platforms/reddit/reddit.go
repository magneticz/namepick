package reddit

import (
	"fmt"

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
func (r *reddit) Check(username string) (bool, error) {
	url := fmt.Sprintf(rootURL, username)
	headers := map[string]string{
		"user-agent": "Chrome/77.0.3865.90",
	}
	body, err := platforms.GetData(url, headers)
	if err != nil {
		// handle error
		fmt.Println("Error:>", err)
		return false, err
	}
	return string(body) == "true", nil
}
