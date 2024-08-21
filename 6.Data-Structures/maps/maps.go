package main

import "fmt"

func main() {
	websites := map[string]string{}
	websites["LinkedIn"] = "https://linkedin.com"
	websites["Google"] = "https://google.com"
	websites["Github"] = "https://github.com"

	fmt.Println(websites["Google"])
	delete(websites, "Google")

}
