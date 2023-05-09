package query

import (
	"context"
	"github.com/wadda0714/My_AI_Assistant/domain/model"
)

// Assistantは、infraから呼び出されるインタフェース
type Assistant interface {
	GetGPTResponse(ctx context.Context, url string) (*model.GPTResponse, error)
	GetWhisperResponse(ctx context.Context, url string) (*model.WhisperResponse, error)
	GetVOICEVOXResponse(ctx context.Context, url string) (*model.VOICEVOXResponse, error)
}
