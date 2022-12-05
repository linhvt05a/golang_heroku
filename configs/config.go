package configs

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	PostgreDriver          string        `mapstructure:"POSTGRES_DRIVER"`
	PostgresUrl            string        `mapstructure:"POSTGRES_URL"`
	Port                   string        `mapstructure:"PORT"`
	Origin                 string        `mapstructure:"ORIGIN"`
	AccessTokenPrivateKey  string        `mapstructure:"ACCESS_TOKEN_PRIVATE_KEY"`
	AccessTokenPublicKey   string        `mapstructure:"ACCESS_TOKEN_PUBLIC_KEY"`
	RefreshTokenPrivateKey string        `mapstructure:"REFRESH_TOKEN_PRIVATE_KEY"`
	RefreshTokenPublicKey  string        `mapstructure:"REFRESH_TOKEN_PUBLIC_KEY"`
	AccessTokenExpiresIn   time.Duration `mapstructure:"ACCESS_TOKEN_EXPIRED_IN"`
	RefreshTokenExpiresIn  time.Duration `mapstructure:"REFRESH_TOKEN_EXPIRED_IN"`
	AccessTokenMaxAge      int           `mapstructure:"ACCESS_TOKEN_MAXAGE"`
	RefreshTokenMaxAge     int           `mapstructure:"REFRESH_TOKEN_MAXAGE"`
	TwilioAccountSid       string        `mapstructure:"TWILIO_ACCOUNT_SID"`
	TwilioAuthenToken      string        `mapstructure:"TWILIO_AUTH_TOKEN"`
	TwilioVerifySid        string        `mapstructure:"VERIFY_SERVICE_SID"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("./")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
