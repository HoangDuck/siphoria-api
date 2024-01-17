package config

import (
	"errors"
	"fmt"
	"go.uber.org/zap"
	"gopkg.in/yaml.v2"
	"hotel-booking-api/logger"
	"hotel-booking-api/model"
	"os"
	"strconv"
	"time"
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
	//open file yml config
	fileContent, err := os.Open(filePath)
	if err != nil {
		logger.Error("Error open file", zap.Error(err))
	}
	//close file config yml (this function runs when function LoadConfig almost done, before this function close
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
		}
	}(fileContent)
	//decode content file
	decoder := yaml.NewDecoder(fileContent)
	err = decoder.Decode(cfg)
	if err != nil {
		logger.Error("Error decode", zap.Error(err))
	}
}

func SetupEnv(cfg *model.Config) {
	//set environment jwt expired
	err := os.Setenv("JWT_EXPIRED", strconv.Itoa(cfg.Server.JwtExpired))
	if err != nil {
		return
	}
	//set environment jwt refresh expired
	err1 := os.Setenv("JWT_REFRESH_EXPIRED", strconv.Itoa(cfg.Server.JwtRefreshExpired))
	if err1 != nil {
		return
	}
	//set environment secret key
	err2 := os.Setenv("SECRET_KEY", cfg.Server.SecretKey)
	if err2 != nil {
		return
	}
	//set environment secret refresh key
	err3 := os.Setenv("SECRET_REFRESH_KEY", cfg.Server.SecretRefreshKey)
	if err3 != nil {
		return
	}
}

// Message is what greeters will use to greet guests.
type Message string

// NewMessage creates a default Message.
func NewMessage(phrase string) Message {
	return Message(phrase)
}

// NewGreeter initializes a Greeter. If the current epoch time is an even
// number, NewGreeter will create a grumpy Greeter.
func NewGreeter(m Message) Greeter {
	var grumpy bool
	if time.Now().Unix()%2 == 0 {
		grumpy = true
	}
	return Greeter{Message: m, Grumpy: grumpy}
}

// Greeter is the type charged with greeting guests.
type Greeter struct {
	Grumpy  bool
	Message Message
}

// Greet produces a greeting for guests.
func (g Greeter) Greet() Message {
	if g.Grumpy {
		return Message("Go away!")
	}
	return g.Message
}

// NewEvent creates an event with the specified greeter.
func NewEvent(g Greeter) (Event, error) {
	if g.Grumpy {
		return Event{}, errors.New("could not create event: event greeter is grumpy")
	}
	return Event{Greeter: g}, nil
}

// Event is a gathering with greeters.
type Event struct {
	Greeter Greeter
}

// Start ensures the event starts with greeting all guests.
func (e Event) Start() {
	msg := e.Greeter.Greet()
	fmt.Println(msg)
}
