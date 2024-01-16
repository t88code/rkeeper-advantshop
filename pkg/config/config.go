package config

import (
	"fmt"
	"gopkg.in/gcfg.v1"
	"os"
	check "rkeeper-advantshop/pkg/license"
	"rkeeper-advantshop/pkg/logging"
)

type config struct {
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
}

var cfg = new(config)

func GetConfig(logger *logging.Logger, appName string) (*config, error) {
	check.Check()
	pwd, err := os.Getwd()
	if err != nil {
		logger.Fatal(err)
	}
	err = gcfg.ReadFileInto(cfg, fmt.Sprintf("%s/config.ini", pwd))
	if err != nil {
		return nil, fmt.Errorf("Config:>Failed to parse gcfg data: %s", err)
	}
	logger.Info("Config:>Config is read")
	return cfg, nil
}
