package config

import "github.com/spf13/viper"

type Config struct {
	Port          string `mapstructure:"PORT"`
	DBHost        string `mapstructure:"DBHost"`
	DBPort        string `mapstructure:"DBPort"`
	DBUser        string `mapstructure:"DBUser"`
	DBPassword    string `mapstructure:"DBPassword"`
	DBName        string `mapstructure:"DBName"`
	ProductSvcUrl string `mapstructure:"PRODUCT_SVC_URL"`
}

func LoadConfig() (confg Config, err error) {
	viper.AddConfigPath("./pkg/config/envs")
	viper.SetConfigName("dev")
	viper.SetConfigType("env")

	viper.SetDefault("PORT", ":50053")
	viper.SetDefault("DBHost", "localhost")
	viper.SetDefault("DBPort", "5432")
	viper.SetDefault("DBUser", "postgres")
	viper.SetDefault("DBPassword", "7356")
	viper.SetDefault("DBName", "order_svc")
	viper.SetDefault("PRODUCT_SVC_URL", "localhost:50052")

	viper.AutomaticEnv()

	if err = viper.ReadInConfig(); err != nil {
		return
	}

	err = viper.Unmarshal(&confg)

	return
}
