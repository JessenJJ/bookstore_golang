package main

import (
	"golang2bookst/internals/routes"
	"golang2bookst/pkg"
	"log"
)

// Depedency Injection (DI)

func main() {

	// inisialisasi DB
	_, err := pkg.InitMysql()
	if err != nil {

		log.Fatal(err)
		// os.Exit() atau log.Fatal(err) sama dengan return
	}
	// inisialisasi Router
	router := routes.InitRouter()
	// inisialisasi Server
	server := pkg.InitServer(router)
	// jalankan server

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

}
