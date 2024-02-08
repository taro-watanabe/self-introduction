package config

import (
	"encoding/json"
	// "fmt"
	"os"
)

// ParseConfigFile reads and parses the config file
func ParseConfigFile(filePath string) (Config, error) {
	var config Config

	// Read the config file
	data, err := os.ReadFile(filePath)
	if err != nil {
		return config, err
	}

	// Unmarshal JSON data into config struct
	err = json.Unmarshal(data, &config)
	if err != nil {
		return config, err
	}

	return config, nil
}
