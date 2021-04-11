package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	response, err := http.Get("https://api.coincap.io/v2/rates/bitcoin")
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
	responseData, err := ioutil.ReadAll(response.Body)
	fmt.Print(response)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
	fmt.Print(string(responseData))

}
