package routers

import (
	"encoding/json"
	"math"
	"net/http"
	"strconv"

	"github.com/JavDomGom/gcat/models"
	"github.com/JavDomGom/gcat/resources"
)

// EllipticCurvePointAddition for elliptic curve: y² = x³ + ax + b mod p
func EllipticCurvePointAddition(w http.ResponseWriter, r *http.Request) {

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
	x_2, err := strconv.ParseInt(r.URL.Query().Get("x_2"), 10, 64)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	y_2, err := strconv.ParseInt(r.URL.Query().Get("y_2"), 10, 64)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	op_x := x_1 - x_2
	op_y := y_1 - y_2

	inv_op_x := resources.ModInv(op_x, p)

	lambda := resources.Mod(op_y*inv_op_x, p)

	x := resources.Mod(int64(math.Pow(float64(lambda), 2))-x_1-x_2, p)

	y := resources.Mod(lambda*(x_1-x)-y_1, p)

	var ecpa = models.EllipticCurvePointAddition{
		A:        a,
		B:        b,
		P:        p,
		X_1:      x_1,
		Y_1:      y_1,
		X_2:      x_2,
		Y_2:      y_2,
		Op_x:     op_x,
		Op_y:     op_y,
		Inv_op_x: inv_op_x,
		Lambda:   lambda,
		X:        x,
		Y:        y,
	}

	w.Header().Set("context-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(ecpa)
}
