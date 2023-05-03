package query

import (
	"bytes"
	"fmt"
	"github.com/wadda0714/My_AI_Assistant/domain/model"
	"github.com/wadda0714/My_AI_Assistant/domain/query"
	"net/http"
)

type assistantQuery struct {
}

func NewAssistantQuery() query.Assistant {
	return &assistantQuery{}
}

func (*assistantQuery) GetGPTResponse(url string) (model.GPTResponse, error) {

	url = "https://api.openai.com/v1/chat/completions"

	jsonStr := []byte(`{"model": "gpt-3.5-turbo","messages": [{"role": "user", "content": "こんにちは"}],"stream": true}`)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return model.GPTResponse{}, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+"")
	req.Header.Set("Accept", "text/event-stream")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return model.GPTResponse{}, err
	}
	defer resp.Body.Close()
	fmt.Println("Response Status:", resp.Status)

	buffer := make([]byte, 1024)
	for {
		n, err := resp.Body.Read(buffer)
		if err != nil {
			break
		}
		fmt.Println(string(buffer[0:n]))
	}

	return model.GPTResponse{}, nil
}

func (*assistantQuery) GetWhisperResponse(url string) (model.WhisperResponse, error) {

	return model.WhisperResponse{}, nil
}

func (*assistantQuery) GetVOICEVOXResponse(url string) (model.VOICEVOXResponse, error) {

	return model.VOICEVOXResponse{}, nil
}
