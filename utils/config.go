package utils

import (
	"github.com/pelletier/go-toml/v2"
	"log"
	"os"
)

type User struct {
	Uid       string `toml:"uid"`
	Pwd       string `toml:"pwd"`
	ProfileID string `toml:"profile_id"`
}

type Config struct {
	Users map[string]*User `toml:"users"`
}

var config Config

func init() {
	data, err := os.ReadFile("../config.toml")
	if err != nil {
		panic(err)
	}
	err = toml.Unmarshal(data, &config)
	if err != nil {
		panic(err)
	}
}

func GetUser() *User {
	for name, user := range config.Users {
		log.Println("user:", name)
		return user
	}
	panic("no user")
}
