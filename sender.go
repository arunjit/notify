package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"net/http"
	"net/http/httputil"
)

var (
	senderDebug = flag.Bool("sender_debug", false, "(for debugging) Print request and response to stdout")
)

type Sender struct {
	APIKey string
	GCMURL string
}

func NewSender(apiKey, gcmURL string) *Sender {
	return &Sender{apiKey, gcmURL}
}

func (s *Sender) Send(n *Notification) error {
	data, err := json.Marshal(n)
	if err != nil {
		return err
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", s.GCMURL, bytes.NewReader(data))
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "key="+s.APIKey)
	if *senderDebug {
		b, _ := httputil.DumpRequest(req, true)
		fmt.Println("---REQUEST---")
		fmt.Println(string(b))
	}

	res, err := client.Do(req)
	defer res.Body.Close()
	if err != nil {
		return err
	}
	if *senderDebug {
		b, _ := httputil.DumpResponse(res, true)
		fmt.Println("---RESPONSE---")
		fmt.Println(string(b))
	}
	if res.StatusCode >= 300 {
		return errors.New("Notification could not be sent. Status: " + res.Status)
	}
	fmt.Println("Notification sent")
	return nil
}
