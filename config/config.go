// config
package config

type Config struct {
	Setting Setting "toml:'setting'"
	URL     URL     "toml:'url'"
	Whisper Whisper "toml:'whisper'"
}
type Setting struct {
	Version      string "toml:'version'"
	Record_Limit int    "toml:'record_limit'"
}

type URL struct {
	Whisper_URL  string "toml:'whisper_URL'"
	GPT3_URL     string "toml:'gpt3_URL'"
	VOICEBOX_URL string "toml:'VOICEVOX_URL'"
}
type Whisper struct {
	URL         string "toml:'URL'"
	ContentType string "toml:'Content-Type'"
	file        string "toml:'file'"
	model       string "toml:'model'"
}
