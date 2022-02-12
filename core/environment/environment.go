package environment

type Environment struct {
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
) *Environment {
	return &Environment{
		Port:          port,
		RedisURL:      redisURL,
		RedisPassword: redisPassword,
		DBURL:         dbURL,
	}
}
