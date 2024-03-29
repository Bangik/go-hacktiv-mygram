package config

import (
	"fmt"
	"hacktiv-assignment-final/utils/common"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type DBConfig struct {
	Url string
}

type APIConfig struct {
	APIHost, APIPort string
}

type TokenConfig struct {
	ApplicationName     string
	JwtSignatureKey     []byte
	JwtSigningMethod    *jwt.SigningMethodHMAC
	AccessTokenLifeTime time.Duration
}

type Config struct {
	APIConfig
	DBConfig
	TokenConfig
}

func (c *Config) ReadConfig() error {
	err := common.LoadENV()
	if err != nil {
		return err
	}

	c.DBConfig = DBConfig{
		Url: os.Getenv("DB_URL"),
	}

	c.APIConfig = APIConfig{
		APIHost: os.Getenv("API_HOST"),
		APIPort: os.Getenv("API_PORT"),
	}

	appTokenExpire, err := strconv.Atoi(os.Getenv("APP_TOKEN_EXPIRE"))
	if err != nil {
		return err
	}

	accessTokenLifeTime := time.Duration(appTokenExpire) * time.Minute

	c.TokenConfig = TokenConfig{
		ApplicationName:     os.Getenv("APP_TOKEN_NAME"),
		JwtSignatureKey:     []byte(os.Getenv("APP_TOKEN_KEY")),
		JwtSigningMethod:    jwt.SigningMethodHS256,
		AccessTokenLifeTime: accessTokenLifeTime,
	}

	if c.DBConfig.Url == "" || c.APIConfig.APIHost == "" || c.APIConfig.APIPort == "" {
		return fmt.Errorf("missing required enivronment variables")
	}

	return nil
}

func NewConfig() (*Config, error) {
	config := &Config{}
	err := config.ReadConfig()
	if err != nil {
		return nil, err
	}

	return config, nil
}
