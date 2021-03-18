package models

type EllipticCurvePointMultiplication struct {
	A         int64 `json:"a"`
	B         int64 `json:"b"`
	P         int64 `json:"p"`
	X_1       int64 `json:"x_1"`
	Y_1       int64 `json:"y_1"`
	Op_x_1    int64 `json:"op_x_1"`
	Inv_2y_1  int64 `json:"inv_2y_1"`
	Lambda    int64 `json:"lambda"`
	X         int64 `json:"x"`
	Op_lambda int64 `json:"op_lambda"`
	Y         int64 `json:"y"`
}
