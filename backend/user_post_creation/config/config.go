package config

import (
	"encoding/json"
	"io/ioutil"
	"user_post_creation/model"
)

var ViperConfig model.Config

func ReadConfigFile(filename string) error {
	// Read the JSON file
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	// Create a new Config struct
	config := &model.Config{}

	// Unmarshal the JSON data into the Config struct
	err = json.Unmarshal(data, config)
	if err != nil {
		return err
	}

	return nil
}
