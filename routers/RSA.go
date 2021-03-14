package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/JavDomGom/gcat/models"
	"github.com/JavDomGom/gcat/resources"
)

func RSA(w http.ResponseWriter, r *http.Request) {

	// Prime number, secret.
	p, err := strconv.Atoi(r.URL.Query().Get("p"))
	if err != nil {
		http.NotFound(w, r)
		return
	}

	// Prime number, secret.
	q, err := strconv.Atoi(r.URL.Query().Get("q"))
	if err != nil {
		http.NotFound(w, r)
		return
	}

	// Private key part (d, n), secret.
	d, err := strconv.Atoi(r.URL.Query().Get("d"))
	if err != nil {
		http.NotFound(w, r)
		return
	}

	// Public key first part (e, n), public.
	e, err := strconv.Atoi(r.URL.Query().Get("e"))
	if err != nil {
		http.NotFound(w, r)
		return
	}

	// Public key second part (e, n), public.
	n := p * q

	// Secret message.
	M, err := strconv.Atoi(r.URL.Query().Get("msg"))
	if err != nil {
		http.NotFound(w, r)
		return
	}

	// 1. Bob encrypts message "M" and sends it to Alice as "C".
	C := resources.ModExp(int64(M), int64(e), int64(n))

	// 2. Alice receives "C", decrypts it and gets "decryptedC".
	decryptedC := resources.ModExp(C, int64(d), int64(n))

	var rsa = models.RSA{
		PlainText:  M,
		CypherText: C,
		DecryptedC: decryptedC,
	}

	w.Header().Set("context-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(rsa)
}
