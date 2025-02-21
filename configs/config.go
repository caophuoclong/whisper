package configs

import (
	"fmt"
	"strings"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	App      AppConfig      `mapstructure:"app"`
	Http     HttpConfig     `mapstructure:"http"`
	Logger   LoggerConfig   `mapstructure:"logger"`
	Postgres PostgresConfig `mapstructure:"POSTGRES"`
	JWT      JWTConfig      `mapstructure:"jwt"`
}
type AppConfig struct {
	Name    string `mapstructure:"name"`
	Version string `mapstructure:"version"`
}
type HttpConfig struct {
	Port string `mapstructure:"port"`
	Host string `mapstructure:"host"`
}

type JWTConfig struct {
	Secret              string `mapstructure:"jwt_secret"`
	AccessTokenExpires  int    `mapstructure:"jwt_expiration"`
	RefreshTokenExpires int    `mapstructure:"jwt_refresh_expiration"`
	Algorithm           string `mapstructure:"jwt_signing_algorithm"`
}

type LoggerConfig struct {
	Level string `mapstructure:"level"`
}

type PostgresConfig struct {
	Host     string `mapstructure:"HOST"`
	Port     int    `mapstructure:"PORT"`
	Name     string `mapstructure:"name"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	SSL      string `mapstructure:"sslmode"`
}

func LoadConfig(path string) (*Config, error) {
	if err := godotenv.Load(".env"); err != nil {
		fmt.Println("No .env file found, using system environment variables")
	}
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(path)

	viper.AutomaticEnv()

	viper.SetEnvPrefix("APP")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Could not read config", err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}
	fmt.Println(config.JWT)
	return &config, nil
}
