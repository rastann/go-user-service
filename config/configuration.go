package config

import (
	"log"
	"os"

	"github.com/magiconair/properties"
)

type Config struct {
	Host     string `properties:"db.host"`
	Port     int    `properties:"db.port,default=5432"`
	Username string `properties:"db.username"`
	Password string `properties:"db.password"`
	DBName   string `properties:"db.dbname"`
	SSLMode  string `properties:"db.sslmode"`
}

func GetConfig() Config {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	p := properties.MustLoadFile(path+"/application.properties", properties.UTF8)
	var cfg Config
	if err := p.Decode(&cfg); err != nil {
		log.Fatal(err)
	}
	return cfg
}
