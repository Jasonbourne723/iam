package config

type Configuration struct {
	Database Database `yaml:"Database"`
}

type Database struct {
	Host string `yaml:"host"`
}
