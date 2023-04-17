// config
package cfg

type Config struct {
	Setting Setting "toml:'setting'"
	URL     URL     "toml:'url'"
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
