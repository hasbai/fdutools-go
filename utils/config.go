package utils

import (
	"github.com/pelletier/go-toml/v2"
	"log"
	"os"
	"os/user"
)

type User struct {
	Uid       string `toml:"uid" json:"uid"`
	Pwd       string `toml:"pwd" json:"pwd"`
	ProfileID string `toml:"profile_id" json:"profileID"`
	Name      string `toml:"name" json:"name"`
}

type Config struct {
	User User `toml:"user"`
}

var config Config

func GetUser() *User {
	return &config.User
}

func exists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	log.Panicf("failed to check if file exists: %v", err)
	return false
}

func configPath() string {
	u, err := user.Current()
	if err != nil {
		log.Panicf("failed to get current user: %v", err)
	}
	dir := u.HomeDir + "/.config/fdutools"
	err = os.MkdirAll(dir, 0700)
	if err != nil {
		log.Panicf("failed to create config directory: %v", err)
	}
	return dir + "/config.toml"
}

func SaveConfig() {
	path := configPath()
	data, err := toml.Marshal(config)
	if err != nil {
		log.Panicf("failed to marshal config: %v", err)
	}
	err = os.WriteFile(path, data, 0600)
	if err != nil {
		log.Panicf("failed to write config file: %v", err)
	}
}

func loadConfig() {
	path := configPath()
	if !exists(path) {
		SaveConfig()
		log.Println("create config file in", path)
	}
	data, err := os.ReadFile(path)
	if err != nil {
		log.Panicf("failed to read config file: %v", err)
	}
	err = toml.Unmarshal(data, &config)
	if err != nil {
		log.Panicf("failed to parse config file: %v", err)
	}
}

func init() {
	loadConfig()
}
