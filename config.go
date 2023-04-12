// config
package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
)

type Config struct {
	Setting setting
	URL     url
}
type setting struct {
	APIKey       string "toml:'SecretAPI_key'"
	Version      string "toml:'version'"
	Record_Limit int    "toml:'record_limit'"
}

type url struct {
	Whisper_URL  string "toml:'whisper_URL'"
	GPT3_URL     string "toml:'gpt3_URL'"
	VOICEBOX_URL string "toml:'VOICEVOX_URL'"
}
