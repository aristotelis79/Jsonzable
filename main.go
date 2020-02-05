package main

import (
	"log"
	"net/http"

	"github.com/aristotelis79/Jsonzable/routers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)

	// statikFS, err := fs.New()
	// if err != nil {
	// 	panic(err)
	// }

	// staticServer := http.FileServer(statikFS)
	// sh := http.StripPrefix("/swaggerui/", staticServer)
	// router.PathPrefix("/swaggerui/").Handler(sh)

	routers.RegisterV1Routes(router)

	log.Fatal(http.ListenAndServe(":8080", router))
}
