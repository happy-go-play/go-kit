package internal

type Message struct {
	MsgType  string    `json:"msgtype"`
	Text     *Text     `json:"text,omitempty"`
	Markdown *Markdown `json:"markdown,omitempty"`
	Link     *Link     `json:"link,omitempty"`
	At       *At       `json:"at,omitempty"`
}

type Text struct {
	Content string `json:"content"`
}

type Markdown struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

type Link struct {
	Title      string `json:"title"`
	Text       string `json:"text"`
	MessageURL string `json:"messageUrl"`
	PicURL     string `json:"picUrl"`
}

type At struct {
	IsAtAll   bool     `json:"isAtAll"`
	AtUserIds []string `json:"atUserIds"`
	AtMobiles []string `json:"atMobiles"`
}
