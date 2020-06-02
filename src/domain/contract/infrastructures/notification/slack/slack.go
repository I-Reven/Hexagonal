package slack

import message "github.com/I-Reven/Hexagonal/src/domain/message/slack"

type Slack interface {
	Send(m message.Message) []error
}
