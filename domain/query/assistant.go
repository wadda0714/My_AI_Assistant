package query

import (
	"github.com/wadda0714/My_AI_Assistant/domain/model"
)

type Assistant interface {
	GetResponseFromAssistant(filename string) (*model.ResAssistant, error)
	ChangeAssistantConfig(assistantConfig *model.Assistant) error
}
