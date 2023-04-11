package main

import (
	"fmt"
	"os/exec"
	"time"
)

func main() {
	fmt.Println("Please enter your voice!")
	cmd := exec.Command("rec", "--encoding", "signed-integer", "--bits", "16", "--channels", "1", "--rate", "16000", "--endian", "little", "--type", "wav", "test.wav")

	//5 seconds to record
	cmd.Start()
	time.Sleep(5 * time.Second)
	cmd.Process.Kill()

	fmt.Println("Sending to Whisper AI...")
	return

}
