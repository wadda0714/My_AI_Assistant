package query

import (
	"bytes"
	"context"
	"fmt"
	"github.com/wadda0714/My_AI_Assistant/domain/model"
	"github.com/wadda0714/My_AI_Assistant/domain/query"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
)

type assistantQuery struct {
}

func NewAssistantQuery() query.Assistant {
	return &assistantQuery{}
}

func (*assistantQuery) GetGPTResponse(ctx context.Context, APIKey string) (*model.GPTResponse, error) {

	url := "https://api.openai.com/v1/chat/completions"

	jsonStr := []byte(`{"model": "gpt-3.5-turbo","messages": [{"role": "user", "content": "こんにちは"}],"stream": true}`)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return &model.GPTResponse{
			Text: "error",
		}, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+APIKey)
	req.Header.Set("Accept", "text/event-stream")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return &model.GPTResponse{
			Text: "error",
		}, err
	}
	defer resp.Body.Close()
	fmt.Println("Response Status:", resp.Status)

	childCtx, cancel := context.WithCancel(ctx)
	buffer := make([]byte, 1024)
	str := ""
	for {
		n, err := resp.Body.Read(buffer)
		if err != nil {
			break
		}
		if strings.Contains(string(buffer[0:n]), "、") || strings.Contains(string(buffer[0:n]), "。") {
			str += string(buffer[0:n])
			go GetResponseFromVOICEVOX(childCtx, str)
			str = ""
		} else {
			str += string(buffer[0:n])
		}
	}

	return &model.GPTResponse{
		Text: "success",
	}, nil
}

func (*assistantQuery) GetWhisperResponse(ctx context.Context, url string) (*model.WhisperResponse, error) {

	return &model.WhisperResponse{}, nil
}

func (*assistantQuery) GetVOICEVOXResponse(ctx context.Context, url string) (*model.VOICEVOXResponse, error) {

	return &model.VOICEVOXResponse{}, nil
}

func GetResponseFromVOICEVOX(ctx context.Context, res string) {

	// リクエストURL
	url := "http://localhost:50021/synthesis"

	// パラメータ設定
	params := map[string]string{
		"text":         "こんにちは、VOICEVOX API。",
		"speaker":      "show",
		"emotion":      "normal",
		"pitch":        "1.0",
		"speed":        "1.0",
		"volume":       "1.0",
		"audio_format": "wav",
	}

	// ファイル読み込み
	file, err := os.Open("/path/to/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// リクエストボディ設定
	requestBody := &bytes.Buffer{}
	writer := multipart.NewWriter(requestBody)
	part, err := writer.CreateFormFile("text", "input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = io.Copy(part, file)
	if err != nil {
		fmt.Println(err)
		return
	}
	for key, value := range params {
		_ = writer.WriteField(key, value)
	}
	err = writer.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	// リクエスト送信
	request, err := http.NewRequest("POST", url, requestBody)
	if err != nil {
		fmt.Println(err)
		return
	}
	request.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer response.Body.Close()

	// レスポンス取得
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(body))

	return

}
