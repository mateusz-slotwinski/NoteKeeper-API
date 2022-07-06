package config

import (
	log "log"
	os "os"
	strconv "strconv"

	gin "github.com/gin-gonic/gin"
	dotenv "github.com/joho/godotenv"
)

var (
	Port               string
	Host               string
	Mode               string
	DB_ClusterEndpoint string
	DB_Username        string
	DB_Password        string
	DB_Collection      string
	JWT_Secret         string
	JWT_ExpiresIn      int
)

func Init() {
	err := dotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	Port = os.Getenv("PORT")
	Host = os.Getenv("HOST")
	Mode = os.Getenv("MODE")
	DB_ClusterEndpoint = os.Getenv("DB_ENDPOINT")
	DB_Username = os.Getenv("DB_USERNAME")
	DB_Password = os.Getenv("DB_PASSWORD")
	DB_Collection = os.Getenv("DB_COLLECTION")
	JWT_Secret = os.Getenv("JWT_SECRET")

	JWT_ExpiresIn, _ = strconv.Atoi(os.Getenv("JWT_EXPIRES_IN"))

	if Mode == "PROD" {
		gin.SetMode(gin.ReleaseMode)
	}
}
