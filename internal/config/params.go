package config

import (
	"flag"
	"os"
)

type Params struct {
	AppAddr         *string
	BasePath        *string
	FileStoragePath *string
}

func NewParams() *Params {
	return &Params{}
}

func (p *Params) InitParams() {
	pFlagAppAddr := flag.String("a", "localhost:8080", "адрес запуска HTTP-сервера")
	pFlagBasePath := flag.String("b", "", "базовый адрес результирующего сокращённого URL")
	pFileStoragePath := flag.String("f", "", "путь до файла, куда сохраняются данные")
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

	FileStoragePath := os.Getenv("FILE_STORAGE_PATH")
	if FileStoragePath == "" {
		if *pFileStoragePath != "" {
			FileStoragePath = *pFileStoragePath
		}
	}

	p.AppAddr = &AppAddr
	p.BasePath = &BasePath
	p.FileStoragePath = &FileStoragePath

	//Если указана переменная окружения, то используется она.
	//Если нет переменной окружения, но есть аргумент командной строки (флаг), то используется он.
	//Если нет ни переменной окружения, ни флага, то используется значение по умолчанию.

}
