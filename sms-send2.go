package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	url := "https://panel.asanak.com/webservice/v1rest/sendsms"
	str := "username=09xxxxxxx&password=xxxx&" +
		"source=98xxxxx&destination=98xxxxxx&message=سلام"
	payload := strings.NewReader(str)

	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	req.Header.Add("content-type", "application/x-www-form-urlencoded")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer res.Body.Close()

	// Replacing ioutil.ReadAll with io.ReadAll
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	fmt.Println(res)
	fmt.Println(string(body))
}
