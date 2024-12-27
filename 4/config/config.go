package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBConnString string
}

func GetDBEnv() (map[string]string, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	env := make(map[string]string)
	env["POSTGRES_HOST"] = os.Getenv("POSTGRES_HOST")
	env["POSTGRES_PORT"] = os.Getenv("POSTGRES_PORT")
	env["POSTGRES_USER"] = os.Getenv("POSTGRES_USER")
	env["POSTGRES_PASSWORD"] = os.Getenv("POSTGRES_PASSWORD")
	env["POSTGRES_DB"] = os.Getenv("POSTGRES_DB")

	return env, nil
}

func GetConfig() *Config {
	env, err := GetDBEnv()
	if err != nil {
		panic(err)
	}

	fmt.Println(env)
	return &Config{
		DBConnString: fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s", env["POSTGRES_HOST"], env["POSTGRES_PORT"], env["POSTGRES_USER"], env["POSTGRES_PASSWORD"], env["POSTGRES_DB"]),
	}
}
