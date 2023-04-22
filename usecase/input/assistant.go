package input

import (
	"github.com/wadda0714/My_AI_Assistant/domain/model"
)

type ExecConversation struct {
	ConversationContext model.ConversationContext
	AssistantConfig     model.AssistantConfig
	QueryText           string
}

type StartConversation struct {
	Record_limit int
	ConfigPath   string
}

type EndConversation struct {
	QueryText             string
	ResponseFromAssistant model.ResponseFromAssistant
}
