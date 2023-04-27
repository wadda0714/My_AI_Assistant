package query

import (
	"github.com/wadda0714/My_AI_Assistant/domain/model"
)

// Assistantは、infraから呼び出されるインタフェース
type Assistant interface {
	GetGPTResponse(url string) (model.GPTResponse, error)
	GetWhisperResponse(url string) (model.WhisperResponse, error)
	GetVOICEVOXResponse(url string) (model.VOICEVOXResponse, error)
}
