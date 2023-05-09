// infraのメソッドをinterfaceを介して呼び出す
package usecase

import (
	"context"
	//	"fmt"

	//	"github.com/wadda0714/My_AI_Assistant/config"
	"github.com/wadda0714/My_AI_Assistant/domain/model"
	"github.com/wadda0714/My_AI_Assistant/domain/query"
	"github.com/wadda0714/My_AI_Assistant/usecase/input"
	"github.com/wadda0714/My_AI_Assistant/usecase/output"
)

type assistant struct {
	assistantQuery query.Assistant
}

func NewAssistant(
	assistantQuery query.Assistant,
) Assistant {
	return &assistant{
		assistantQuery,
	}
}

func (u *assistant) TalkToAssistant(ctx context.Context, p *input.TalkToAssistantInput) (*output.TalkToAssistantOutput, error) {

	var AssistantDomain *model.GPTResponse
	var err error
	if AssistantDomain, err = u.assistantQuery.GetGPTResponse(ctx, p.APIKey); err != nil {
		return nil, err

	}

	return &output.TalkToAssistantOutput{
		FilePath: AssistantDomain.Text,
	}, nil
}
