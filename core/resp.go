package core

import (
	"errors"
	"fmt"
	"log"
)

func DecodeArrayStr(data []byte) ([]string, error) {
	val, err := decode(data)
	if err != nil {
		return nil, err
	}
	ts := val.([]interface{})
	tokens := make([]string, len(ts))
	for i := range tokens {
		tokens[i] = ts[i].(string)
	}
	return tokens, nil
}

func Encode(value interface{}) []byte {
	switch v := value.(type) {
	case string:
		return []byte(fmt.Sprintf("+%s\r\n", v))
	}
	return []byte{}
}
func decode(data []byte) (interface{}, error) {
	if len(data) == 0 {
		return nil, errors.New("No data")
	}
	val, _, err := decodeOne(data)
	log.Println(val)
	return val, err
}

func decodeOne(data []byte) (interface{}, int, error) {
	switch data[0] {
	case '+':
		str, n, err := readSimpleStr(data)
		if err != nil {
			return nil, 0, err
		}
		return []interface{}{str}, n, nil
	default:
		return nil, 0, errors.New("Unsupported Type")
	}
}

func readSimpleStr(data []byte) (string, int, error) {
	// first character is +
	pos := 1
	for ; data[pos] != '\r'; pos++ {
	}
	return string(data[1:pos]), pos + 2, nil
}
