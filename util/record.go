package record

import (
	"fmt"
	"github.com/wadda0714/My_AI_Assistant/config"

	"os/exec"
	"time"
)

func Record(cfg cfg.Config) error {

	fmt.Println("Please enter your voice...")
	cmd := exec.Command("rec", "--encoding", "signed-integer", "--bits", "16", "--channels", "1", "--rate", "16000", "--endian", "little", "--type", "wav", "test.wav")

	//time.Duration(cfg.Setting.Record_Limit) seconds to record

	cmd.Start()

	fmt.Println(cfg.Setting.Record_Limit)

	time.Sleep(time.Duration(cfg.Setting.Record_Limit) * time.Second)

	cmd.Process.Kill()

	return nil
}
