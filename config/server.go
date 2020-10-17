package config

// Server holds http server related config
// read and write timeout are in seconds
type ServerConfig struct {
	Port int
	ReadTimeout int
	WriteTimeout int
}

func newServerConfig() ServerConfig {
	return ServerConfig{
		Port: mustGetInt("SERVER_PORT"),
		ReadTimeout: mustGetInt("SERVER_READ_TIMEOUT"),
		WriteTimeout: mustGetInt("SERVER_WRITE_TIMEOUT"),
	}
}
