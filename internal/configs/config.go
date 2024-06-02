package configs

import "flag"

type Config struct {
	URL         string
	Requests    int
	Concurrency int
}

func LoadConfig() *Config {
	cfg := &Config{}

	flag.StringVar(&cfg.URL, "url", "", "URL of the service to test")
	flag.IntVar(&cfg.Requests, "requests", 100, "Number of total requests to make")
	flag.IntVar(&cfg.Concurrency, "concurrency", 10, "Number of concurrent (simultaneous) requests to make")
	flag.Parse()

	return cfg
}
