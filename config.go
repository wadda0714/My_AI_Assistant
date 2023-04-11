// config
package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Config struct {
	Port string
	Db   string
	Host string
	User string
}
