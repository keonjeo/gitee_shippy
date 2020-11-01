package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// MySQL Configuration
var (
	DBHost               string // DBHost xxx
	DBPort               string // DBPort xxx
	DBDatabase           string // DBDatabase xxx
	DBUser               string // DBUser xxx
	DBPassword           string // DBPassword xxx
	DBCharset            string // DBCharset xxx
	DBMaxOpenConnections string // DBMaxOpenConnections xxx
	DBMaxIdleConnections string // DBMaxIdleConnections xxx
	DBConnMaxLifetime    string // DBConnMaxLifetime xxx
)

func loadEnvFile() {
	err := godotenv.Load("./config/.env")
	if err != nil {
		log.Fatalf("init => Fail to godotenv.Load, err: %v", err)
	}

	DBHost = os.Getenv("DB_HOST")
	if DBHost == "" {
		log.Printf("Not Configure DB_HOST ENV, will use default value")
		DBHost = "127.0.0.1"
	}

	DBPort = os.Getenv("DB_PORT")
	if DBPort == "" {
		log.Printf("Not Configure DB_PORT ENV, will use default value")
		DBPort = "3306"
	}

	DBDatabase = os.Getenv("DB_DATABASE")
	if DBDatabase == "" {
		log.Printf("Not Configure DB_DATABASE ENV, will use default value")
		DBDatabase = "gorm_demo"
	}

	DBUser = os.Getenv("DB_USER")
	if DBUser == "" {
		log.Printf("Not Configure DB_USER ENV, will use default value")
		DBUser = "root"
	}

	DBPassword = os.Getenv("DB_PASSWORD")
	if DBPassword == "" {
		log.Printf("Not Configure DB_PASSWORD ENV, will use default value")
		DBPassword = "123456"
	}

	DBCharset = os.Getenv("DB_CHARSET")
	if DBCharset == "" {
		log.Printf("Not Configure DB_CHARSET ENV, will use default value")
		DBCharset = "utf8"
	}

	DBMaxOpenConnections = os.Getenv("DB_MAX_OPEN_CONNECTIONS")
	if DBMaxOpenConnections == "" {
		log.Printf("Not Configure DB_MAX_OPEN_CONNECTIONS ENV, will use default value")
		DBMaxOpenConnections = "100"
	}

	DBMaxIdleConnections = os.Getenv("DB_MAX_IDLE_CONNECTIONS")
	if DBMaxIdleConnections == "" {
		log.Printf("Not Configure DB_MAX_IDLE_CONNECTIONS ENV, will use default value")
		DBMaxIdleConnections = "10"
	}

	DBConnMaxLifetime = os.Getenv("DB_CONN_MAX_LIFE_TIME")
	if DBConnMaxLifetime == "" {
		log.Printf("Not Configure DB_CONN_MAX_LIFE_TIME ENV, will use default value")
		DBConnMaxLifetime = "100"
	}
}
