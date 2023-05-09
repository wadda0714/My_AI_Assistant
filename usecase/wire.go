package usecase

import (
	"github.com/google/wire"
	"github.com/wadda0714/My_AI_Assistant/infra/query"
)

func InitializeAssistantDI() (_ Assistant) {
	wire.Build(
		NewAssistant,
		query.NewAssistantQuery,
	)
	return
}
