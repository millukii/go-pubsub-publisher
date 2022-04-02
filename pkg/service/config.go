package service

import (
	"encoding/json"
	"io/ioutil"
	"log"

)

type Configuration struct {
 Topics []Topic `json:"topics"`
	Subscriptions []Subscription `json:"subscriptions"`
	Subscriber []Subscriber `json:"subscribers"`
	Publishers []Publisher `json:"publishers"`
	Events []Event `json:"events"`
}
type	Publisher struct {
		ID    int    `json:"id"`
		Topic string `json:"topic"`
	} 
type	Subscription struct {
		ID           int    `json:"id"`
		Subscription string `json:"subscription"`
		Topic        string `json:"topic"`
		ProjectID    string `json:"projectId"`
	} 

type	Topic struct {
		ID        int    `json:"id"`
		Topic     string `json:"topic"`
		ProjectID string `json:"projectId"`
	}

type	Subscriber struct {
		ID                                    int    `json:"id"`
		Subscription                          string `json:"subscription"`
		ReceiveSettingsMaxExtension           int    `json:"receiveSettingsMaxExtension"`
		ReceiveSettingsMaxExtensionPeriod     int    `json:"receiveSettingsMaxExtensionPeriod"`
		ReceiveSettingsMaxOutstandingMessages int    `json:"receiveSettingsMaxOutstandingMessages"`
		ReceiveSettingsMaxOutstandingBytes    int    `json:"receiveSettingsMaxOutstandingBytes"`
		ReceiveSettingsNumGoroutines          int    `json:"receiveSettingsNumGoroutines"`
		ReceiveSettingsSynchronous            bool   `json:"receiveSettingsSynchronous"`
}
type	Attributes struct {
				EventName string `json:"eventName"`
				EventType string `json:"eventType"`
} 
type Output struct {
	Attributes Attributes `json:"attributes"`
}
type Filters struct {
	Attributes Attributes `json:"attributes"`
}

type Event struct{
	Filters *Filters  `json:"filters"`
	Output *Output  `json:"output"`
	Subscriber int  `json:"subscriber"`
	Publisher int  `json:"publisher"`
}
func  ReadConfig() (*Configuration,error) {
	file, err := ioutil.ReadFile("config.json")

	if err != nil {
		log.Printf("error reading file %s", err)
		return nil, err
	}
	data := Configuration{}

	err= json.Unmarshal([]byte(file), &data)
	if err != nil {
		log.Printf("error unmarchal config %s", err)
		return nil, err
	}	

	return &data,nil
}