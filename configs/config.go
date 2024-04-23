package configs

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

// This struct isn't designed to be used externally.
// Use TeleConfig variable instead to get config
// variables.
// Config holds the configuration settings for the Telegram bot.
type Config struct {
    // BotToken is the API token required for authenticating the bot with Telegram.
    BotToken string `yaml:"bot token"`

    // Username of the telegram account.
    Username string `yaml:"username"`

    // Password of the telegram account
    Password string `yaml:"password"`

    // AdminID is the Telegram user ID of the administrator for the bot.
    AdminID string `yaml:"admin"`
}

// GetBotToken returns the bot token from the configuration.
func (c *Config) GetBotToken() string {
    return c.BotToken
}

// GetUsername returns the username from the configuration.
func (c *Config) GetUsername() string {
    return c.Username
}

// GetPassword returns the password from the configuration.
func (c *Config) GetPassword() string {
    return c.Password
}

// GetAdminID returns the Telegram user ID of the administrator.
func (c *Config) GetAdminID() string {
    return c.AdminID
}

// Don not change it at runtime, can lead to bugs
var TeleConfig *Config

// Initialization of config form config.yaml file.
// Path must be specifeid in production
func init() {
    // TODO: Update the file path in production.
    // Currently set to a specific file path for development purposes.
    file, err := os.Open("D:\\Work\\TeleBot\\configs\\config.yaml")
    if err != nil {
        log.Fatalf("[FATAL]: Error opening config.yaml - %s\n", err)
    }
    defer file.Close()

    // Unmarshal the config file
    decoder := yaml.NewDecoder(file)
    TeleConfig = new(Config)

    err = decoder.Decode(TeleConfig)
    if err != nil {
        log.Fatalf("[FATAL]: Error unmarshalling config file: %s\n", err)
    }

    log.Println("[INFO]: Config has been loaded successfully")
}
