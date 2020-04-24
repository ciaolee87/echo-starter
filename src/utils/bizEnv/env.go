package bizEnv

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"path/filepath"
)

func init() {
	LoadEnv()
}

func LoadEnv() {
	getPath := func() string {
		envName := os.Getenv("GO_ENV")

		log.Println("GO_ENV", envName)
		if envName == "" {
			return ".env"
		} else {
			return ".evn." + envName
		}
	}

	execPath, err := os.Getwd()
	if err != nil {
		log.Fatal("! Env. Get exec path error!")
	}

	path := filepath.Join(execPath, "env", getPath())
	log.Println("envPath", path)
	errEnvLoad := godotenv.Load(path)
	if errEnvLoad != nil {
		log.Fatal("! Error Loading DotEnv")
	}

	for _, pair := range os.Environ() {
		log.Println(pair)
	}
}

func Get(key string) string {
	return os.Getenv(key)
}
