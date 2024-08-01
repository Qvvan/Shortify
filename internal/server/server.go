package server

// ⚙️ Оптимизация серверной логики и маршрутизации
func StartServer() {
	router := InitRoutes()
	router.Run()
}
