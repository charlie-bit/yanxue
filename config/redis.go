package config

// Redis config
type Redis struct {
	Addr     string `yaml:"addr"`
	DB       int    `yaml:"db"`
	Username string `yaml:"user_name"`
	Password string `yaml:"password"`
}
