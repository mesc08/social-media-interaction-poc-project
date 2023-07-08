package config

import (
	"encoding/json"
	"io/ioutil"
)

var ViperConfig Config

type Config struct {
	AwsRegion    string `json:"awsregion"`
	AwsID        string `json:"awsid"`
	AwsSecretKey string `json:"awssecret"`
	PSGUser      string `json:"psguser"`
	PSGHost      string `json:"psghost"`
	PSGPass      string `json:"psgpass"`
	PSGDB        string `json:"psgdb"`
	PSGPort      string `json:"psgport"`
	S3Bucket     string `json:"s3bucket"`
}

func ReadConfigFile(filename string) error {
	// Read the JSON file
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	// Create a new Config struct
	config := &Config{}

	// Unmarshal the JSON data into the Config struct
	err = json.Unmarshal(data, config)
	if err != nil {
		return err
	}

	return nil
}
