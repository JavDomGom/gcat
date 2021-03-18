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
	router.HandleFunc("/aes-256-cbc", routers.AES256CBC).Methods("GET")
	router.HandleFunc("/aes-256-ctr", routers.AES256CTR).Methods("GET")
	router.HandleFunc("/rsa", routers.RSA).Methods("GET")
	router.HandleFunc("/elgamal", routers.ElGamal).Methods("GET")
	router.HandleFunc("/ellipticcurves/pointaddition", routers.EllipticCurvePointAddition).Methods("GET")
	router.HandleFunc("/ellipticcurves/pointmultiplication", routers.EllipticCurvePointMultiplication).Methods("GET")

	port := os.Getenv("GCAT_PORT")
	if port == "" {
		port = "8080"
	}
	handler := cors.AllowAll().Handler(router)

	log.Println("Starting server on http://localhost:" + port)
	log.Fatal(http.ListenAndServe(":"+port, handler))
}
