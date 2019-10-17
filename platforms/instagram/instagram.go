package instagram

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/magneticz/namepick/platforms"
)

type instagram struct{}

type instagramResponse struct {
	Errors struct {
		Username []struct {
			Message string `json:"message"`
			Code    string `json:"code"`
		} `json:"username"`
	} `json:"errors"`
}

var rootURL = "https://www.instagram.com/web/__mid/"
var attemptURL = "https://www.instagram.com/accounts/web_create_ajax/attempt/"

func init() {
	assignToMap(&instagram{})
}

func assignToMap(p platforms.Platform) {
	platforms.Platforms["instagram"] = p
}

// Check checks if username exists in twitters
func (i instagram) Check(username string) (bool, error) {
	data := instagramResponse{}
	body, err := performCheckRequests(username)
	if err != nil {
		// handle error
		fmt.Println("Error:>", err)
		return false, err
	}

	json.Unmarshal(body, &data)
	if len(data.Errors.Username) > 0 && data.Errors.Username[0].Code != "username_is_taken" {
		return false, nil
	}

	return true, nil
}

func performCheckRequests(username string) ([]byte, error) {
	var token string
	headers := map[string]string{
		"cookie": "ig_cb=1",
	}
	tokenResp, err := getData("GET", rootURL, headers, nil)
	for _, cookie := range tokenResp.Header {
		if strings.Contains(cookie[0], "csrftoken") {
			cookies := strings.Split(cookie[0], ";")
			token = strings.Split(cookies[0], "=")[1]
		}
	}
	headers = map[string]string{
		"x-csrftoken":  token,
		"Content-type": "application/x-www-form-urlencoded",
	}
	body := url.Values{}
	body.Set("username", username)
	result, err := getData("POST", attemptURL, headers, strings.NewReader(body.Encode()))
	if err != nil {
		// handle error
		fmt.Println("Error:>", err)
		return nil, err
	}
	defer result.Body.Close()
	defer tokenResp.Body.Close()
	return ioutil.ReadAll(result.Body)
}

func getData(method string, url string, headers map[string]string, body io.Reader) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		// handle error
		fmt.Println("Error:>", err)
		return nil, err
	}
	for k, v := range headers {
		req.Header.Add(k, v)

	}
	resp, err := client.Do(req)
	if err != nil {
		// handle error
		fmt.Println("Error:>", err)
		return nil, err
	}
	return resp, nil
}
