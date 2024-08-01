package server

func StartServer() {
	router := InitRoutes()
	router.Run()
}
