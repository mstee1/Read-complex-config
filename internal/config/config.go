package config

type Config struct {
	Logger  Logger   `mapstructure:"logger"`
	Workers []Worker `mapstructure:"workers"`
	Version string
}

type Logger struct {
	LogLevel   string `mapstructure:"logLevel"`
	LogDir     string `mapstructure:"logDir"`
	LogMode    string `mapstructure:"logMode"`
	RewriteLog bool   `mapstructure:"rewriteLog"`
}

type Worker struct {
	Use      bool     `mapstructure:"use"`
	Database Database `mapstructure:"database"`
	Sql      Sql      `mapstructure:"sql"`
}

type Database struct {
	DbPort      string `mapstructure:"dbPort"`
	DbHost      string `mapstructure:"dbHost"`
	DbName      string `mapstructure:"dbName"`
	DbUser      string `mapstructure:"dbUser"`
	DbPass      string `mapstructure:"dbPassword"`
	UseContract bool   `mapstructure:"useContract"`
}

type Sql struct {
	Select1 string `mapstructure:"select1"`
	Update1 string `mapstructure:"update1"`
}
