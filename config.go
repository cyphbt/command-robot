package main

import (
	"encoding/json"
	"log"
	"os"
)

type config struct {
	Path string   `json:"path"`
	Port string   `json:"port"`
	Cmd  string   `json:"cmd"`
	Args []string `json:"args"`
}

var Config = new(config)

func init() {
	content, err := os.ReadFile("./config.json")
	if err != nil {
		log.Fatal("read file error: ", err)
	}

	if err = json.Unmarshal(content, Config); err != nil {
		log.Fatal("unmarshal error: ", err)
	}
}
