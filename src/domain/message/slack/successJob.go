package slack

import (
	"github.com/ashwanthkumar/slack-go-webhook"
	"os"
)

type SuccessJob struct {
	JobName string
	Message string
}

func (m *SuccessJob) GetPayload() Payload {

	attachment1 := slack.Attachment{}
	attachment1.AddField(slack.Field{Title: "Success Job", Value: m.JobName})
	attachment1.AddField(slack.Field{Title: "Message", Value: m.Message})

	return Payload{
		Text:        "Hexagonal Success Job ✅️",
		Attachments: []slack.Attachment{attachment1},
	}
}

func (m *SuccessJob) GetWebHook() string {
	return os.Getenv("SLACK_JOB_WEB_HOOK")
}
