package config

type Config struct {
	//LogLevel is the log level used by Slog
	LogLevel string `envconfig:"LOG_LEVEL" default:"info"`

	//LogPath is the path of the app's log file
	LogPath string `envconfig:"LOG_PATH" default:"application.log"`

	//DatabaseUrl is the connection string of the database
	DatabaseURL string `envconfig:"DATABASE_URL" required:"true"`
}
