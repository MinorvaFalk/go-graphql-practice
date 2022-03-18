package utils

import (
	"log"

	"github.com/spf13/viper"
)

type Env struct {
	DBHost     string `mapstructure:"DB_HOST"`
	DBUser     string `mapstructure:"DB_USERNAME"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBName     string `mapstructure:"DB_NAME"`
	DBPort     string `mapstructure:"DB_PORT"`
	Port       string `mapstructure:"PORT"`
}

func NewEnv(l *log.Logger) *Env {
	env := Env{}

	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {

		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Check if no .env file
			l.Panicln(".env file not found")

		} else {
			l.Panicln(err.Error())
		}
	}

	if err := viper.Unmarshal(&env); err != nil {
		l.Panicln(err.Error())
	}

	l.Println(env)

	return &env
}
