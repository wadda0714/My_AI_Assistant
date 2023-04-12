package main

import (
	"fmt"
	"github.com/wadda0714/My_AI_Assistant/config"

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
	time.Sleep(time.Duration(config.Setting.Record_Limit) * time.Second)
	cmd.Process.Kill()
	//Send to Whisper AI with curl
	fmt.Println("Sending to Whisper AI...(" + config.URL.Whisper_URL + ")")

	//Create data to send to Whisper AI with curl

	return

}
