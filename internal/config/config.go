package config

import "github.com/spf13/viper"

type EnvConfig struct {
	DBPassword string `mapstructure:"POSTGRES_PASSWORD"`
	DBUser     string `mapstructure:"POSTGRES_USER"`
	DBName     string `mapstructure:"POSTGRES_DB"`
	DBHost     string `mapstructure:"POSTGRES_HOST"`
	DBPort     string `mapstructure:"POSTGRES_PORT"`

	KafkaPort string `mapstructure:"KAFKA_PORT"`
	KafkaHost string `mapstructure:"KAFKA_HOST"`
}

func LoadEnvConfig(path string) (config EnvConfig, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
