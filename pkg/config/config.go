package config

import (
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

type Config struct {
	Port             string `mapstructure:"PORT"`
	DBHost           string `mapstructure:"DB_HOST"`
	DBName           string `mapstructure:"DB_NAME"`
	DBUser           string `mapstructure:"DB_USER"`
	DBPort           string `mapstructure:"DB_PORT"`
	DBPassword       string `mapstructure:"DB_PASSWORD"`
	GoauthClientID   string `mapstructure:"GOAUTH_CLIENTID"`
	GoauthClientSRC  string `mapstructure:"GOAUTH_CLIENR_SRC"`
	Redirect_URL     string `mapstructure:"REDIRECT_URL"`
	AccessToken      string `mapstructure:"ACCESS_TOKEN"`
	RefreshToken     string `mapstructure:"REFRESH_TOKEN"`
	RazorKey         string `mapstructure:"RZOR_KEYID"`
	RazorSec         string `mapstructure:"RAZOR_KEYSCR"`
	TWILIOACCOUNTSID string `mapstructure:"TWILIO_ACCOUNT_SID"`
	TWILIOAUTHTOKEN  string `mapstructure:"TWILIO_AUTHTOKEN"`
	TWILIOSERVICESID string `mapstructure:"TWILIO_SERVICES_ID"`
	StripeKey        string `mapstructure:"SECRET_KEY"`
	PublishableKey   string `mapstructure:"PUBLISHABLE_KEY"`
}

var envs = []string{
	"PORT", "DB_HOST", "DB_NAME", "DB_USER", "DB_PORT", "DB_PASSWORD", "SECRET_KEY",
	"GOAUTH_CLIENTID", "GOAUTH_CLIENR_SRC", "REDIRECT_URL", "ACCESS_TOKEN", "REFRESH_TOKEN", "PUBLISHABLE_KEY",
	"RZOR_KEYID", "RAZOR_KEYSCR", "TWILIO_SERVICES_ID", "TWILIO_ACCOUNT_SID", "TWILIO_AUTHTOKEN",
}
var config Config

func LoadConfig() (Config, error) {

	viper.AddConfigPath("./")
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	for _, env := range envs {
		if err := viper.BindEnv(env); err != nil {
			return config, err
		}
	}

	if err := viper.Unmarshal(&config); err != nil {
		return config, err
	}

	if err := validator.New().Struct(&config); err != nil {
		return config, err
	}

	return config, nil
}

func GetConfig() Config {
	return config
}
