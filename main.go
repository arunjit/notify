package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"strings"

	proto "github.com/golang/protobuf/proto"
)

const (
	defaultGCMURL      = "https://gcm-http.googleapis.com/gcm/send"
	defaultTopicURL    = "/topic/global"
	defaultMessageFmt  = "Command (@cmd@) exited with status (@status@)"
	defaultFailureOnly = true
)

var (
	configFile           = flag.String("config", "~/.config/notify.pb", "Path to config file")
	sendTestNotification = flag.Bool("send_test_notification", false, "Send a test notification")
)

func readConfig() (*Config, error) {
	path := *configFile
	if strings.HasPrefix(path, "~") {
		path = strings.Replace(path, "~", os.Getenv("HOME"), 1)
	}
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	config := new(Config)
	err = proto.UnmarshalText(string(data), config)
	return config, err
}

func main() {
	flag.Parse()
	log.SetPrefix("notify:")
	config, err := readConfig()
	if err != nil {
		log.Fatalf("Error reading config: %v\n", err)
	}

	sender := NewSender(config.GetApiKey(), config.GetGcmUrl())
	notification := NewNotification(config.GetTopicUrl(), config.GetMessageFmt())

	if *sendTestNotification {
		notification.SetData(NewCommandStatus("test", 0))
		err := sender.Send(notification)
		if err != nil {
			log.Fatalf("Unable to send notification: %v\n", err)
		}
		return
	}
}
