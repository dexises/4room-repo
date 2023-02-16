package main

type config struct {
	port int
	env  string
	db   struct {
		dsn          string
		maxOpenConns int
		maxIdleConns int
		maxIdleTime  string
	}
}

func (c *config) GetInstanceConfig(dsn string, maxOpenCon, maxidleCon int, maxidleTime string) *config {
}
