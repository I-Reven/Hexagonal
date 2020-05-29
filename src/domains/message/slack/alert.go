package slack

import "github.com/ashwanthkumar/slack-go-webhook"

type Alert struct {
	Title       string
	Description string
}

func (m *Alert) GetPayload() Payload {

	attachment1 := slack.Attachment{}
	attachment1.AddField(slack.Field{Title: "Title", Value: m.Title})
	attachment1.AddField(slack.Field{Title: "Description", Value: m.Description})

	return Payload{
		Text:        "Hexagonal Alert!",
		Username:    "AlertBot",
		Channel:     "#i-raven-core-alive",
		IconEmoji:   ":monkey_face:",
		Attachments: []slack.Attachment{attachment1},
	}
}

func (m *Alert) GetWebHook() string {
	return "https://hooks.slack.com/services/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"
}
