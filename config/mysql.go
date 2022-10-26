package config

// Mysql config
type Mysql struct {
	DSN          string `yaml:"dsn"`
	MaxOpenConns int    `yaml:"max_open_conns"`
	MaxIdleConns int    `yaml:"max_idle_conns"`
	Path         string `yaml:"path"`
	FromVersion  int    `yaml:"from_version"`
	ToVersion    uint   `yaml:"to_version"`
}
