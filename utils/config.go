package utils

import (
	"github.com/rs/zerolog/log"
	"github.com/shakh9006/numbers-store/consts"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName(consts.ENV_FILE_NAME)
	viper.SetConfigType(consts.ENV_FILE_TYPE)
	viper.AddConfigPath(consts.TEST_ENV_FILE_DIRECTORY)
	viper.AddConfigPath(consts.ENV_FILE_DIRECTORY)

	err := viper.ReadInConfig()
	if err != nil {
		log.Debug().Err(err).
			Msg("Error occurred while reading env file, might fallback to OS env config")
	}
	viper.AutomaticEnv()
}

func GetEnvVar(name string) string {
	if !viper.IsSet(name) {
		log.Debug().Msgf("Environment variable %s is not set", name)
		return ""
	}
	value := viper.GetString(name)
	return value
}
