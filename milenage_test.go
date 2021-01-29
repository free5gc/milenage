package milenage_test

import (
	"encoding/hex"
	"fmt"
	// "math/rand"
	"reflect"
	"strings"
	"testing"

	"github.com/free5gc/milenage"
)

type f1Test struct {
	K           string
	RAND        string
	SQN         string
	AMF         string
	OP          string
	ExpectedOPc string
	f1          string
	f1Start     string
}

type f2f5f3Test struct {
	K           string
	RAND        string
	OP          string
	ExpectedOPc string
	ExpectedRES string
	ExpectedAK  string
	ExpectedCK  string
}

type f4f5StarTest struct {
	K              string
	RAND           string
	OP             string
	ExpectedOPc    string
	ExpectedIK     string
	ExpectedAKStar string
}

const (
	f1Test1       = "f1Test1"
	f1Test2       = "f1Test2"
	f1Test3       = "f1Test3"
	f1Test4       = "f1Test4"
	f1Test5       = "f1Test5"
	f1Test6       = "f1Test6"
	f2f5f3Test1   = "f2f5f3Test1"
	f2f5f3Test2   = "f2f5f3Test2"
	f2f5f3Test3   = "f2f5f3Test3"
	f2f5f3Test4   = "f2f5f3Test4"
	f2f5f3Test5   = "f2f5f3Test5"
	f4f5StarTest1 = "f4f5StarTest1"
	f4f5StarTest2 = "f4f5StarTest2"
	f4f5StarTest3 = "f4f5StarTest3"
	f4f5StarTest4 = "f4f5StarTest4"
	f4f5StarTest5 = "f4f5StarTest5"
	f4f5StarTest6 = "f4f5StarTest6"
)

var Testf1TestTable = make(map[string]*f1Test)
var Testf2f5f3TestTable = make(map[string]*f2f5f3Test)
var Testf4f5StarTestTable = make(map[string]*f4f5StarTest)

func init() {
	Testf1TestTable[f1Test1] = &f1Test{
		K:           "465b5ce8b199b49faa5f0a2ee238a6bc",
		RAND:        "23553cbe9637a89d218ae64dae47bf35",
		SQN:         "ff9bb4d0b607",
		AMF:         "b9b9",
		OP:          "cdc202d5123e20f62b6d676ac72cb318",
		ExpectedOPc: "cd63cb71954a9f4e48a5994e37a02baf",
		f1:          "4a9ffac354dfafb3",
		f1Start:     "01cfaf9ec4e871e9",
	}
	Testf1TestTable[f1Test2] = &f1Test{
		K:           "0396eb317b6d1c36f19c1c84cd6ffd16",
		RAND:        "c00d603103dcee52c4478119494202e8",
		SQN:         "fd8eef40df7d",
		AMF:         "af17",
		OP:          "ff53bade17df5d4e793073ce9d7579fa",
		ExpectedOPc: "53c15671c60a4b731c55b4a441c0bde2",
		f1:          "5df5b31807e258b0",
		f1Start:     "a8c016e51ef4a343",
	}

	Testf1TestTable[f1Test3] = &f1Test{
		K:           "fec86ba6eb707ed08905757b1bb44b8f",
		RAND:        "9f7c8d021accf4db213ccff0c7f71a6a",
		SQN:         "9d0277595ffc",
		AMF:         "725c",
		OP:          "dbc59adcb6f9a0ef735477b7fadf8374",
		ExpectedOPc: "1006020f0a478bf6b699f15c062e42b3",
		f1:          "9cabc3e99baf7281",
		f1Start:     "95814ba2b3044324",
	}

	Testf1TestTable[f1Test4] = &f1Test{
		K:           "9e5944aea94b81165c82fbf9f32db751",
		RAND:        "ce83dbc54ac0274a157c17f80d017bd6",
		SQN:         "0b604a81eca8",
		AMF:         "9e09",
		OP:          "223014c5806694c007ca1eeef57f004f",
		ExpectedOPc: "a64a507ae1a2a98bb88eb4210135dc87",
		f1:          "74a58220cba84c49",
		f1Start:     "ac2cc74a96871837",
	}

	Testf1TestTable[f1Test5] = &f1Test{
		K:           "4ab1deb05ca6ceb051fc98e77d026a84",
		RAND:        "74b0cd6031a1c8339b2b6ce2b8c4a186",
		SQN:         "e880a1b580b6 ",
		AMF:         "9f07",
		OP:          "2d16c5cd1fdf6b22383584e3bef2a8d8",
		ExpectedOPc: "dcf07cbd51855290b92a07a9891e523e",
		f1:          "49e785dd12626ef2",
		f1Start:     "9e85790336bb3fa2",
	}

	Testf1TestTable[f1Test6] = &f1Test{
		K:           "6c38a116ac280c454f59332ee35c8c4f",
		RAND:        "ee6466bc96202c5a557abbeff8babf63",
		SQN:         "414b98222181",
		AMF:         "4464",
		OP:          "1ba00a1a7c6700ac8c3ff3e96ad08725",
		ExpectedOPc: "3803ef5363b947c6aaa225e58fae3934",
		f1:          "078adfb488241a57",
		f1Start:     "80246b8d0186bcf1",
	}

	Testf2f5f3TestTable[f2f5f3Test1] = &f2f5f3Test{
		K:           "465b5ce8b199b49faa5f0a2ee238a6bc",
		RAND:        "23553cbe9637a89d218ae64dae47bf35",
		OP:          "cdc202d5123e20f62b6d676ac72cb318",
		ExpectedOPc: "cd63cb71954a9f4e48a5994e37a02baf",
		ExpectedRES: "a54211d5e3ba50bf",
		ExpectedAK:  "aa689c648370",
		ExpectedCK:  "b40ba9a3c58b2a05bbf0d987b21bf8cb",
	}

	Testf2f5f3TestTable[f2f5f3Test2] = &f2f5f3Test{
		K:           "0396eb317b6d1c36f19c1c84cd6ffd16",
		RAND:        "c00d603103dcee52c4478119494202e8",
		OP:          "ff53bade17df5d4e793073ce9d7579fa",
		ExpectedOPc: "53c15671c60a4b731c55b4a441c0bde2",
		ExpectedRES: "d3a628ed988620f0",
		ExpectedAK:  "c47783995f72",
		ExpectedCK:  "58c433ff7a7082acd424220f2b67c556",
	}

	Testf2f5f3TestTable[f2f5f3Test3] = &f2f5f3Test{
		K:           "fec86ba6eb707ed08905757b1bb44b8f",
		RAND:        "9f7c8d021accf4db213ccff0c7f71a6a",
		OP:          "dbc59adcb6f9a0ef735477b7fadf8374",
		ExpectedOPc: "1006020f0a478bf6b699f15c062e42b3",
		ExpectedRES: "8011c48c0c214ed2",
		ExpectedAK:  "33484dc2136b",
		ExpectedCK:  "5dbdbb2954e8f3cde665b046179a5098",
	}

	Testf2f5f3TestTable[f2f5f3Test4] = &f2f5f3Test{
		K:           "9e5944aea94b81165c82fbf9f32db751",
		RAND:        "ce83dbc54ac0274a157c17f80d017bd6",
		OP:          "223014c5806694c007ca1eeef57f004f",
		ExpectedOPc: "a64a507ae1a2a98bb88eb4210135dc87",
		ExpectedRES: "f365cd683cd92e96",
		ExpectedAK:  "f0b9c08ad02e",
		ExpectedCK:  "e203edb3971574f5a94b0d61b816345d",
	}

	Testf2f5f3TestTable[f2f5f3Test4] = &f2f5f3Test{
		K:           "4ab1deb05ca6ceb051fc98e77d026a84",
		RAND:        "74b0cd6031a1c8339b2b6ce2b8c4a186",
		OP:          "2d16c5cd1fdf6b22383584e3bef2a8d8",
		ExpectedOPc: "dcf07cbd51855290b92a07a9891e523e",
		ExpectedRES: "5860fc1bce351e7e",
		ExpectedAK:  "31e11a609118",
		ExpectedCK:  "7657766b373d1c2138f307e3de9242f9",
	}

	Testf4f5StarTestTable[f4f5StarTest1] = &f4f5StarTest{
		K:              "465b5ce8b199b49faa5f0a2ee238a6bc",
		RAND:           "23553cbe9637a89d218ae64dae47bf35",
		OP:             "cdc202d5123e20f62b6d676ac72cb318",
		ExpectedOPc:    "cd63cb71954a9f4e48a5994e37a02baf",
		ExpectedIK:     "f769bcd751044604127672711c6d3441",
		ExpectedAKStar: "451e8beca43b",
	}

	Testf4f5StarTestTable[f4f5StarTest2] = &f4f5StarTest{
		K:              "0396eb317b6d1c36f19c1c84cd6ffd16",
		RAND:           "c00d603103dcee52c4478119494202e8",
		OP:             "ff53bade17df5d4e793073ce9d7579fa",
		ExpectedOPc:    "53c15671c60a4b731c55b4a441c0bde2",
		ExpectedIK:     "21a8c1f929702adb3e738488b9f5c5da",
		ExpectedAKStar: "30f1197061c1",
	}

	Testf4f5StarTestTable[f4f5StarTest3] = &f4f5StarTest{
		K:              "fec86ba6eb707ed08905757b1bb44b8f",
		RAND:           "9f7c8d021accf4db213ccff0c7f71a6a",
		OP:             "dbc59adcb6f9a0ef735477b7fadf8374",
		ExpectedOPc:    "1006020f0a478bf6b699f15c062e42b3",
		ExpectedIK:     "59a92d3b476a0443487055cf88b2307b",
		ExpectedAKStar: "deacdd848cc6",
	}

	Testf4f5StarTestTable[f4f5StarTest4] = &f4f5StarTest{
		K:              "9e5944aea94b81165c82fbf9f32db751",
		RAND:           "ce83dbc54ac0274a157c17f80d017bd6",
		OP:             "223014c5806694c007ca1eeef57f004f",
		ExpectedOPc:    "a64a507ae1a2a98bb88eb4210135dc87",
		ExpectedIK:     "0c4524adeac041c4dd830d20854fc46b",
		ExpectedAKStar: "6085a86c6f63",
	}

	Testf4f5StarTestTable[f4f5StarTest5] = &f4f5StarTest{
		K:              "4ab1deb05ca6ceb051fc98e77d026a84",
		RAND:           "74b0cd6031a1c8339b2b6ce2b8c4a186",
		OP:             "2d16c5cd1fdf6b22383584e3bef2a8d8",
		ExpectedOPc:    "dcf07cbd51855290b92a07a9891e523e",
		ExpectedIK:     "1c42e960d89b8fa99f2744e0708ccb53",
		ExpectedAKStar: "fe2555e54aa9",
	}

	Testf4f5StarTestTable[f4f5StarTest6] = &f4f5StarTest{
		K:              "6c38a116ac280c454f59332ee35c8c4f",
		RAND:           "ee6466bc96202c5a557abbeff8babf63",
		OP:             "1ba00a1a7c6700ac8c3ff3e96ad08725",
		ExpectedOPc:    "3803ef5363b947c6aaa225e58fae3934",
		ExpectedIK:     "a7466cc1e6b2a1337d49d3b66e95d7b4",
		ExpectedAKStar: "1f53cd2b1113",
	}

}

func TestF1Test35207(t *testing.T) {

	for i, testTable := range Testf1TestTable {
		K, _ := hex.DecodeString(strings.Repeat(testTable.K, 1))
		RAND, _ := hex.DecodeString(strings.Repeat(testTable.RAND, 1))
		SQN, _ := hex.DecodeString(strings.Repeat(testTable.SQN, 1))
		AMF, _ := hex.DecodeString(strings.Repeat(testTable.AMF, 1))
		OP, _ := hex.DecodeString(strings.Repeat(testTable.OP, 1))
		ExpectedOPc, _ := hex.DecodeString(strings.Repeat(testTable.ExpectedOPc, 1))
		f1, _ := hex.DecodeString(strings.Repeat(testTable.f1, 1))
		f1Start, _ := hex.DecodeString(strings.Repeat(testTable.f1Start, 1))

		OPC, err := milenage.GenerateOPC(K, OP)
		if err != nil {
			t.Errorf("err: %+v\n", err)
		}

		fmt.Printf("K=%x\nSQN=%x\nOP=%x\nOPC=%x\n", K, SQN, OP, OPC)
		if !reflect.DeepEqual(OPC, ExpectedOPc) {
			t.Errorf("Testf1Test35207%s \t OPC[0x%x] \t ExpectedOPc[0x%x]\n", i, OPC, ExpectedOPc)
		}
		MAC_A, MAC_S := make([]byte, 8), make([]byte, 8)
		milenage.F1(OPC, K, RAND, SQN, AMF, MAC_A, MAC_S)

		if !reflect.DeepEqual(MAC_A, f1) {
			t.Errorf("Testf1Test35207%s \t MAC_A[0x%x] \t f1[0x%x]\n", i, MAC_A, f1)
		}
		if !reflect.DeepEqual(MAC_S, f1Start) {
			t.Errorf("Testf1Test35207%s \t MAC_S[0x%x] \t f1Start[0x%x]\n", i, MAC_S, f1Start)
		}

	}
}

func TestF2F5F3Test35207(t *testing.T) {
	for i, testTable := range Testf2f5f3TestTable {
		K, _ := hex.DecodeString(strings.Repeat(testTable.K, 1))
		OP, _ := hex.DecodeString(strings.Repeat(testTable.OP, 1))
		ExpectedOPc, _ := hex.DecodeString(strings.Repeat(testTable.ExpectedOPc, 1))
		ExpectedRES, _ := hex.DecodeString(strings.Repeat(testTable.ExpectedRES, 1))
		ExpectedAK, _ := hex.DecodeString(strings.Repeat(testTable.ExpectedAK, 1))
		ExpectedCK, _ := hex.DecodeString(strings.Repeat(testTable.ExpectedCK, 1))
		RAND, _ := hex.DecodeString(strings.Repeat(testTable.RAND, 1))

		OPC, err := milenage.GenerateOPC(K, OP)
		if err != nil {
			t.Errorf("err: %+v\n", err)
		}

		if !reflect.DeepEqual(OPC, ExpectedOPc) {
			t.Errorf("TestF2F5F3Test35207%s \t OPC[0x%x] \t ExpectedOPc[0x%x]\n", i, OPC, ExpectedOPc)
		}
		CK, IK := make([]byte, 16), make([]byte, 16)
		RES := make([]byte, 8)
		AK, AKstar := make([]byte, 6), make([]byte, 6)
		milenage.F2345(OPC, K, RAND, RES, CK, IK, AK, AKstar)
		if !reflect.DeepEqual(RES, ExpectedRES) {
			t.Errorf("TestF2F5F3Test35207%s \t RES[0x%x] \t ExpectedRES[0x%x]\n", i, RES, ExpectedRES)
		}
		if !reflect.DeepEqual(AK, ExpectedAK) {
			t.Errorf("TestF2F5F3Test35207%s \t AK[0x%x] \t ExpectedAK[0x%x]\n", i, AK, ExpectedAK)
		}
		if !reflect.DeepEqual(CK, ExpectedCK) {
			t.Errorf("TestF2F5F3Test35207%s \t CK[0x%x] \t ExpectedCK[0x%x]\n", i, CK, ExpectedCK)
		}
	}
}

func TestF4F5StarTest35207(t *testing.T) {
	for i, testTable := range Testf4f5StarTestTable {
		K, _ := hex.DecodeString(strings.Repeat(testTable.K, 1))
		OP, _ := hex.DecodeString(strings.Repeat(testTable.OP, 1))
		ExpectedOPc, _ := hex.DecodeString(strings.Repeat(testTable.ExpectedOPc, 1))

		ExpectedAKStar, _ := hex.DecodeString(strings.Repeat(testTable.ExpectedAKStar, 1))
		ExpectedIK, _ := hex.DecodeString(strings.Repeat(testTable.ExpectedIK, 1))
		RAND, _ := hex.DecodeString(strings.Repeat(testTable.RAND, 1))

		OPC, err := milenage.GenerateOPC(K, OP)
		if err != nil {
			t.Errorf("err: %+v\n", err)
		}

		if !reflect.DeepEqual(OPC, ExpectedOPc) {
			t.Errorf("TestF4F5StarTest35207%s \t OPC[0x%x] \t ExpectedOPc[0x%x]\n", i, OPC, ExpectedOPc)
		}
		CK, IK := make([]byte, 16), make([]byte, 16)
		RES := make([]byte, 8)
		AK, AKstar := make([]byte, 6), make([]byte, 6)
		milenage.F2345(OPC, K, RAND, RES, CK, IK, AK, AKstar)
		if !reflect.DeepEqual(AKstar, ExpectedAKStar) {
			t.Errorf("TestF4F5StarTest35207%s \t AKstar[0x%x] \t ExpectedAKStar[0x%x]\n", i, AKstar, ExpectedAKStar)
		}
		if !reflect.DeepEqual(IK, ExpectedIK) {
			t.Errorf("TestF4F5StarTest35207%s \t IK[0x%x] \t ExpectedIK[0x%x]\n", i, IK, ExpectedIK)
		}

	}
}

func TestGenerateOPC(t *testing.T) {

	//K_str := "3016ebeae2c45bd0060923dbbb402be6"
	K_str := "000102030405060708090a0b0c0d0e0f" // CHT
	K, _ := hex.DecodeString(K_str)

	// OP_str := "00000000000000000000000000000000"
	OP_str := "00112233445566778899aabbccddeeff" // CHT
	OP, _ := hex.DecodeString(OP_str)
	fmt.Println("K:", K)

	fmt.Println("OP:", OP)

	OPCbyGo, err := milenage.GenerateOPC(K, OP)
	if err != nil {
		t.Errorf("err: %+v\n", err)
	}

	fmt.Println("OPCbyGo:", OPCbyGo)

}

func TestRAND(t *testing.T) {
	/*
		K, RAND, CK, IK: 128 bits (16 bytes) (hex len = 32)
		SQN, AK: 48 bits (6 bytes) (hex len = 12) TS33.102 - 6.3.2
		AMF: 16 bits (2 bytes) (hex len = 4) TS33.102 - Annex H
	*/

	K, OP := make([]byte, 16), make([]byte, 16)
	K_str := "5122250214c33e723a5dd523fc145fc0"
	OP_str := "c9e8763286b5b9ffbdf56e1297d0887b"
	SQN_str := "16f3b3f70fc2"

	K, _ = hex.DecodeString(K_str)
	OP, _ = hex.DecodeString(OP_str)
	SQN, _ := hex.DecodeString(SQN_str)

	OPC, err := milenage.GenerateOPC(K, OP)
	if err != nil {
		t.Errorf("err: %+v\n", err)
	}

	fmt.Printf("K=%x\nSQN=%x\nOP=%x\nOPC=%x\n", K, SQN, OP, OPC)
	RAND := make([]byte, 16)
	RAND, _ = hex.DecodeString("81e92b6c0ee0e12ebceba8d92a99dfa5")

	AMF, _ := hex.DecodeString("c3ab")
	fmt.Printf("RAND=%x\nAMF=%x\n", RAND, AMF)

	// for test
	// RAND, _ = hex.DecodeString(TestGenAuthData.MilenageTestSet19.RAND)
	// AMF, _ = hex.DecodeString(TestGenAuthData.MilenageTestSet19.AMF)
	fmt.Printf("For test: RAND=%x, AMF=%x\n", RAND, AMF)

	// Run milenage
	MAC_A, MAC_S := make([]byte, 8), make([]byte, 8)
	CK, IK := make([]byte, 16), make([]byte, 16)
	RES := make([]byte, 8)
	AK, AKstar := make([]byte, 6), make([]byte, 6)

	// Generate MAC_A, MAC_S

	milenage.F1(OPC, K, RAND, SQN, AMF, MAC_A, MAC_S)

	// Generate RES, CK, IK, AK, AKstar
	// RES == XRES (expected RES) for server
	milenage.F2345(OPC, K, RAND, RES, CK, IK, AK, AKstar)
	fmt.Printf("milenage RES = %s\n", hex.EncodeToString(RES))
	//
	fmt.Printf("RES=%x\n", RES)
	expRES, _ := hex.DecodeString("28d7b0f2a2ec3de5")
	if !reflect.DeepEqual(RES, RES) {
		t.Errorf("RES[0x%x] \t expected[0x%x]\n", RES, expRES)
	}
	// Generate AUTN
	fmt.Printf("CK=%x\n", CK)
	expCK, _ := hex.DecodeString("5349fbe098649f948f5d2e973a81c00f")
	if !reflect.DeepEqual(CK, expCK) {
		t.Errorf("CK[0x%x] \t expected[0x%x]\n", CK, expCK)
	}
	fmt.Printf("IK=%x\n", IK)
	expIK, _ := hex.DecodeString("9744871ad32bf9bbd1dd5ce54e3e2e5a")
	if !reflect.DeepEqual(IK, expIK) {
		t.Errorf("IK[0x%x] \t expected[0x%x]\n", IK, expIK)
	}
	// fmt.Printf("SQN=%x\nAK =%x\n", SQN, AK)
	// fmt.Printf("AK=%x\n", AK)
	expAK, _ := hex.DecodeString("ada15aeb7bb8")
	if !reflect.DeepEqual(AK, expAK) {
		t.Errorf("AK[0x%x] \t expected[0x%x]\n", AK, expAK)
	}
	// fmt.Printf("AMF=%x, MAC_A=%x\n", AMF, MAC_A)
	SQNxorAK := make([]byte, 6)
	for i := 0; i < len(SQN); i++ {
		SQNxorAK[i] = SQN[i] ^ AK[i]
	}

	fmt.Printf("SQN xor AK = %x\n", SQNxorAK)
	AUTN := append(append(SQNxorAK, AMF...), MAC_A...)

	// fmt.Printf("MAC_A = %x\n", MAC_A)
	// fmt.Printf("MAC_S = %x\n", MAC_S)

	// fmt.Printf("AUTN = %x\n", AUTN)

	expAUTN, _ := hex.DecodeString("bb52e91c747ac3ab2a5c23d15ee351d5")
	if !reflect.DeepEqual(AUTN, expAUTN) {
		t.Errorf("AUTN[0x%x] \t expected[0x%x]\n", AUTN, expAUTN)
	}

}
