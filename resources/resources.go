package resources

import (
	"math/big"
)

func ModExp(base, exponent, modulus int64) int64 {
	return new(big.Int).Exp(
		big.NewInt(base),
		big.NewInt(exponent),
		big.NewInt(modulus),
	).Int64()
}

func ModInv(a, m int64) int64 {
	return new(big.Int).ModInverse(
		big.NewInt(a),
		big.NewInt(m),
	).Int64()
}

func Mod(a, p int64) int64 {
	m := a % p
	if a < 0 && p < 0 {
		m -= p
	}
	if a < 0 && p > 0 {
		m += p
	}
	return m
}
