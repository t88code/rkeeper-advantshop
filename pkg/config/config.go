package config

import (
	"fmt"
	"gopkg.in/gcfg.v1"
	"io"
	"log"
	"os"
	check "rkeeper-advantshop/pkg/license"
	"sync"
)

type (
	Config struct {
		TELEGRAM struct {
			BotToken string
			Debug    bool
		}
		LOG struct {
			Debug bool
		}
		SERVICE struct {
			PORT int
		}
		ADVANTSHOP struct {
			URL                     string
			ApiKey                  string
			Username                string
			Password                string
			RPS                     int
			Timeout                 int
			ApiKeyExpire            int
			OrderPrefix             string
			OrderSource             string
			Currency                string
			BonusCost               int
			CheckOrderItemAvailable bool
		}
		MAXMA struct {
			URL                     string
			ApiKey                  string
			Username                string
			Password                string
			RPS                     int
			Timeout                 int
			ApiKeyExpire            int
			OrderPrefix             string
			OrderSource             string
			Currency                string
			BonusCost               int
			CheckOrderItemAvailable bool
		}
		RK7REF struct {
			URL  string
			User string
			Pass string
		}
		RK7MID struct {
			URL           string
			User          string
			Pass          string
			OrderTypeCode int
			TableCode     int
			StationCode   int
			TimeoutError  int
		}
		XMLINTERFACE struct {
			Type      int
			UserName  string
			Password  string
			Token     string
			RestCode  int
			ProductID string
			Guid      string
			URL       string
		}
	}
)

var cfg Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		check.Check()
		err := os.MkdirAll("logs", 0770)
		if err != nil {
			fmt.Println(err)
		}

		file, err := os.OpenFile("logs/config.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0640)
		if err != nil {
			fmt.Println(err)
		}

		multiWriter := io.MultiWriter(file, os.Stdout)

		logger := log.New(multiWriter, "MAIN ", log.Ldate|log.Ltime|log.Lshortfile)
		logger.Print("Config:>Read application configurations")

		pwd, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		err = gcfg.ReadFileInto(&cfg, fmt.Sprintf("%s/config.ini", pwd))
		if err != nil {
			logger.Fatalf("Config:>Failed to parse gcfg data: %s", err)
		} else {
			logger.Print("Config:>Config is read")
		}
	})

	return &cfg
}
