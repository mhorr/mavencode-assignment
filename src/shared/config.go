package shared

import "time"

// Config holds configuration values for the program.
type Config struct {
	Redis           string // Connection string for redis
	Rabbit          string // Connection string for RabbitMQ
	TimestampFormat string // Format to use when parsing / marshalling timestamp values
	PersonChannel   string // Name of channel for storing person objects
	RedisTTL        int
}

// GetConfig gets the current configuration. For now, it's hardcoded.
// For future work, add code to read environment variables or config file.
func GetConfig() Config {
	return Config{
		Redis:           "localhost:6379",                     // redis connection info
		Rabbit:          "amqp://guest:guest@localhost:5672/", // URL for accessing rabbit server
		TimestampFormat: time.RFC3339,                         // TODO: maybe get rid of this.
		PersonChannel:   "person.store",                       // name of RabbitMQ channel for storing person objects
		RedisTTL:        3600,                                 // TTL for keys inserted into redis
	}
}
