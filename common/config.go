package common

import "github.com/spf13/viper"

type selfConfig struct {
	Host     string `yaml:"host"`
	Port     uint16 `yaml:"port"`
	Language string `yaml:"language"`
}

type dependencyConfig struct {
	AuthorizationEndpoint  string `yaml:"authorization_endpoint"`
	UserManagementEndpoint string `yaml:"user_management_endpoint"`
}

var (
	SelfConfig       selfConfig
	DependencyConfig dependencyConfig
)

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := viper.UnmarshalKey("self", &SelfConfig); err != nil {
		panic(err)
	}
	if err := viper.UnmarshalKey("dependency", &DependencyConfig); err != nil {
		panic(err)
	}
}
