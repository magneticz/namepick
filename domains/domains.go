package domains

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
	"sync"

	"github.com/domainr/whois"
	"github.com/magneticz/namepick/platforms"
)

type domains []struct {
	Domain   string `json:"domain"`
	NotFound string `json:"notFound"`
}

var domainsMap = make(map[string]string)

// Domains is a list of supported domains
var Domains = []string{"com", "dev", "io", "me", "ai", "co"}

func init() {
	getDomainsConfig()
}

// Check domain name availability
func Check(uname string, tld string, c chan platforms.CheckResult, w *sync.WaitGroup) {
	defer w.Done()

	domain := uname + "." + tld
	result := platforms.CheckResult{Name: tld, Value: false}
	req, err := whois.NewRequest(domain)
	if err != nil {
		fmt.Println("Error:>", err)
	}

	resp, err := whois.DefaultClient.Fetch(req)
	if err != nil {
		fmt.Println("Error:>", err)
	}

	t := fmt.Sprintf("%s", resp)
	result.Value = strings.Contains(t, domainsMap[tld])

	c <- result
}

func getDomainsConfig() {
	d, err := ioutil.ReadFile("domains.json")
	var domainsArray = domains{}

	if err != nil {
		fmt.Println(err)
	}

	json.Unmarshal(d, &domainsArray)
	for _, j := range domainsArray {
		domainsMap[j.Domain] = j.NotFound
	}
}
