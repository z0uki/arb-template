package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

var Conf Config

type Config struct {
	PrivateKey      string `yaml:"privateKey"`
	SigningKey      string `yaml:"signingKey"`
	EthNode         string `yaml:"ethNode"`
	EthNodeWss      string `yaml:"ethNodeWss"`
	OpenseaKey      string `yaml:"openseaKey"`
	FeishuWebhook   string `yaml:"feishuWebhook"`
	FeishuSecretKey string `yaml:"feishuSecretKey"`
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
