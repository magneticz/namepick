package main

import (
	"fmt"

	"github.com/magneticz/namepick/domains"
	p "github.com/magneticz/namepick/platforms"
	_ "github.com/magneticz/namepick/platforms/github"
	_ "github.com/magneticz/namepick/platforms/instagram"
	_ "github.com/magneticz/namepick/platforms/reddit"
	_ "github.com/magneticz/namepick/platforms/twitter"
)

func main() {
	name := "magnetic"

	fmt.Println("PLATFORMS")

	for key, platforms := range p.Platforms {
		result, err := platforms.Check(name)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Printf("%s:> %v\n", key, result)
	}

	fmt.Println("DOMAINS")

	for _, domain := range domains.Domains {
		fmt.Printf("%s:> %v\n", domain, domains.Check(name, domain))
	}
}
