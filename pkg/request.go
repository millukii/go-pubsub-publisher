package pkg


type Request struct {
	Message      Message `json:"message"` 
	Subscription string  `json:"subscription"`
}
type Message struct {
	Data        string    `json:"data"`
	Attributes        map[string]string `json:"attributes"`    
	MessageID   string    `json:"messageId"`
	PublishTime string `json:"publishTime"`
}
