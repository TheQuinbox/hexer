package hextools

import (
	"encoding/hex"
)

func Decode(input string) string {
	result, err := hex.DecodeString(input)
	if err != nil {
		panic(err)
	}
	return string(result)
}

func Encode(input string) string {
	result := hex.EncodeToString([]byte(input))
	return result
}
