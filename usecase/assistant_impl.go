// infraのメソッドをinterfaceを介して呼び出す
package usecase

import (
	"fmt"

	"github.com/wadda0714/My_AI_Assistant/config"
	"github.com/wadda0714/My_AI_Assistant/infra/query"
	"github.com/wadda0714/My_AI_Assistant/usecase/input"
	"github.com/wadda0714/My_AI_Assistant/usecase/output"
)

type assistant struct {
	whisperURL  string
	gpt3URL     string
	voicevoxURL string
	recordLimit int
}

func NewAssistant(cfg config.Config) Assistant {
	return &assistant{
		whisperURL:  cfg.URL.Whisper_URL,
		gpt3URL:     cfg.URL.GPT3_URL,
		voicevoxURL: cfg.URL.VOICEVOX_URL,
		recordLimit: cfg.Setting.Record_Limit,
	}
}

func (*assistant) TalkToAssistant(p *input.TalkToAssistantInput) (*output.TalkToAssistantOutput, error) {

	p := NewAssistantQuery()
	res, err := p.GetGPTResponse("")
	if err != nil {
		return nil, err
	}
	fmt.Println(res)
	return &output.TalkToAssistantOutput{
		FilePath: "test",
	}, nil
}
