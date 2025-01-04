package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	url := "https://panel.asanak.com/webservice/v1rest/sendsms"
	str := "username=989xxxxx&password=xxxxx&" +
		"source=9821xxxxxxx&destination=xxxxxxx&message=hello"
	payload := strings.NewReader(str)

	// ایجاد درخواست HTTP و بررسی ارور
	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		fmt.Println("خطا در ایجاد درخواست:", err)
		return
	}

	req.Header.Add("content-type", "application/x-www-form-urlencoded")

	// ارسال درخواست و بررسی ارور
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("خطا در ارسال درخواست:", err)
		return
	}
	defer res.Body.Close()

	// خواندن بدنه پاسخ و بررسی ارور
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("خطا در خواندن بدنه پاسخ:", err)
		return
	}

	fmt.Println("وضعیت پاسخ:", res.Status)
	fmt.Println("بدنه پاسخ:", string(body))
}
