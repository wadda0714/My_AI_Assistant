package main

import (
	"fmt"
	"os"

	"context"
	"github.com/wadda0714/My_AI_Assistant/config"
	"github.com/wadda0714/My_AI_Assistant/usecase/input"

	"github.com/wadda0714/My_AI_Assistant/usecase"
	//"bytes"
	//"encoding/json"

	"github.com/BurntSushi/toml"
	"github.com/joho/godotenv"
)

func main() {
	ctx := context.Background()
	var config config.Config
	fmt.Println("Reading config.toml...")
	_, err := toml.DecodeFile("../config.toml", &config)
	if err != nil {
		fmt.Println("failed to read config.toml")

		return
	}
	fmt.Println(config.URL.Whisper_URL)

	//Read .env file
	err = godotenv.Load("../.env")
	if err != nil {
		fmt.Println("failed to read .env file")
		return
	}
	APIKey := os.Getenv("API_Key")

	Assistant := usecase.GeneratedInitializeAssistantDI()
	res, err := Assistant.TalkToAssistant(ctx, &input.TalkToAssistantInput{APIKey: APIKey})
	fmt.Println(res.FilePath)

	/*	//Record your voice

		err = record.Record(config)

		if err != nil {
			fmt.Println("failed to record your voice", err)
			fmt.Println("Retry...")
			//Retry
		}

		//Send to Whisper AI
		fmt.Println("Sending to Whisper AI...")

		API_KEY := os.Getenv("API_Key")

		header_map := map[string]string{"Content-Type": "multipart/form-data", "Authorization": "Bearer " + API_KEY}

		resp, err := http.MakeHTTPRequest(config.URL.Whisper_URL, header_map)

		if err != nil {
			fmt.Println("failed to send to Whisper AI:", err)
		}
		//Convert result to string and print
		fmt.Println(resp) // this is not working */
	return

}
