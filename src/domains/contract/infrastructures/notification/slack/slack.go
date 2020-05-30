package slack

import message "github.com/I-Reven/Hexagonal/src/domains/message/slack"

type Slack interface {
	Send(m message.Message) []error
}
