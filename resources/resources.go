package resources

import "math/big"

func ModExp(base, exponent, modulus int64) int64 {
	return new(big.Int).Mod(new(big.Int).Exp(big.NewInt(base), big.NewInt(exponent), nil), big.NewInt(modulus)).Int64()
}
