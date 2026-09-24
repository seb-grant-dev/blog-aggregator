package config

import (
	"fmt"
	"os"
	"path/filepath"
	"encoding/json"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DBUrl string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

// Struct Methods
func (c *Config) SetUser(username string) {
	// Update config struct with username
	c.CurrentUserName = username
	// Write back to config file
	err := write(c)
	if err != nil {
		fmt.Errorf("Config could not be updated. %s",err)
	}

	fmt.Println("Config updated.")
}

// Package-global public functions
func Read() (Config) {
	configPath,err := getConfigFilePath()
	if err != nil {
		fmt.Printf("Error retrieving config file path. %s",err)
		os.Exit(1)
	}
	fmt.Printf("Reading config from %s\n",configPath)
	// Read config file from home directory
	fileContents, err := os.ReadFile(configPath)
	if err != nil {
		fmt.Errorf("Error reading config file: %s",err)
	}


	var config = Config{}
	// Unmarshal config json to Config object
	err = json.Unmarshal(fileContents,&config)
	if err != nil {
		fmt.Printf("Error reading JSON contents: %s\n",err)
	}
	// Return Config object
	return config
}

// Package-private helper functions
func getConfigFilePath() (string,error) {
	configDir,err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	fullPathParts := []string{configDir,"gator",configFileName}
	fullPath := filepath.Join(fullPathParts...)
	return fullPath,nil
}

func write(config *Config) error {
	configPath,err := getConfigFilePath()	
	if err != nil {
		fmt.Printf("Error retrieving config file path. %s",err)
		os.Exit(1)
	}

	// Encode the config object to json
	configBytes,err := json.Marshal(config)
	if err != nil {
		fmt.Errorf("Error writing config to file: %s",err)
	}
	// Write the encoded json to the config file
	err = os.WriteFile(configPath,configBytes, 0666)
	// Error if write fails
	if err != nil {
		return err
	}
	
	return nil
} 

