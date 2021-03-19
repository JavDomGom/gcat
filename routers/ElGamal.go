package routers

import (
	"encoding/json"
	"math"
	"net/http"
	"strconv"

	"github.com/JavDomGom/gcat/models"
	"github.com/JavDomGom/gcat/resources"
)

func ElGamal(w http.ResponseWriter, r *http.Request) {

	// Bob's public number.
	p, err := strconv.ParseInt(r.URL.Query().Get("p"), 10, 64)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	// Bob's public number.
	a, err := strconv.ParseInt(r.URL.Query().Get("a"), 10, 64)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	// Bob's private key.
	priK, err := strconv.ParseInt(r.URL.Query().Get("pri_k"), 10, 64)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	// Bob's public key.
	pubK := resources.ModExp(a, priK, p)

	// Secret number.
	N, err := strconv.ParseInt(r.URL.Query().Get("n"), 10, 64)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	// Alice's random session number.
	v, err := strconv.ParseInt(r.URL.Query().Get("v"), 10, 64)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	// Alice's first part of pair (N1, N2).
	N1 := resources.ModExp(a, v, p)

	// Alice's second part of pair (N1, N2).
	N2 := int64(
		math.Mod(
			math.Mod(
				math.Mod(float64(N), float64(p))*float64(resources.ModExp(pubK, v, p)),
				float64(p),
			),
			float64(p),
		),
	)

	// Alice send to Bob the pair (N1, N2), and Bob calculates:
	N3 := resources.ModExp(N1, priK, p)
	NDecrypted := int64(
		math.Mod(
			float64(N2*resources.ModInv(N3, p)),
			float64(p),
		),
	)

	var elgamal = models.ElGamal{
		Bob_p:        p,
		Bob_a:        a,
		Bob_PriK:     priK,
		Bob_PubK:     pubK,
		SecretNumber: N,
		Alice_v:      v,
		Alice_N1:     N1,
		Alice_N2:     N2,
		Bob_N3:       N3,
		N_decrypted:  NDecrypted,
	}

	w.Header().Set("context-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(elgamal)
}
