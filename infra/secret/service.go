package secret

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type adapter struct{}

func CreateSecret() SecretAdapter {
	return &adapter{}
}

func (adapter *adapter) GetSecret(key string) string {
	value := os.Getenv(key)
	if value == "" {
		panic(fmt.Sprintf("env variable %s not found", key))
	}
	return value
}

func (adapter *adapter) InitEnvs() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	fmt.Println("Successfully loaded .env file")
}
