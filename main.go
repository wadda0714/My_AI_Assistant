package main

import (
	"fmt"
	"os/exec"
	"time"
)

func main() {
	fmt.Println("Please enter your voice...")
	cmd := exec.Command("rec", "--encoding", "signed-integer", "--bits", "16", "--channels", "1", "--rate", "16000", "--endian", "little", "--type", "wav", "test.wav")

	//5 seconds to record
	cmd.Start()
	time.Sleep(5 * time.Second)
	cmd.Process.Kill()
	//Send to Whisper AI with curl
	fmt.Println("Sending to Whisper AI...")

	//Create data to send to Whisper AI with curl
	//curl -X POST -H "Content-Type: application/json" -d '{"audio": "test.wav"}' http://localhost:5000/recognize

	return

}
