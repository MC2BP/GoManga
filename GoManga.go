package main

import (
	"GoManga/db"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Db db.DbConfig `yaml:"db"`
}

func main() {
	//Logging
	/*	Log output to file
		file, err := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		log.SetOutput(file)
		log.SetOutput(os.Stdout)
	*/
	log.SetOutput(os.Stdout)
	log.Print("Starting GoManga")

	log.Print("Loading config file \"config.yaml\"")
	config, err := loadConfig()
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.GetEngine(config.Db)
	if err != nil {
		log.Fatal(err)
	}

}

func loadConfig() (Config, error) {
	config := Config{}
	configFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		return config, err
	}
	err = yaml.Unmarshal(configFile, &config)
	return config, err
}
