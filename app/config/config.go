package config

type DatabaseDetail struct {
	URL      string `yaml:"url"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type Config struct {
	InsertCount    int            `yaml:"insertCount"`
	DatabaseConfig DatabaseDetail `yaml:"databaseConfig"`
}
