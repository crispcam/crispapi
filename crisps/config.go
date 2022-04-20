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
	DB        DBConfig        `mapstructure:"db"`
	Port      string          `mapstructure:"port"`
	Project   string          `mapstructure:"project"`
	Redis     RedisConfig     `mapstructure:"redis"`
	App       AppConfig       `mapstructure:"app"`
	OIDC      OIDC            `mapstructure:"oidc"`
	Session   Session         `mapstructure:"session"`
}

type FirestoreConfig struct {
	CollectionName string `mapstructure:"collection_name"`
}

type CrispCamConfig struct {
	Services  ServicesConfig `mapstructure:"services"`
	Paths     PathsConfig    `mapstructure:"paths"`
	Review    ReviewConfig   `mapstructure:"review"`
	Threshold float64        `mapstructure:"threshold"`
	Save      SaveConfig     `mapstructure:"save"`
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
	Search         string `mapstructure:"search"`
	SearchPublic   string `mapstructure:"search_public"`
}

type PathsConfig struct {
	Catalog      Path `mapstructure:"catalog"`
	Reviews      Path `mapstructure:"reviews"`
	AutoML       Path `mapstructure:"automl"`
	CrispCam     Path `mapstructure:"crispcam"`
	CrispCamSave Path `mapstructure:"crispcam_save"`
	Search       Path `mapstructure:"search"`
	Auth         Path `mapstructure:"auth"`
}

type Path struct {
	Path       string `mapstructure:"path"`
	All        string `mapstructure:"all"`
	Many       string `mapstructure:"many"`
	Single     string `mapstructure:"single"`
	Reviews    string `mapstructure:"reviews"`
	Review     string `mapstructure:"review"`
	Rating     string `mapstructure:"rating"`
	Ratings    string `mapstructure:"ratings"`
	Save       string `mapstructure:"save"`
	Scan       string `mapstructure:"scan"`
	Categories string `mapstructure:"categories"`
	Search     string `mapstructure:"search"`
	User       string `mapstructure:"user"`
	BasicUser  string `mapstructure:"basic-user"`
	Reviewed   string `mapstructure:"reviewed"`
	Flavours   string `mapstructure:"flavours"`
	Delete     string `mapstructure:"delete"`
	Update     string `mapstructure:"update"`
	CSV        string `mapstructure:"csv"`
}

type ReviewConfig struct {
	Colour  string `mapstructure:"colour"`
	Broken  bool   `mapstructure:"broken"`
	Timeout int    `mapstructure:"timeout"`
}

type SaveConfig struct {
	Bucket string `mapstructure:"bucket"`
	Folder string `mapstructure:"folder"`
}

type DBConfig struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Name     string `mapstructure:"name"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

type RedisConfig struct {
	Host      string `mapstructure:"host"`
	Port      string `mapstructure:"port"`
	KeyPrefix string `mapstructure:"key_prefix"`
}

type AppConfig struct {
	Name    string `mapstructure:"name"`
	Version string `mapstructure:"version"`
	Cache   bool   `mapstructure:"cache"`
}

type OIDC struct {
	Endpoint     string   `mapstructure:"endpoint"`
	ClientID     string   `mapstructure:"client_id"`
	ClientSecret string   `mapstructure:"client_secret"`
	Redirect     string   `mapstructure:"redirect"`
	Provider     string   `mapstructure:"provider"`
	Scopes       []string `mapstructure:"scopes"`
	Callback     string   `mapstructure:"callback"`
}
type Session struct {
	Key string `mapstructure:"key"`
}
