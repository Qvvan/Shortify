package server

func StartServer() {
	// Инициализация роутов
	router := InitRoutes()
	router.Run()
}
