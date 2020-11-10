package db

import (
	"fmt"

	_ "github.com/lib/pq"
	"xorm.io/xorm"
)

type DbConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Database string `yaml:"database"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	SSL      string `yaml:"ssl"`
}

func GetEngine(config DbConfig) (*xorm.Engine, error) {
	engine, err := xorm.NewEngine("postgres", fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=%s",
		config.Username,
		config.Password,
		config.Host,
		config.Port,
		config.Database,
		config.SSL))
	if err != nil {
		return engine, err
	}

	{
		err = engine.Sync2(new(Manga))
		if err != nil {
			return engine, err
		}
		err = engine.Sync2(new(Chapter))
		if err != nil {
			return engine, err
		}
	}

	return engine, err
}

func CloseEngine(engine *xorm.Engine) error {
	return engine.Close()
}
