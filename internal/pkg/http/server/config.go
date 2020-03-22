package server

import "time"

type (
	// Config hold HTTP server configurations
	Config struct {
		Address           string
		Port              int
		ReadTimeout       time.Duration
		WriteTimeout      time.Duration
		ReadHeaderTimeout time.Duration
		ShutdownTimeout   time.Duration
		TLSCertFile       string
		TLSKeyFile        string
	}
)
