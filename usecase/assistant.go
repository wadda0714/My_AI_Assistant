package usecase

import (
	"github.com/wadda0714/My_AI_Assistant/usecase/input"
	"github.com/wadda0714/My_AI_Assistant/usecase/output"
)

// "github.com/wadda0714/My_AI_Assistant/usecase/input"
// "github.com/wadda0714/My_AI_Assistant/usecase/output"

// userがやりたい処理の全体の流れであり、細かい処理はinrfaに任せる(infraに依存しない)
type Assistant interface {
	TalkToAssistant(*input.TalkToAssistantInput) (*output.TalkToAssistantOutput, error)
}
