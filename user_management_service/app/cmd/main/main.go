package main

import (
	"flag"
)

const version = "1.0.0"

type config struct {
	port int
	env  string
}

type application struct {
	config config
	logger string
}

func main() {
	var cfg config
	var logsUrl string

	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&logsUrl, "logsUrl", ":localhost:9090", "url of log service")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")

	flag.Parse()
}
