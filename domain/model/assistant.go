package model

import (
// "github.com/volatiletech/null/v8"
)

type Assistant struct {
	Name         string `json:"name"`
	Whisper_URL  string `json:"whisper_url"`
	GPT3_URL     string `json:"gpt3_url"`
	VOICEVOX_URL string `json:"voicevox_url"`
	Record_Limit int64  `json:"record_limit"`
}

type ResAssistant struct {
	ResponseFromAssistant  string `json:"response_from_assistant"`
	query                  string `json:"query"`
	AssistantConfigMessage string `json:"assistant_config_message"`
}

func NewResAssistant(
	ResponseFromAssistant string,
	query string,
	AssistantConfigMessage string,
) *ResAssistant {

	return &ResAssistant{
		ResponseFromAssistant:  ResponseFromAssistant,
		query:                  query,
		AssistantConfigMessage: AssistantConfigMessage,
	}
}

func NewAssistant(
	Name string,
	Whisper_URL string,
	GPT3_URL string,
	VOICEVOX_URL string,
	Record_Limit int64,
) *Assistant {

	return &Assistant{
		Name:         Name,
		Whisper_URL:  Whisper_URL,
		GPT3_URL:     GPT3_URL,
		VOICEVOX_URL: VOICEVOX_URL,
		Record_Limit: Record_Limit,
	}
}
