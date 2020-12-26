package main

import (
	"flag"
	"fmt"
	// "geobase/internal/database"
	models "geobase/internal/models"
	"geobase/internal/server"
	"geobase/internal/storage"
	"github.com/ilyakaznacheev/cleanenv"
	// "io"
	// "io/ioutil"
	"log"
	"os"
	"path/filepath"
)

const address = ":1234"

func main() {
	cfg := prepareConfig()
	// dat, err := ioutil.ReadFile("/static/data")
	store := storage.New(cfg.Datapath)
	srv := server.New(store, cfg)
	log.Println("starting server")
	err := srv.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func prepareConfig() *models.Config {
	var cfg models.Config
	// configFile := getConfigFile()
	dir, _ := os.Getwd()
	dir = filepath.Join(dir, "config.yml")
	fmt.Println(dir)
	if err := cleanenv.ReadConfig(dir, &cfg); err != nil {
		fmt.Printf("Unable to get app configuration due to: %s\n", err.Error())
	}

	if err := cleanenv.ReadEnv(&cfg); err != nil {
		fmt.Printf("Unable to retrieve app configuration due to: %s\n", err.Error())
		os.Exit(1)
	}
	return &cfg
}

func getConfigFile() string {
	configFile := flag.String("config", "config.yml", "config file")
	return *configFile
}
