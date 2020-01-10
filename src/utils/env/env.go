package env

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"path/filepath"
)

func LoadEnv() {
	getPath := func() string {
		envName := os.Getenv("GO_ENV")
		if envName == "" {
			return ".env"
		} else {
			return ".env." + envName
		}
	}

	execPath, err := os.Getwd()
	if err != nil {
		log.Fatal("! Env. Get exec path error!")
		os.Exit(1)
	}

	path := filepath.Join(execPath, "env", getPath())
	errEnvLoad := godotenv.Load(path)
	if errEnvLoad != nil {
		log.Fatal("! Error Loading DotEnv")
		os.Exit(1)
	}

	for _, pair := range os.Environ() {
		log.Println(pair)
	}
}
