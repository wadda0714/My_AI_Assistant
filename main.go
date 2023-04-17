package main

import (
	"fmt"
	"github.com/wadda0714/My_AI_Assistant/config"

	"github.com/wadda0714/My_AI_Assistant/util"

	"os"
	"os/exec"

	"github.com/BurntSushi/toml"
)

func main() {
	var config config.Config
	fmt.Println("Reading config.toml...")
	_, err := toml.DecodeFile("config.toml", &config)
	if err != nil {
		fmt.Println("failed to read config.toml")

		return
	}
	fmt.Println(config.URL.Whisper_URL)

	//Record your voice

	err = util.Record(config)

	if err != nil {
		fmt.Println("failed to record your voice", err)
		fmt.Println("Retry...")
		//Retry
	}

	//Send to Whisper AI with curl to use exec Commandd
	fmt.Println("Sending to Whisper AI...")
	API_KEY := os.Getenv("API_Key")

	cmd := exec.Command("curl", "https://api.whisper.com/v1/convert",
		"-H", "Authorization: Bearer "+API_KEY, "-H",
		"Content-Type:", "audio/wav", "-F", "model='whisper-1'", "-F", "file='test.wav'")

	result, err := cmd.Output()
	if err != nil {
		fmt.Println("failed to send to Whisper AI:", err)
	}
	//Convert result to string and print
	fmt.Println(string(result))
	return

}
