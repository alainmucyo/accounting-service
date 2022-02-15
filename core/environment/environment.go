package environment

type Environment struct {
	KafkaBroker   string
	KafkaGroupId  string
	Port          string
	RedisURL      string
	RedisPassword string
	DBURL         string
}

func New(
	port string,
	redisURL string,
	redisPassword string,
	dbURL string,
	kafkaBroker string,
	kafkaGroupId string,
) *Environment {
	return &Environment{
		Port:          port,
		RedisURL:      redisURL,
		RedisPassword: redisPassword,
		DBURL:         dbURL,
		KafkaGroupId:  kafkaGroupId,
		KafkaBroker:   kafkaBroker,
	}
}
