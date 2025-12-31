package nibirudb

import "fmt"

type Config struct {
	Driver                string
	Host                  string
	User                  string
	Password              string
	Name                  string
	Port                  int
	MaxConnectionLifeTime string
	MaxConnectionIdleTime string
	MaxIdleConnections    int
	MaxOpenConnections    int
	DebugMode             bool
	Timeout               string
	WriteTimeout          string
	ReadTimeout           string
	SSLMode               string
	Timezone              string
}

func (cfg *Config) GetDSN() string {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?timeout=%s&writeTimeout=%s&readTimeout=%s&charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Name,
		cfg.Timeout,
		cfg.WriteTimeout,
		cfg.ReadTimeout,
	)

	if cfg.Driver == "postgres" {
		dsn = fmt.Sprintf("%s://%s:%s@%s:%d/%s?sslmode=%s",
			cfg.Driver,
			cfg.User,
			cfg.Password,
			cfg.Host,
			cfg.Port,
			cfg.Name,
			cfg.SSLMode,
		)
	}

	return dsn
}
