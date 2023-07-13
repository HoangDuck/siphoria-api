package config

import (
	"fmt"
	"go.uber.org/zap"
	"gopkg.in/yaml.v2"
	"hotel-booking-api/logger"
	"hotel-booking-api/model"
	"os"
	"strconv"
)

func LoadConfig(cfg *model.Config, env *string) {
	//file env(.yml) path
	var filePath string
	//check if environment is "dbdev" load file env.dev.yml
	if *env == "dbdev" {
		filePath = fmt.Sprintf("../env.%s.yml", "dev")
		fmt.Println("loading config from ", filePath)

	} else {
		//check if environment is "dev" or "pro" load file env.dev.yml or env.pro.yml
		//depend on command run program "make dev" --> load env.dev.yml; "make pro" --> load env.pro.yml
		filePath = fmt.Sprintf("../env.%s.yml", *env)
		fmt.Println("loading config from ", filePath)
	}

	fileContent, err := os.Open(filePath)
	if err != nil {
		logger.Error("Error open file", zap.Error(err))
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
		}
	}(fileContent)

	decoder := yaml.NewDecoder(fileContent)
	err = decoder.Decode(cfg)
	if err != nil {
		logger.Error("Error decode", zap.Error(err))
	}
}

func SetupEnv(cfg *model.Config) {
	err := os.Setenv("JWT_EXPIRED", strconv.Itoa(cfg.Server.JwtExpired))
	if err != nil {
		return
	}
	err1 := os.Setenv("JWT_REFRESH_EXPIRED", strconv.Itoa(cfg.Server.JwtRefreshExpired))
	if err1 != nil {
		return
	}
	err2 := os.Setenv("SECRET_KEY", cfg.Server.SecretKey)
	if err2 != nil {
		return
	}
	err3 := os.Setenv("SECRET_REFRESH_KEY", cfg.Server.SecretRefreshKey)
	if err3 != nil {
		return
	}
}
