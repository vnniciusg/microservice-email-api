package config

import "github.com/spf13/viper"

type conf struct {
	SMTPServer   string `mapstructure:"SMTPServer"`
	SMTPPort     int    `mapstructure:"SMTPPort"`
	SMTPUsername string `mapstructure:"SMTPUsername"`
	SMTPPassowrd string `mapstructure:"SMTPPassowrd"`
}

func LoadConfig() (*conf, error) {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()

	if err != nil {
		panic(err)
	}

	cfg := &conf{
		SMTPServer:   viper.GetString("SMTPServer"),
		SMTPPort:     viper.GetInt("SMTPPort"),
		SMTPUsername: viper.GetString("SMTPUsername"),
		SMTPPassowrd: viper.GetString("SMTPPassowrd"),
	}

	return cfg, err
}
