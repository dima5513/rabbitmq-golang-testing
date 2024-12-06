package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)



type AmqpServerURL struct {
	URL string
}
type Config struct {
	AmqpServerURL string
}


func NewRabbitConfig() *Config {

	if err := godotenv.Load(); err != nil {
        log.Print("No .env file found")
    }

    return &Config{
        AmqpServerURL: GetEnv("AMQP_SERVER_URL", "123"),
	}
    
}


func GetEnv(key string, defaultVal string) string {
	
    value, exists := os.LookupEnv(key)
	if exists {
		return value
	}
	return defaultVal

}