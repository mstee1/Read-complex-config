package config

import (
	"flag"
	"os"
	"strings"

	"github.com/spf13/viper"
)

func GetConfig() (*Config, error) {
	var v = viper.New()
	var config Config
	var absoluteConfigPath, configPath, configName, configType string
	args := os.Args
	for _, arg := range args {
		if arg == "--version" {
			config.Version = "1.0.0"
			return &config, nil
		}
		if strings.Split(arg, "=")[0][1:] == "configPath" {
			absoluteConfigPath = strings.Split(arg, "=")[1]
			configPath, configName, configType = getParamsConf(absoluteConfigPath)
			break
		}
	}

	if absoluteConfigPath == "" {
		configPath = "./"
		configName = "config"
		configType = "yaml"
	}

	err := readParametersFromConfig(*v, configPath, configName, configType, &config)
	if err != nil {
		return nil, err
	}
	readFlags(&config)
	return &config, nil
}

func getParamsConf(absoluteConfigPath string) (string, string, string) {

	pathSplit := strings.Split(absoluteConfigPath, "/")
	configNameType := strings.Split(pathSplit[len(pathSplit)-1], ".")
	configPath := strings.Join(pathSplit[:len(pathSplit)-1], "/") + "/"

	return configPath, configNameType[0], configNameType[1]

}

func readParametersFromConfig(viper viper.Viper, configPath, configName, configType string, cfg *Config) error {

	viper.SetConfigName(configName)
	viper.SetConfigType(configType)
	viper.AddConfigPath(configPath)

	// Попытка чтения конфига
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	// Попытка заполнение структуры Config полученными данными
	if err := viper.Unmarshal(&cfg); err != nil {
		return err
	}
	return nil
}

func readFlags(cfg *Config) {
	var configPath string

	flag.StringVar(&cfg.Logger.LogLevel, "logLevel", cfg.Logger.LogLevel, "Aviable LogLevael: INFO,DEBUG,TEST")
	flag.StringVar(&cfg.Logger.LogDir, "logDir", cfg.Logger.LogDir, "Full path to save log file")
	flag.StringVar(&cfg.Logger.LogMode, "logMode", cfg.Logger.LogMode, "Aviable LogMode:stdout,file,empty field")
	flag.BoolVar(&cfg.Logger.RewriteLog, "rewriteLog", cfg.Logger.RewriteLog, "Overwriting a log file")

	flag.StringVar(&configPath, "configPath", "./", "The configPath parameter")
	flag.Parse()
}
