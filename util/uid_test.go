package util

import "testing"

func TestGenerateOrderSN(t *testing.T) {
	// 240708214434867543453543
	println(GenerateOrderSN(23))
}

func TestGenerateRefundSN(t *testing.T) {

	println(GenerateRefundSN(1, 1))
}

func TestGenerateCode(t *testing.T) {
	println(GenerateCode())
}
