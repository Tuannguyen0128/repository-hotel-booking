package model

type (
	Config struct {
		DB               DBConfig     `mapstructure:"database"`
		Server           ServerConfig `mapstructure:"server"`
		KafkaServer      string       `mapstructure:"kafka_server"`
		KafkaBrokerTopic string       `mapstructure:"kafka_broker-topic"`
		KafkaRepoTopic   string       `mapstructure:"kafka_repo-topic"`
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
