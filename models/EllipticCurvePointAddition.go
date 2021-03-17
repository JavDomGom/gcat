package models

type EllipticCurvePointAddition struct {
	A        int64 `json:"a"`
	B        int64 `json:"b"`
	P        int64 `json:"p"`
	X_1      int64 `json:"x_1"`
	Y_1      int64 `json:"y_1"`
	X_2      int64 `json:"x_2"`
	Y_2      int64 `json:"y_2"`
	Op_x     int64 `json:"op_x"`
	Op_y     int64 `json:"op_y"`
	Inv_op_x int64 `json:"inv_op_x"`
	Lambda   int64 `json:"lambda"`
	X        int64 `json:"x"`
	Y        int64 `json:"y"`
}
