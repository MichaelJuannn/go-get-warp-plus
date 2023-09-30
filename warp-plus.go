package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/http/httputil"
	"time"
)

type BodyType struct {
	key          string
	install_id   string
	fcm_token    string
	referrer     string
	warp_enabled bool
	tos          string
}

func main() {
	var refferer string
	fmt.Print("Insert Your Client Id")
	fmt.Scan(&refferer)
	url := fmt.Sprintf("https://api.cloudflareclient.com/v0a%s/reg", digitString(3))
	install_id := genString(22)
	body := map[string]any{
		"key":          fmt.Sprintf("%s=", genString(43)),
		"install_id":   install_id,
		"fcm_token":    fmt.Sprintf("%s:APA91b%s", install_id, genString(134)),
		"referrer":     refferer,
		"warp_enabled": false,
		"tos":          fmt.Sprintf("%s+02:00", time.Now().Format("2006-01-02T15:04:05.071")), // iso date
		"type":         "Android",
		"locale":       "es_ES",
	}
	reqBody, err := json.Marshal(body)
	if err != nil {
		panic(err)
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqBody))
	if err != nil {
		panic(err)
	}
	// req.Header.Set("Host", "api.cloudflareclient.com")
	req.Header.Add("Content-Type", "application/json; charset=UTF-8")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Accept-Encoding", "gzip")
	req.Header.Add("User-Agent", "okhttp/3.12.1")
	req.Header.Add("Accept", "*/*")

	reqDump, err := httputil.DumpRequestOut(req, true)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("REQUEST:\n%s \n", string(reqDump))

	// client := &http.Client{}
	// resp, err := client.Do(req)

	// if err != nil {
	// 	panic(err)
	// }
	// defer resp.Body.Close()

	// fmt.Println(resp.Status)

	// respDump, err := httputil.DumpResponse(resp, true)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Printf("RESPONSE:\n%s", string(respDump))

}

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func genString(length int) string {
	letter := asciiLetter() + "0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = letter[seededRand.Intn(len(letter))]
	}
	return string(b)
}

func digitString(length int) string {
	digits := "0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = digits[seededRand.Intn(len(digits))]
	}
	return string(b)
}

func asciiLetter() string {
	var character string
	for i := 65; i <= 122; i++ {
		if 91 <= i && i <= 96 {
			continue
		}
		character += string(rune(i))
	}
	return character
}
