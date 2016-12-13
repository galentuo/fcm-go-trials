package main

import (
	"fmt"
	"io/ioutil"
	"log"

	yaml "gopkg.in/yaml.v2"

	fcm "github.com/NaySoftware/go-fcm"
)

// Config something
type Config struct {
	ServerKey string   `yaml:"serverkey"`
	GcmIDs    []string `yaml:"gcmids"`
}

// pushData
type pushData struct {
	Title       string
	Image       string
	Description string
}

// SendAndroidPushNotification something
func SendAndroidPushNotification(serverKey string, gcmIDs []string) {
	pushMsg := pushData{
		Title:       "Title test",
		Image:       "testimg.img",
		Description: "test description",
	}

	client := fcm.NewFcmClient(serverKey)
	client.NewFcmRegIdsMsg(gcmIDs, pushMsg)
	status, err := client.Send()
	if err != nil {
		log.Printf(err.Error())
	}
	fmt.Println(status)
}

func main() {
	bytes, err := ioutil.ReadFile("config.yml")
	if err != nil {
		fmt.Println(err)
	}

	config := &Config{}
	if err := yaml.Unmarshal(bytes, config); err != nil {
		fmt.Println(err)
	}
	SendAndroidPushNotification(config.ServerKey, config.GcmIDs)
}
