package config

type Configuration struct {
	Database Database `yaml:"Database"`
}

type Database struct {
	Host   string `yaml:"Host"`
	Port   int    `yaml:"Port"`
	User   string `yaml:"User"`
	Passwd string `yaml:"Passwd"`
	Db     string `yaml:"Db"`
}
