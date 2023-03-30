package config

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
)

var Conf Config

type Config struct {
	PrivateKey string `yaml:"privateKey"`
	EthNode    string `yaml:"ethNode"`
	EthNodeWss string `yaml:"ethNodeWss"`
}

func Load(path string) error {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(file, &Conf)

	return err
}

func GenerateDefaultConfig() {
	defaultConfig := Config{}
	yamlData, err := yaml.Marshal(&defaultConfig)

	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile("config.yml", yamlData, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
