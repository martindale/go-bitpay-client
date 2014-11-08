package btcbase58

import (
	"encoding/hex"
	"testing"
)

var TestTable = [][]string{
	{"61", "2g", "C2dGTwc"},
	{"626262", "a3gV", "4jF5uERJAK"},
	{"636363", "aPEr", "4mT4krqUYJ"},
	{"73696d706c792061206c6f6e6720737472696e67", "2cFupjhnEsSn59qHXstmK2ffpLv2", "BXF1HuEUCqeVzZdrKeJjG74rjeXxqJ7dW"},
	{"00eb15231dfceb60925886b67d065299925915aeb172c06647", "1NS17iag9jJgTHD1VXjvLCEnZuQ3rJDE9L", "13REmUhe2ckUKy1FvM7AMCdtyYq831yxM3QeyEu4"},
	{"516b6fcd0f", "ABnLTmg", "237LSrY9NUUas"},
	{"bf4f89001e670274dd", "3SEo3LWLoPntC", "GwDDDeduj1jpykc27e"},
	{"572e4794", "3EFU7m", "FamExfqCeza"},
	{"ecac89cad93923c02321", "EJDM8drfXA6uyA", "2W1Yd5Zu6WGyKVtHGMrH"},
	{"10c8511e", "Rt5zm", "3op3iuGMmhs"},
	{"00000000000000000000", "1111111111", "111111111146Momb"},
	{"", "", "3QJmnh"},
}

func TestEncode(t *testing.T) {

	for _, line := range TestTable {
		test := []byte(line[0])

		b10 := make([]byte, hex.DecodedLen(len(test)))
		hex.Decode(b10, test)
		// This method accepts base10
		hash, _ := Encode(b10)

		if string(hash) != line[1] {
			t.Error(line[0])
			t.Errorf("Encode mismatch expected: %s was: %s",
				line[1], hash)
		}

		if b := Check(b10); b != line[2] {
			t.Errorf("Encode check failed expected: %s was: %s",
				line[2], b)
		}
	}
}

func TestDecode(t *testing.T) {
	for _, line := range TestTable {
		hash, _, err := Decode([]byte(line[1]))
		if err != nil {
			t.Error(err)
		}
		b10 := make([]byte, hex.EncodedLen(len(hash)))
		hex.Encode(b10, hash)
		if string(b10) != line[0] {
			t.Error(line[0])
			t.Errorf("Decode mismatch expected: %s was: %s",
				line[1], b10)
		}

	}
}
