package model

type (
	Config struct {
		Kafka  KafkaConfig  `mapstructure:"kafka"`
		DB     DBConfig     `mapstructure:"database"`
		Server ServerConfig `mapstructure:"server"`
	}

	KafkaConfig struct {
		Brokers  []string `mapstructure:"brokers"`
		Username string   `mapstructure:"username"`
		Password string   `mapstructure:"password"`
	}

	DBConfig struct {
		Host     string `mapstructure:"host"`
		Username string `mapstructure:"username"`
		Password string `mapstructure:"password"`
		Schema   string `mapstructure:"schema"`
	}

	ServerConfig struct {
		Port string `mapstructure:"port"`
	}
)
