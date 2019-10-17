package platforms

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

//Platforms is a var
var Platforms = make(map[string]Platform)

// CheckResult is a struct for results of check
type CheckResult struct {
	Name  string
	Value bool
}

// Platform is an interface
type Platform interface {
	Check(string, chan CheckResult, *sync.WaitGroup)
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
