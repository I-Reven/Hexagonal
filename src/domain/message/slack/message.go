package slack

import "github.com/ashwanthkumar/slack-go-webhook"

type (
	Payload slack.Payload

	Message interface {
		GetPayload() Payload
		GetWebHook() string
	}
)
