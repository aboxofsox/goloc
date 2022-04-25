package goloc

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"
)

type Config struct {
	Colors map[string]interface{}
}

// Check if the config file exists
func configCheck() bool {
	if _, err := os.Stat("goloc.json"); errors.Is(err, os.ErrNotExist) {
		return false
	}
	return true
}

// Load a config file if one exists.
func ConfigLoader() {
	var config Config

	// If a config file doesn't exist, exit the function.
	if !configCheck() {
		return
	}
	file, err := ioutil.ReadFile("goloc.json")
	if err != nil {
		log.Println("Failed to load config")
		return
	}

	err = json.Unmarshal(file, &config)
	if err != nil {
		log.Println("Failed to unmarshal config file.")
		return
	}
}
