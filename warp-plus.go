package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
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
	for true {
		makeWarpReq(refferer)
		time.Sleep(35 * time.Second)
	}
}

func makeWarpReq(refferer string) {
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
	req.Header.Set("Host", "api.cloudflareclient.com")
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Accept-Encoding", "gzip")
	req.Header.Set("User-Agent", "okhttp/3.12.1")

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	client := &http.Client{Transport: tr}
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	fmt.Println(resp.Status)
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
