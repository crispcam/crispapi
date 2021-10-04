package crisps

import (
	"fmt"
	"github.com/spf13/viper"
)

func LoadConfig() (Config, error) {
	var config Config
	viper.SetConfigName("config")
	viper.AddConfigPath("config")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		return config, err
	}
	if err := viper.Unmarshal(&config); err != nil {
		fmt.Println(err)
		return config, err
	}
	return config, nil
}

type Config struct {
	Firestore FirestoreConfig `mapstructure:"firestore"`
	CrispCam  CrispCamConfig  `mapstructure:"crispcam"`
	Port      string          `mapstructure:"port"`
	Project   string          `mapstructure:"project"`
}

type FirestoreConfig struct {
	CollectionName string `mapstructure:"collection_name"`
}

type CrispCamConfig struct {
	Services  ServicesConfig `mapstructure:"services"`
	Paths     PathsConfig    `mapstructure:"paths"`
	Threshold float64        `mapstructure:"threshold"`
}

type ServicesConfig struct {
	Catalog        string `mapstructure:"catalog"`
	AutoML         string `mapstructure:"automl"`
	LocalML        string `mapstructure:"localml"`
	Assets         string `mapstructure:"assets"`
	Auth           string `mapstructure:"auth"`
	Barcodes       string `mapstructure:"barcodes"`
	CrispCam       string `mapstructure:"crispcam"`
	CrispCamReview string `mapstructure:"crispcam_review"`
	CrispCamSave   string `mapstructure:"crispcam_save"`
	Frontend       string `mapstructure:"frontend"`
	Reviews        string `mapstructure:"reviews"`
}

type PathsConfig struct {
	Catalog      Path `mapstructure:"catalog"`
	Reviews      Path `mapstructure:"reviews"`
	AutoML       Path `mapstructure:"automl"`
	CrispCamSave Path `mapstructure:"crispcam_save"`
}

type Path struct {
	Path    string `mapstructure:"path"`
	All     string `mapstructure:"all"`
	Many    string `mapstructure:"many"`
	Single  string `mapstructure:"single"`
	Reviews string `mapstructure:"reviews"`
	Review  string `mapstructure:"review"`
	Rating  string `mapstructure:"rating"`
	Ratings string `mapstructure:"ratings"`
	Save    string `mapstructure:"save"`
}