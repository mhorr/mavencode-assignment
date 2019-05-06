package shared

import "github.com/spf13/viper"

func init() {
	viper.SetEnvPrefix("MW")
	viper.BindEnv("Redis")
	viper.BindEnv("Rabbit")
	viper.BindEnv("PersonChannel")
	viper.BindEnv("RedisTTL")
	viper.SetDefault("Redis", ":6379")
	viper.SetDefault("Rabbit", "amqp://guest:guest@rabbit:5672")
	viper.SetDefault("PersonChannel", "person.store")
	viper.SetDefault("RedisTTL", 3600)
}

// Config holds configuration values for the program.
type Config struct {
	Redis         string // Connection string for redis
	Rabbit        string // Connection string for RabbitMQ
	PersonChannel string // Name of channel for storing person objects
	RedisTTL      int
}

// GetConfig gets the current configuration. For now, it's hardcoded.
// For future work, add code to read environment variables or config file.
func GetConfig() Config {

	return Config{
		Redis:         viper.GetString("Redis"),         // redis connection info
		Rabbit:        viper.GetString("Rabbit"),        // URL for accessing rabbit server
		PersonChannel: viper.GetString("PersonChannel"), // name of RabbitMQ channel for storing person objects
		RedisTTL:      viper.GetInt("RedisTTL"),         // TTL for keys inserted into redis
	}
}
