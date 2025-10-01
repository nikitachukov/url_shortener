package config

import (
	"flag"
)

var (
	AppAddr  *string
	BasePath *string
)

func ParseParams() {
	AppAddr = flag.String("a", "localhost:8080", "адрес запуска HTTP-сервера")
	BasePath = flag.String("b", "", "базовый адрес результирующего сокращённого URL")
	flag.Parse()
}
