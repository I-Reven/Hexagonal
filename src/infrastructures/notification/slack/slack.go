package slack

import (
	message "github.com/I-Reven/Hexagonal/src/domains/message/slack"
	"github.com/ashwanthkumar/slack-go-webhook"
)

type Slack struct{}

func (*Slack) Send(m message.Message) []error {
	return slack.Send(m.GetWebHook(), "", slack.Payload(m.GetPayload()))
}
