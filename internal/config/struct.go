package config

type Config struct {
	//LogLevel is the log level used by Slog
	LogLevel string `envconfig:"LOG_LEVEL" default:"info"`

	//LogPath is the path of the app's log file
	LogPath string `envconfig:"LOG_PATH" default:"application.log"`

	//DBConfig is the configuration for the database
	DBConfig DBConfig
}

type DBConfig struct {
	//DBHost is the host of the database
	DBUrl string `envconfig:"DB_URL" required:"true"`
}
