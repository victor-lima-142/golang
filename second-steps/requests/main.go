package requests

import (
	"fmt"
	"net/http"
	"os"
)

func GetRequest() {
	response, err := http.Get("https://www.github.com")

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println("\nStatus:", response.Status)
	fmt.Println("\nStatus Code:", response.StatusCode)
	if response.StatusCode == 200 {
		println("\nRequest was successful")
		return
	}
	println("\nRequest was not successful")
}

func GetRequestToUrl(url string) {
	response, err := http.Get(url)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println("\nStatus:", response.Status)
	fmt.Println("\nStatus Code:", response.StatusCode)
	if response.StatusCode == 200 {
		println("\nRequest was successful")
		return
	}
	println("\nRequest was not successful")
}

func RawGetRequestToUrl(url string) (*http.Response, error) {
	return http.Get(url)
}
