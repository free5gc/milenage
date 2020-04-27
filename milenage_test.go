package milenage_test

import (
	"encoding/hex"
	"fmt"
	"free5gc/lib/milenage"
	"testing"
)

func TestGenerateOPC(t *testing.T) {
	_, _, OPC := make([]byte, 16), make([]byte, 16), make([]byte, 16)

	// OPC_str := ""
	// OPC, _ = hex.DecodeString(OPC_str)
	// fmt.Println("OPC", OPC)

	K_str := "3016ebeae2c45bd0060923dbbb402be6"
	K, _ := hex.DecodeString(K_str)

	OP_str := "00000000000000000000000000000000"
	OP, _ := hex.DecodeString(OP_str)
	fmt.Println("K:", K)

	fmt.Println("OP:", OP)

	milenage.GenerateOPC(K, OP, OPC)
}
