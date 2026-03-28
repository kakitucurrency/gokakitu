package util

import (
	"errors"
	"math/big"
)

// KakituAmount wraps a raw amount in KSHS (Kakitu).
type KakituAmount struct {
	Raw *big.Int
}

func (KakituAmount) exp() *big.Int {
	x := big.NewInt(10)
	return x.Exp(x, big.NewInt(30), nil)
}

// KakituAmountFromString parses KSHS amounts in strings.
func KakituAmountFromString(s string) (n KakituAmount, err error) {
	r, ok := new(big.Rat).SetString(s)
	if !ok {
		err = errors.New("unable to parse kakitu amount")
		return
	}
	r = r.Mul(r, new(big.Rat).SetInt(n.exp()))
	if !r.IsInt() {
		err = errors.New("unable to parse kakitu amount")
		return
	}
	n.Raw = r.Num()
	return
}

func (n KakituAmount) String() string {
	r := new(big.Rat).SetFrac(n.Raw, n.exp())
	s := r.FloatString(30)
	return s[:len(s)-24]
}
