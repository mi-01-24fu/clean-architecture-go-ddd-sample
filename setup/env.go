package setup

import (
	"github.com/spf13/viper"
	"log"
)

type Env struct {
	DB_DBMS    string
	DB_DBNAME  string
	DB_HOST    string
	DB_PORT    string
	DB_USER    string
	DB_PASS    string
	DB_SSLMODE string
}

func NewEnv() *Env {
	env := Env{}

	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Can't find the file .env : ", err)
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatal("Environment can't be loaded: ", err)
	}

	return &env
}
