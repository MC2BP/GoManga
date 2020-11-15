package main

import (
	"GoManga/collector"
	"GoManga/collector/mangadex"
	"GoManga/db"
	"GoManga/storage"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Db           db.DbConfig  `yaml:"db"`
	SourceConfig SourceConfig `yaml:"source"`
	MangaFolder  string       `yaml:"mangaFolder"`
}

type SourceConfig struct {
	MangaDexConfig mangadex.MangaDexConfig `yaml:"mangadex"`
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
		return
	}

	//Get db engine
	_, err = db.GetEngine(config.Db)
	if err != nil {
		log.Fatal(err)
		return
	}

	//Get collectors
	collectors := getCollectors(config.SourceConfig)
	err = storage.PrepareMangaFolder(config.MangaFolder, collectors)
	if err != nil {
		log.Fatal("Failed to prepare the folder for mangas. ", err)
	}
	run(collectors)
}

func run(collectors map[string]collector.Collector) error {
	for _, collector := range collectors {
		manga, _ := collector.GetTitle(55529)
		chapters, _ := collector.GetChapters(manga)
		log.Print(chapters)
	}
	return nil
}

func getCollectors(config SourceConfig) map[string]collector.Collector {
	collectors := map[string]collector.Collector{}

	//Mangadex
	mangaDex := mangadex.CreateMangaDex(config.MangaDexConfig)
	collectors[mangaDex.GetName()] = mangaDex

	return collectors
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
