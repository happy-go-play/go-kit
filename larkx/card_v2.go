package larkx

import "encoding/json"

type cardV2 struct {
	Schema string        `json:"schema"`
	Header *cardV2Header `json:"header,omitempty"`
	Body   cardV2Body    `json:"body"`
}

type cardV2Header struct {
	Title    cardV2PlainText `json:"title"`
	Template string          `json:"template,omitempty"`
}

type cardV2PlainText struct {
	Tag     string `json:"tag"`
	Content string `json:"content"`
}

type cardV2Body struct {
	Elements []cardV2MarkdownElement `json:"elements"`
}

type cardV2MarkdownElement struct {
	Tag     string `json:"tag"`
	Content string `json:"content"`
}

func buildMarkdownCardV2(body, title, headerTemplate string) (string, error) {
	card := cardV2{
		Schema: "2.0",
		Body: cardV2Body{
			Elements: []cardV2MarkdownElement{
				{Tag: "markdown", Content: body},
			},
		},
	}
	if title != "" {
		if headerTemplate == "" {
			headerTemplate = "blue"
		}
		card.Header = &cardV2Header{
			Title:    cardV2PlainText{Tag: "plain_text", Content: title},
			Template: headerTemplate,
		}
	}
	b, err := json.Marshal(card)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
