package main

import (
	"fmt"
	"github.com/wadda0714/My_AI_Assistant/config"

	"os"
	"os/exec"
	"time"

	"github.com/BurntSushi/toml"
)

func main() {
	var config cfg.Config
	fmt.Println("Reading config.toml...")
	_, err := toml.DecodeFile("config.toml", &config)
	if err != nil {
		fmt.Println("failed to read config.toml")

		return
	}
	fmt.Println("Please enter your voice...")
	cmd := exec.Command("rec", "--encoding", "signed-integer", "--bits", "16", "--channels", "1", "--rate", "16000", "--endian", "little", "--type", "wav", "test.wav")

	//time.Duration(config.Setting.Record_Limit) seconds to record

	cmd.Start()
	fmt.Println(config.Setting.Record_Limit)
	time.Sleep(time.Duration(config.Setting.Record_Limit) * time.Second)
	cmd.Process.Kill()
	//Send to Whisper AI with curl to use exec Commandd
	fmt.Println("Sending to Whisper AI...")
	API_KEY := os.Getenv("API_Key")

	cmd = exec.Command("curl", "https://api.whisper.com/v1/convert",
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
