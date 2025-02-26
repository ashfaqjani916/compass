package config

import (
	"log"

	"github.com/spf13/viper"
)

func FetchSource() ([]string){

	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yml")

	if err:= viper.ReadInConfig(); err != nil{
		log.Fatalf("Error reading the config file %s",err)
	}

	sources := viper.GetStringSlice("hacakthon_sources")

	if len(sources) == 0 {
		log.Fatal("No sources found in the configuration file")
	}

	return sources
	
}
