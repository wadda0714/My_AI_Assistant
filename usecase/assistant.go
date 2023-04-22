package usecase

import (
	"github.com/wadda0714/My_AI_Assistant/usecase/input"
        "github.com/wadda0714/My_AI_Assistant/usecase/output"
)
//userがやりたい処理の全体の流れであり、細かい処理はinrfaに任せる(infraに依存しない)
type Assistant interface {
	// 会話を開始する
	StartConversation(p input.StartConversation) (result *output.StartConversation, err error)
	// 会話を終了する
	EndConversation(p input.EndConversation) error (result *output.EndConversation, err error)

	// 会話を実行する
	ExecConversation(p input.ExecConversation) error (result *output.ExecConversation, err error)
}
