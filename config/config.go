package config

import "github.com/spf13/viper"

type Config struct {
	DBDriver           string `mapstructure:"DB_DRIVER"`
	DBSource           string `mapstructure:"DB_SOURCE"`
	CacheSource        string `mapstructure:"CACHE_SOURCE"`
	KafkaSource        string `mapstructure:"KAFKA_SOURCE"`
	NewRelicAppName    string `mapstructure:"NEW_RELIC_APP_NAME"`
	NewRelicLicenseKey string `mapstructure:"NEW_RELIC_LICENSE_KEY"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
