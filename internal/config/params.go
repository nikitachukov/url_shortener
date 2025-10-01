package config

import (
	"flag"
	"os"
)

type Params struct {
	AppAddr  *string
	BasePath *string
}

func NewParams() *Params {
	return &Params{}
}

func (p *Params) InitParams() {
	pFlagAppAddr := flag.String("a", "localhost:8080", "адрес запуска HTTP-сервера")
	pFlagBasePath := flag.String("b", "", "базовый адрес результирующего сокращённого URL")
	flag.Parse()

	AppAddr := os.Getenv("SERVER_ADDRESS")
	if AppAddr == "" {
		if *pFlagAppAddr != "" {
			AppAddr = *pFlagAppAddr
		}
	}

	BasePath := os.Getenv("BASE_URL")
	if BasePath == "" {
		if *pFlagBasePath != "" {
			BasePath = *pFlagBasePath
		}
	}

	p.AppAddr = &AppAddr
	p.BasePath = &BasePath

	//Если указана переменная окружения, то используется она.
	//Если нет переменной окружения, но есть аргумент командной строки (флаг), то используется он.
	//Если нет ни переменной окружения, ни флага, то используется значение по умолчанию.

}
