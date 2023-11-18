package main

import (
	"fmt"
	"math/big"
)

type bigNumbers struct {
	a *big.Int
	b *big.Int
}

func (bn *bigNumbers) add() *big.Int {
	ans := new(big.Int)
	return ans.Add(bn.a, bn.b)
}

func (bn *bigNumbers) sub() *big.Int {
	ans := new(big.Int)
	return ans.Sub(bn.a, bn.b)
}

func (bn *bigNumbers) mul() *big.Int {
	ans := new(big.Int)
	return ans.Mul(bn.a, bn.b)
}

func (bn *bigNumbers) div() *big.Int {
	ans := new(big.Int)
	return ans.Div(bn.a, bn.b)
}

func main() {
	a := big.NewInt(16e9)
	b := big.NewInt(12e8)

	bn := bigNumbers{a, b}

	fmt.Println(bn.add())
	fmt.Println(bn.sub())
	fmt.Println(bn.mul())
	fmt.Println(bn.div())
}
