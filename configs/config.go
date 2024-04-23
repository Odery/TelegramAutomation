package configs

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

// This struct isn't designed to be used externally.
// Use TeleConfig variable instead to get config
// variables.
type Config struct {
	BotToken string `yaml:"bot token"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	AdminID string `yaml:"admin"`
}

func (c *Config) GetBotToken() string {
	return c.BotToken
}

func (c *Config) GetUsername() string {
	return c.Username
}

func (c *Config) GetPassword() string {
	return c.Password
}

func (c *Config) GetAdminID() string {
	return c.AdminID
}

// Don not change it at runtime, can lead to bugs
var TeleConfig *Config

func init() {
	// Load config file
	// TODO must change path in production
	file, err := os.Open("D:\\Work\\TeleBot\\configs\\config.yaml")
	if err != nil {
		log.Fatalf("[FATAL]: Error getting config.yaml - %s\n", err)
	}
	defer file.Close()

	// Unmarshall config file
	decoder := yaml.NewDecoder(file)
	TeleConfig = new(Config)

	err = decoder.Decode(TeleConfig)
	if err != nil {
		log.Fatalf("[FATAL]: Error unmarshalling config file. %s\n", err)
	}

	log.Println("[INFO]: Config has been loaded successfully")
}
