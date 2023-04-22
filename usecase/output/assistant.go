package output

import "github.com/wadda0714/My_AI_Assistant/domain/model"

type StartConversation struct {
	AssistantConfig model.AssistantConfig
	QueryText       string
}

type ExecConversation struct {
	ResponseFromAssistant string
}

type EndConversation struct {
	e error
}
