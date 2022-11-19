package config

import (
	"fmt"
	"os"
	"strconv"
)

var (
	Port = 8000
	Env  = "dev"
)

func Init() error {
	getIntConfigFromENV("PORT", &Port)
	getStrConfigFromENV("ENV", &Env)

	// Validate Env
	if Env != "dev" && Env != "prod" {
		return fmt.Errorf("invalid env \"%s\" has been set", Env)
	}

	return nil
}

func getStrConfigFromENV(envVar string, configVar *string) {
	temp := os.Getenv(envVar)

	if temp != "" {
		*configVar = temp
	}
}

func getIntConfigFromENV(envVar string, configVar *int) {
	temp := os.Getenv(envVar)

	value, err := strconv.Atoi(temp)

	if err == nil {
		*configVar = value
	}
}
