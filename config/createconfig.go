package config

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func WriteConfigFile(filePath string, config Config) error {
	// Marshal config into JSON
	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}

	// Write JSON data to file
	err = os.WriteFile(filePath, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

func CreateConfigFile() Config {
	newConfig := Config{}

	fmt.Println("Answer the following questions to create a new config:")
	fmt.Print("Enter your name: ")
	fmt.Scanln(&newConfig.Name)

	fmt.Print("Enter your Birthday (and your birthyear if you feel fancy): ")
	fmt.Scanln(&newConfig.Birthday)

	fmt.Print("Enter your location: ")
	fmt.Scanln(&newConfig.Location)

	fmt.Print("Enter your workplace: ")
	fmt.Scanln(&newConfig.Workplace)

	fmt.Print("Enter your role: ")
	fmt.Scanln(&newConfig.Role)

	fmt.Print("Enter your Github username: ")
	fmt.Scanln(&newConfig.Github)

	fmt.Print("Enter your email: ")
	fmt.Scanln(&newConfig.Email)

	fmt.Print("Enter your programming languages (comma separated): ")
	var programming string
	fmt.Scanln(&programming)
	newConfig.Programming = append(newConfig.Programming, strings.Split(programming, ",")...)

	fmt.Print("Enter your hobbies (comma separated): ")
	var hobby string
	fmt.Scanln(&hobby)
	newConfig.Hobby = append(newConfig.Hobby, strings.Split(hobby, ",")...)

	// preview the config by calling the main
	fmt.Println(GenerateSelfIntroduction(newConfig))
	fmt.Println("Does the preview look good? (y/n)")
	var confirmation string
	fmt.Scanln(&confirmation)
	if confirmation == "y" {
		// Write new config to file
		err := WriteConfigFile("config/config.json", newConfig)
		if err != nil {
			fmt.Println("Error writing config file:", err)
			os.Exit(1)
		}
		fmt.Println("Config file created successfully.")
		return newConfig
	}
	fmt.Println("Config file not created.")
	os.Exit(1)
	return newConfig
}

func updateConfig(filePath string) Config {
	// Read existing config
	existingConfig, err := ParseConfigFile(filePath)
	if err != nil {
		fmt.Println("Error reading config file:", err)
		os.Exit(1)
	}

	fmt.Println("Update the following fields (press Enter to keep existing values):")
	fmt.Printf("Name [%s]: ", existingConfig.Name)
	var newName string
	fmt.Scanln(&newName)
	if newName != "" {
		existingConfig.Name = newName
	}

	fmt.Printf("Birthday [%s]: ", existingConfig.Birthday)
	var newBirthday string
	fmt.Scanln(&newBirthday)
	if newBirthday != "" {
		existingConfig.Birthday = newBirthday
	}

	fmt.Printf("Location [%s]: ", existingConfig.Location)
	var newLocation string
	fmt.Scanln(&newLocation)
	if newLocation != "" {
		existingConfig.Location = newLocation
	}

	fmt.Printf("Workplace [%s]: ", existingConfig.Workplace)
	var newWorkplace string
	fmt.Scanln(&newWorkplace)
	if newWorkplace != "" {
		existingConfig.Workplace = newWorkplace
	}

	fmt.Printf("Role [%s]: ", existingConfig.Role)
	var newRole string
	fmt.Scanln(&newRole)
	if newRole != "" {
		existingConfig.Role = newRole
	}

	fmt.Printf("Github [%s]: ", existingConfig.Github)
	var newGithub string
	fmt.Scanln(&newGithub)
	if newGithub != "" {
		existingConfig.Github = newGithub
	}

	fmt.Printf("Email [%s]: ", existingConfig.Email)
	var newEmail string
	fmt.Scanln(&newEmail)
	if newEmail != "" {
		existingConfig.Email = newEmail
	}

	fmt.Printf("Programming [%s]: ", strings.Join(existingConfig.Programming, ", "))
	var newProgramming string
	fmt.Scanln(&newProgramming)
	if newProgramming != "" {
		existingConfig.Programming = strings.Split(newProgramming, ",")
	}

	fmt.Printf("Hobby [%s]: ", strings.Join(existingConfig.Hobby, ", "))
	var newHobby string
	fmt.Scanln(&newHobby)
	if newHobby != "" {
		existingConfig.Hobby = strings.Split(newHobby, ",")
	}

	return existingConfig
}

func GenerateConfig(new bool, filePath string) {
	var newConfig Config
	if new {
		newConfig = CreateConfigFile()
	} else {
		newConfig = updateConfig(filePath)
	}

	// Write new config to file
	err := WriteConfigFile(filePath, newConfig)
	if err != nil {
		fmt.Println("Error writing config file:", err)
		os.Exit(1)
	}
	fmt.Println("Config file updated successfully.")
}
