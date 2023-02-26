package config

import (
	"encoding/json"
	"os"
	"sync"

	"github.com/dexises/4room/api_service/pkg/logging"
)

type Config struct {
	Port       string `json:"port"`
	IAMService struct {
		URL string `json:"url" env-required:"true"`
	} `json:"iam_service" env-required:"true"`
	UserService struct {
		URL string `json:"url" env-required:"true"`
	} `json:"user_service" env-required:"true"`
	PostService struct {
		URL string `json:"url" env-required:"true"`
	} `json:"post_service" env-required:"true"`
	VoteService struct {
		URL string `json:"url" env-required:"true"`
	} `json:"vote_service" env-required:"true"`
}

var (
	instance *Config
	once     sync.Once
)

func GetConfigInstance() *Config {
	once.Do(func() {
		logger := logging.GetLoggerInstance()
		logger.PrintInfo("read application config")
		instance = &Config{}
		if err := instance.configParser("config.json"); err != nil {
			logger.PrintError(err)
		}
	})
	return instance
}

func (c *Config) configParser(path string) error {
	f, err := os.OpenFile(path, os.O_RDONLY|os.O_SYNC, 0)
	if err != nil {
		return err
	}
	defer f.Close()

	return json.NewDecoder(f).Decode(c)
}
