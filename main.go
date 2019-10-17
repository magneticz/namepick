package main

import (
	"fmt"
	"sync"

	"github.com/magneticz/namepick/domains"
	p "github.com/magneticz/namepick/platforms"

	_ "github.com/magneticz/namepick/platforms/github"
	_ "github.com/magneticz/namepick/platforms/instagram"
	_ "github.com/magneticz/namepick/platforms/reddit"
	_ "github.com/magneticz/namepick/platforms/twitter"
)

func main() {
	name := "magnetic"
	platformsChannel := make(chan p.CheckResult)
	var wg sync.WaitGroup

	fmt.Println("PLATFORMS")

	for _, platform := range p.Platforms {
		wg.Add(1)
		go platform.Check(name, platformsChannel, &wg)
	}

	fmt.Println("DOMAINS")

	for _, domain := range domains.Domains {
		wg.Add(1)
		go domains.Check(name, domain, platformsChannel, &wg)
	}

	for range p.Platforms {
		result := <-platformsChannel
		fmt.Printf("%s:> %v\n", result.Name, result.Value)
	}

	for range domains.Domains {
		result := <-platformsChannel
		fmt.Printf("%s:> %v\n", result.Name, result.Value)
	}
	wg.Wait()
}
