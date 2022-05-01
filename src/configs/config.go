package configs

import (
	"strings"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type Config struct {
	*viper.Viper
}

func NewConfig() (*Config, error) {
	v := viper.New()
	v.SetConfigType("yaml")
	v.SetConfigName("app")
	v.AddConfigPath("./resources")
	v.AddConfigPath("../resources")
	v.AddConfigPath("../../resources")
	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	if err := v.ReadInConfig(); err != nil {
		return nil, errors.Wrap(err, "Cannot read config file")
	}
	return &Config{v}, nil
}
