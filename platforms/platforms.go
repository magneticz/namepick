package platforms

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

//Platforms is a var
var Platforms = make(map[string]Platform)

// Platform is an interface
type Platform interface {
	Check(string) (bool, error)
}

// GetData gets data
func GetData(url string, headers map[string]string) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
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
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
