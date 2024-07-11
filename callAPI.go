package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func ExamCallAPI() {

	url := "http://auth.tiktok-shops.com/api/v2/token/get?app_key=6c7alu3q7j6ut&auth_code=ROW_ML1_cgAAAACMK_DYeSq6Ny28z8DxmHf5JsppyoOiJxKa9rC5-XIZEZFbQJpTTAakuw82UGcB71_A_VGcjUX76oDJcJ2bE9xbd0x6FyRd_nkK77dJUYZrSjLincIKTsZWQy5ybQdwY6Pnwc0y6OuoJMYZCvSLwfDi1-nIqIm4hb0ejSlqPNmdCw&app_secret=f886d52543752603fb486f57a4a924e94bff8b91&grant_type=authorized_code"
	method := "POST"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
