package main

import (
	"fmt"
	"net/url"
)

func main() {
	const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&depth=hooks"
	fmt.Println("	welcome to handling urls in go lang")
	// parsing
	result, _ := url.Parse(myurl)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)
	qparams := result.Query()
	fmt.Println("The type of query params are%T\n", qparams)
	fmt.Print(qparams["coursename"], "\n")
	for _, val := range qparams {
		fmt.Println(val)
	}

	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=Soham",
	}
	creatingUrlFromPartsOfUrl := partsOfUrl.String()
	fmt.Println(creatingUrlFromPartsOfUrl)
}
