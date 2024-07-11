package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"mime"
	"net/http"
	"sort"
)

//ใช้สำหรับ tiktok
func ExamGenerateSHA() string {
	// queries := "http://open-api.tiktokglobalshop.com/event/202309/webhooks?app_key=68xu9ks5p4i8&timestamp=1696909648&sign=ca75e39eb5f778005c7ec17a816830ec733f17c95317d4321976e2587e9e2424&shop_cipher=ROW_xkMbgAAAeVAQra0eZWebFQq5aIK"
	// a := CalSign()
	// fmt.Println(a)

	// var token = "ROW_54uEiQAwZbK9tV0zsbKWR-_dmt8sBnE_lwDk-RhKPYT7evVyeXRM0dqQT7-p29B-2IJZTvhpgG2s0njn86HEkxMcqf2CZOe176jl80KdQEjBs9M1K8cGGPLA7a625eV8Y5V9eXA"
	var url = "/event/202309/webhooksapp_key=68xu9ks5p4i8timestamp=1696909648shop_cipher=ROW_xkMbgAAAeVAQra0eZWebFQq5aIK"
	var path = url

	var queries = map[string]string{
		"address":    "https://partner.tiktokshop.com",
		"event_type": "PACKAGE_UPDATE",
	}
	var secret = "e59af819cc"

	keys := make([]string, len(queries))
	idx := 0
	for k, _ := range queries {
		keys[idx] = k
		idx++
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})
	input := path
	for _, key := range keys {
		input = input + key + queries[key]
	}
	input = secret + `/event/202309/webhooksapp_key68xu9ks5p4i8shop_cipherROW_xkMbgAAAeVAQra0eZWebFQq5aIKtimestamp1696909648{"address":"https://partner.tiktokshop.com","event_type":"PACKAGE_UPDATE"}` + secret
	// fmt.Println("input : ", input)

	a := generateSHA256(input, secret)
	fmt.Println(a)
	return a

	// h := hmac.New(sha256.New, []byte(secret))
	// if _, err := h.Write([]byte(input)); err != nil {
	// 	// todo: log error
	// 	return ""
	// }

	// fmt.Println(hex.EncodeToString(h.Sum(nil)))
	// return hex.EncodeToString(h.Sum(nil))
}

func CalSign(req *http.Request) string {
	queries := req.URL.Query()
	secret := "e59af819cc"

	// extract all query parameters excluding sign and access_token
	keys := make([]string, len(queries))
	idx := 0
	for k := range queries {
		// params except 'sign' & 'access_token'
		if k != "sign" && k != "access_token" {
			keys[idx] = k
			idx++
		}
	}

	// reorder the parameters' key in alphabetical order
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})

	// Concatenate all the parameters in the format of {key}{value}
	input := ""
	for _, key := range keys {
		input = input + key + queries.Get(key)
	}

	// append the request path
	input = req.URL.Path + input

	// if the request header Content-type is not multipart/form-data, append body to the end
	mediaType, _, _ := mime.ParseMediaType(req.Header.Get("Content-type"))
	if mediaType != "multipart/form-data" {
		body, _ := io.ReadAll(req.Body)
		input = input + string(body)

		req.Body.Close()
		// reset body after reading from the original
		req.Body = io.NopCloser(bytes.NewReader(body))
	}

	// wrap the string generated in step 5 with the App secret
	input = secret + input + secret

	return generateSHA256(input, secret)
}

func generateSHA256(input, secret string) string {
	// encode the digest byte stream in hexadecimal and use sha256 to generate sign with salt(secret)
	h := hmac.New(sha256.New, []byte(secret))

	if _, err := h.Write([]byte(input)); err != nil {
		// TODO error log
		return ""
	}

	return hex.EncodeToString(h.Sum(nil))
}

func XXXXXX() string {

	var token = "ROW_5BEy3gAAAAB9lI0bdaI-iJcamaVF4_yNI6_5Z8pRhfyoGL5fS-bPCItPwsco14LL8DAdugwxVXjU1gawCBimWff-VwbAOcjSFOSpZpWGwtpH_LhcQgXMJBi2pin_CRK2uuF-YEpkZnbblMCXTMBNrEn-f91ddp-MG7Poojy7zpM1PnxROudAvQ"
	var url = "https://open-api.tiktokglobalshop.com/api/products/search?"
	url += "access_token=" + token
	url += "&app_key=6c7alu3q7j6ut"
	url += "&shop_cipher="
	url += "&shop_id="
	url += "&timestamp=1713416480"
	url += "&version=202212"
	var path = url

	var queries = map[string]string{
		"q": "hello",
	}
	var secret = "f886d52543752603fb486f57a4a924e94bff8b91"

	keys := make([]string, len(queries))
	idx := 0
	for k, _ := range queries {
		keys[idx] = k
		idx++
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})
	input := path
	for _, key := range keys {
		input = input + key + queries[key]
	}
	input = secret + input + secret

	h := hmac.New(sha256.New, []byte(secret))
	if _, err := h.Write([]byte(input)); err != nil {
		// todo: log error
		return ""
	}

	fmt.Println(hex.EncodeToString(h.Sum(nil)))
	return hex.EncodeToString(h.Sum(nil))
}
