package models

//SMSConfig : ""
type SMSConfig struct {
	URL      string
	User     string
	Password string
	Senderid string
	Channel  string
	DCS      int
	Flashsms int
	Route    string
}

//SingleSMS : ""
type SingleSMS struct {
	Mobile string `json:"mobile" bson:"mobile,omitempty"`
	Msg    string `json:"msg" bson:"msg,omitempty"`
}
