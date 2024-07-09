package configs

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppConfig     `json:"appConfig"`
	DBConfig      `json:"dbConfig"`
	ServiceConfig `json:"serviceConfig"`
	JwtSecret     string `json:"jwtSecret"`
}

type ServiceConfig struct {
	RestPort string `json:"restPort"`
	GrpcPort string `json:"grpcPort"`
}

type AppConfig struct {
	DomainName string `json:"domainName"`
	AppEmailId string `json:"appEmailId"`
}

type DBConfig struct {
	DbUrl      string `json:"dbUrl"`
	DbName     string `json:"dbName"`
	DbUser     string `json:"dbUser"`
	DbPassword string `json:"dbPassword"`
}

// various loading techniques for config
func LoadConfigFronEnvFile() (*Config, error) {
	err := godotenv.Load()

	if err != nil {
		return nil, err
	}

	config := &Config{
		AppConfig: AppConfig{
			DomainName: os.Getenv("DOMAIN_NAME"),
		},
		DBConfig: DBConfig{
			DbUrl:      os.Getenv("DB_URL"),
			DbName:     os.Getenv("DB_NAME"),
			DbUser:     os.Getenv("DB_USER"),
			DbPassword: os.Getenv("DB_PASSWORD"),
		},
		ServiceConfig: ServiceConfig{
			RestPort: os.Getenv("REST_PORT"),
			GrpcPort: os.Getenv("GRPC_PORT"),
		},
		JwtSecret: os.Getenv("JWT_SECRET"),
	}

	return config, nil
}
