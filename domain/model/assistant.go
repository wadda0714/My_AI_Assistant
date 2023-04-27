package model

import (
// "github.com/volatiletech/null/v8"
)

type GPTResponse struct {
	Text string `json:"text"`
}

func NewGPTResponse(Text string) *GPTResponse {
	return &GPTResponse{
		Text: Text,
	}
}

type WhisperResponse struct {
	Text string `json:"file"`
}

func NewWhisperResponse(Text string) *WhisperResponse {
	return &WhisperResponse{
		Text: Text,
	}
}

type VOICEVOXResponse struct {
	file string `json:"file"`
}

func NewVOICEVOXResponse(file string) *VOICEVOXResponse {
	return &VOICEVOXResponse{
		file: file,
	}
}
