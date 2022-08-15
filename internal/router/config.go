package router

type Config struct {
	DefaultGateWay    string `koanf:"defaultGateWay"`
	NumberOfInterface int    `koanf:"numberOfInterface"`
}
