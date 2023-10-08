package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Config struct {
	Listen           string        `envconfig:"LISTEN" default:":8080"`
	PostgresLogin    string        `envconfig:"POSTGRES_LOGIN" default:"user"`
	PostgresPassword string        `envconfig:"POSTGRES_PASSWORD" default:"super-secret-password"`
	PgConnectUrl     string        `envconfig:"PG_CONNECT_URL" default:"postgres://user:pass@localhost:5432/postgres"`
	Timeout          time.Duration `envconfig:"TIMEOUT" default:"10s"`
}

func (c Config) Print() {
	tmp := Config{
		Listen:           c.Listen,
		PostgresLogin:    c.PostgresLogin,
		PostgresPassword: "***", //c.PostgresPassword,
		PgConnectUrl:     c.PgConnectUrl,
		Timeout:          c.Timeout,
	}

	b, _ := json.MarshalIndent(tmp, "", "  ")

	fmt.Println(string(b))
}
