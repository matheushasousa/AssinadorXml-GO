package rabbit

type Config struct {
	Url string
}

func LoadConfig() Config {

	return Config{
		Url: "amqp://guest:guest@localhost:5672/",
	}
}
