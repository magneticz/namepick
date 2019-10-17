package domains

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/domainr/whois"
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
func Check(uname string, tld string) (isAvailable bool) {
	domain := uname + "." + tld
	req, err := whois.NewRequest(domain)
	if err != nil {
		return
	}

	resp, err := whois.DefaultClient.Fetch(req)
	if err != nil {
		return
	}

	t := fmt.Sprintf("%s", resp)
	isAvailable = strings.Contains(t, domainsMap[tld])

	return isAvailable
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
