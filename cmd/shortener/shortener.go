package main

import "github.com/nikitachukov/url_shortener.git/internal/app"

func main() {
	app.StartServer()
}

//post / работает как есть, но возвращает куку,
//post /api/shorten работает как есть, но тоже возвращает куку,
//get /api/user/urls возвращает 401 при отсутствии auth info и взвращает куку,
//                              204 при отсутствии куки и возвращает куку,
//                              200 при наличии данных и куки, от себя,
//                              204 при наличии куки и отсутствии данных
