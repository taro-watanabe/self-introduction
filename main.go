package main

import (
	"flag"
	"fmt"
	"main/config"
)

func readConfigFromFile(filePath string) (config.Config, error) {
	// Parse the config file
	return config.ParseConfigFile(filePath)
}

func main() {
	filePath := "config/config.json"
	edit := flag.Bool("edit", false, "Edit existing config")
	flag.Parse()

	if *edit {
		config.GenerateConfig(false, filePath)
		return
	}
	parsedConfig, err := readConfigFromFile(filePath)
	if err != nil {
		fmt.Println("Config file not found. Creating a new one...")
		config.GenerateConfig(true, filePath)
		return
	}

	selfIntroduction := config.GenerateSelfIntroduction(parsedConfig)
	fmt.Println(selfIntroduction)
}
