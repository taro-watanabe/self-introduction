package config

import (
	"bufio"
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
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Answer the following questions to create a new config:")
	fmt.Print("Enter your name: ")
	scanner.Scan()
	newConfig.Name = scanner.Text()

	fmt.Print("Enter your Birthday (and your birthyear if you feel fancy): ")
	scanner.Scan()
	newConfig.Birthday = scanner.Text()

	fmt.Print("Enter your location: ")
	scanner.Scan()
	newConfig.Location = scanner.Text()

	fmt.Print("Enter your workplace: ")
	scanner.Scan()
	newConfig.Workplace = scanner.Text()

	fmt.Print("Enter your role: ")
	scanner.Scan()
	newConfig.Role = scanner.Text()

	fmt.Print("Enter your Github username: ")
	scanner.Scan()
	newConfig.Github = scanner.Text()

	fmt.Print("Enter your email: ")
	scanner.Scan()
	newConfig.Email = scanner.Text()

	fmt.Print("Enter your programming languages (comma separated, no need for space after comma): ")
	scanner.Scan()
	newConfig.Programming = strings.Split(scanner.Text(), ",")

	fmt.Print("Enter your hobbies (comma separated, no need for space after comma): ")
	scanner.Scan()
	newConfig.Hobby = strings.Split(scanner.Text(), ",")

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
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Update the following fields (press Enter to keep existing values):")
	fmt.Printf("Name [%s]: ", existingConfig.Name)
	scanner.Scan()
	var newName string = scanner.Text()
	if scanner.Text() != "" {
		existingConfig.Name = newName
	}

	fmt.Printf("Birthday [%s]: ", existingConfig.Birthday)
	var newBirthday string
	scanner.Scan()
	newBirthday = scanner.Text()
	if newBirthday != "" {
		existingConfig.Birthday = newBirthday
	}

	fmt.Printf("Location [%s]: ", existingConfig.Location)
	var newLocation string
	scanner.Scan()
	newLocation = scanner.Text()
	if newLocation != "" {
		existingConfig.Location = newLocation
	}

	fmt.Printf("Workplace [%s]: ", existingConfig.Workplace)
	var newWorkplace string
	scanner.Scan()
	newWorkplace = scanner.Text()
	if newWorkplace != "" {
		existingConfig.Workplace = newWorkplace
	}

	fmt.Printf("Role [%s]: ", existingConfig.Role)
	var newRole string
	scanner.Scan()
	newRole = scanner.Text()
	if newRole != "" {
		existingConfig.Role = newRole
	}

	fmt.Printf("Github [%s]: ", existingConfig.Github)
	var newGithub string
	scanner.Scan()
	newGithub = scanner.Text()
	if newGithub != "" {
		existingConfig.Github = newGithub
	}

	fmt.Printf("Email [%s]: ", existingConfig.Email)
	var newEmail string
	scanner.Scan()
	newEmail = scanner.Text()
	if newEmail != "" {
		existingConfig.Email = newEmail
	}

	fmt.Printf("Programming [%s]: ", strings.Join(existingConfig.Programming, ", "))
	var newProgramming string
	scanner.Scan()
	newProgramming = scanner.Text()
	if newProgramming != "" {
		existingConfig.Programming = strings.Split(newProgramming, ",")
	}

	fmt.Printf("Hobby [%s]: ", strings.Join(existingConfig.Hobby, ", "))
	var newHobby string
	scanner.Scan()
	newHobby = scanner.Text()
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
