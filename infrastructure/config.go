package infrastructure

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DB struct {
		Production struct {
			Host     string
			User     string
			Password string
			DbName   string
			Port     string
			SslMode  string
			TimeZone string
		}
		Development struct {
			Host     string
			User     string
			Password string
			DbName   string
			Port     string
			SslMode  string
			TimeZone string
		}
	}
}

func NewConfig() *Config {
	err := godotenv.Load(".env." + "local")
	if err != nil {
		log.Fatal("Error loading .env file")
		panic(err)
	}

	c := new(Config)

	c.DB.Production.Host = os.Getenv("HOST")
	c.DB.Production.User = os.Getenv("USER")
	c.DB.Production.Password = os.Getenv("PASSWORD")
	c.DB.Production.DbName = os.Getenv("DBNAME")
	c.DB.Production.Port = os.Getenv("PORT")
	c.DB.Production.SslMode = os.Getenv("SSLMODE")
	c.DB.Production.TimeZone = os.Getenv("TIMEZONE")

	c.DB.Development.Host = os.Getenv("HOST")
	c.DB.Development.User = os.Getenv("USER")
	c.DB.Development.Password = os.Getenv("PASSWORD")
	c.DB.Development.DbName = os.Getenv("DBNAME")
	c.DB.Development.Port = os.Getenv("PORT")
	c.DB.Development.SslMode = os.Getenv("SSLMODE")
	c.DB.Development.TimeZone = os.Getenv("TIMEZONE")

	return c
}
