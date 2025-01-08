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
	fmt.Println("/nStatus:", response.Status)
	fmt.Println("/nStatus Code:", response.StatusCode)
	if response.StatusCode == 200 {
		os.Exit(1)
	}
	os.Exit(0)
}
