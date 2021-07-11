package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/mrinjamul/go-json-server/models"
)

// ConfigFile the `config.json`
var ConfigFile string = "config.json"

// GetJSON return json from files
func GetJSON(filename string) interface{} {
	var jsondata interface{}
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Println(err)
	}
	err = json.Unmarshal(data, &jsondata)
	if err != nil {
		log.Println(err)
	}
	return jsondata
}

// ReadConfig reads config file from `config.json`
func ReadConfig() (models.Config, error) {
	var config models.Config
	b, err := ioutil.ReadFile(ConfigFile)
	if err != nil {
		return models.Config{}, nil
	}
	err = json.Unmarshal(b, &config)
	if err != nil {
		return models.Config{}, nil
	}
	return config, nil
}
