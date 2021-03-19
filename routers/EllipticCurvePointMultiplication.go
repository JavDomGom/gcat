package routers

import (
	"encoding/json"
	"math"
	"net/http"
	"strconv"

	"github.com/JavDomGom/gcat/models"
	"github.com/JavDomGom/gcat/resources"
)

// EllipticCurvePointMultiplication for elliptic curve: y² = x³ + ax + b mod p
func EllipticCurvePointMultiplication(w http.ResponseWriter, r *http.Request) {

	// Parameters.
	a, err := strconv.ParseInt(r.URL.Query().Get("a"), 10, 64)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	b, err := strconv.ParseInt(r.URL.Query().Get("b"), 10, 64)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	p, err := strconv.ParseInt(r.URL.Query().Get("p"), 10, 64)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	x_1, err := strconv.ParseInt(r.URL.Query().Get("x_1"), 10, 64)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	y_1, err := strconv.ParseInt(r.URL.Query().Get("y_1"), 10, 64)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	op_x_1 := 3*int64(math.Pow(float64(x_1), 2)) + 1
	inv_2y_1 := resources.ModInv(2*y_1, p)

	lambda := resources.Mod(op_x_1*inv_2y_1, p)

	x := resources.Mod(int64(math.Pow(float64(lambda), 2))-2*x_1, p)

	op_lambda := lambda * (x_1 - x)

	y := resources.Mod(op_lambda-y_1, p)

	var ecpm = models.EllipticCurvePointMultiplication{
		A:         a,
		B:         b,
		P:         p,
		X_1:       x_1,
		Y_1:       y_1,
		Op_x_1:    op_x_1,
		Inv_2y_1:  inv_2y_1,
		Lambda:    lambda,
		X:         x,
		Op_lambda: op_lambda,
		Y:         y,
	}

	w.Header().Set("context-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(ecpm)
}
