package server

import "time"

type (
	// Config hold HTTP server configurations
	Config struct {
		Address           string        `mapstructure:"address"`
		Port              int           `mapstructure:"port"`
		ReadTimeout       time.Duration `mapstructure:"read_timeout"`
		WriteTimeout      time.Duration `mapstructure:"write_timeout"`
		ReadHeaderTimeout time.Duration `mapstructure:"read_header_timeout"`
		ShutdownTimeout   time.Duration `mapstructure:"shutdown_timeout"`
		TLSCertFile       string        `mapstructure:"tls_cert_file"`
		TLSKeyFile        string        `mapstructure:"tls_key_file"`
	}
)
