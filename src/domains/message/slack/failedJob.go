package slack

import "github.com/ashwanthkumar/slack-go-webhook"

type FailedJob struct {
	JobName string
	Message string
	Error   error
}

func (m *FailedJob) GetPayload() Payload {

	attachment1 := slack.Attachment{}
	attachment1.AddField(slack.Field{Title: "Failed Job", Value: m.JobName})
	attachment1.AddField(slack.Field{Title: "Message", Value: m.Message})
	attachment1.AddField(slack.Field{Title: "Error", Value: m.Error.Error()})

	return Payload{
		Text:        "Hexagonal Failed Job!",
		Username:    "FailedJobBot",
		Channel:     "#i-raven-core-alive",
		IconEmoji:   ":monkey_face:",
		Attachments: []slack.Attachment{attachment1},
	}
}

func (m *FailedJob) GetWebHook() string {
	return "https://hooks.slack.com/services/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"
}
