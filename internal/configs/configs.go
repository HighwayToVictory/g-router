package configs

type Config struct {
}

func Load() Config {
	return Default()
}
