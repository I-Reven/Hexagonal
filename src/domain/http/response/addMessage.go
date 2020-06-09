package response

type AddMessage struct {
	MessageId string `json:"message_id"`
}

func (r *AddMessage) Make(messageId string) *AddMessage {
	r.MessageId = messageId
	return r
}
