package slack

import (
	"github.com/ashwanthkumar/slack-go-webhook"
	"os"
)

type FailedJob struct {
	JobName   string
	Message   string
	RetryUrl  string
	CancelUrl string
	Error     error
}

func (m *FailedJob) GetPayload() Payload {

	attachment1 := slack.Attachment{}
	attachment1.AddField(slack.Field{Title: "Failed Job", Value: m.JobName})
	attachment1.AddField(slack.Field{Title: "Message", Value: m.Message})
	attachment1.AddField(slack.Field{Title: "Error", Value: m.Error.Error()})
	if m.RetryUrl != "" {
		attachment1.AddAction(slack.Action{Type: "button", Text: "Try Again", Url: m.RetryUrl, Style: "primary"})
	}
	if m.CancelUrl != "" {
		attachment1.AddAction(slack.Action{Type: "button", Text: "Cancel", Url: m.CancelUrl, Style: "danger"})
	}

	return Payload{
		Text:        "Hexagonal Failed Job ⚠️",
		Attachments: []slack.Attachment{attachment1},
	}
}

func (m *FailedJob) GetWebHook() string {
	return os.Getenv("SLACK_JOB_WEB_HOOK")
}
