package config

import (
	"encoding/json"
	"os"
)

func Load(_path string) (c *Config) {
	res, e := os.ReadFile(_path)
	if e != nil {
		panic(e)
	}
	e = json.Unmarshal(res, &c)
	if e != nil {
		panic(e)
	}
	return
}
