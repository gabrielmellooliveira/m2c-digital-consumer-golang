package configs

import "github.com/spf13/viper"

type conf struct {
	M2CDigitalApiUrl string `mapstructure:"M2C_DIGITAL_API_URL"`
	M2CDigitalApiKey string `mapstructure:"M2C_DIGITAL_API_KEY"`
	MongoDBUrl       string `mapstructure:"MONGODB_URL"`
	RabbitMQUrl      string `mapstructure:"RABBITMQ_URL"`
	RedisUrl         string `mapstructure:"REDIS_URL"`
}

func LoadConfig() (*conf, error) {
	var cfg *conf

	viper.SetConfigName("cmd/.env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
