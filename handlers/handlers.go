package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/JavDomGom/gcat/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func Handlers() {
	router := mux.NewRouter()

	router.HandleFunc("/md5", routers.Hash).Methods("GET")
	router.HandleFunc("/sha1", routers.Hash).Methods("GET")
	router.HandleFunc("/sha256", routers.Hash).Methods("GET")
	router.HandleFunc("/sha512", routers.Hash).Methods("GET")
	router.HandleFunc("/fnv", routers.Hash).Methods("GET")
	router.HandleFunc("/adler32", routers.Hash).Methods("GET")
	router.HandleFunc("/crc32", routers.Hash).Methods("GET")
	router.HandleFunc("/crc64", routers.Hash).Methods("GET")

	GCAT_PORT := os.Getenv("GCAT_PORT")
	if GCAT_PORT == "" {
		GCAT_PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)

	log.Println("Starting server on http://localhost:" + GCAT_PORT)
	log.Fatal(http.ListenAndServe(":"+GCAT_PORT, handler))
}
