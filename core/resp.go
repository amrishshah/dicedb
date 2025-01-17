package core

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"strconv"
)

func readLength(data []byte) (int, int) {
	pos, length := 0, 0
	for ; data[length] != '\r'; length++ {

	}
	// for pos = range data {
	// 	b := data[pos]
	// 	if !(b >= '0' && b <= '9') {
	// 		log.Println(length)
	// 		log.Println(pos)
	// 		log.Println("ReadLength End")
	// 		return length, pos + 2
	// 	}
	// 	length = length*10 + int(b-'0')
	// }
	i, err := strconv.Atoi(string(data[pos : pos+length]))
	if err != nil {
		// ... handle error
		panic(err)
	}
	log.Println(i)
	log.Println(pos)
	log.Println("WE")
	return i, pos + length + 2
}

func readSimpleString(data []byte) (string, int, error) {

	pos := 1
	for ; data[pos] != '\r'; pos++ {

	}
	return string(data[1:pos]), pos + 2, nil
}

func readError(data []byte) (string, int, error) {
	return readSimpleString(data)
}

func readInt64(data []byte) (int64, int, error) {
	pos := 1
	var value int64 = 0
	for ; data[pos] != '\r'; pos++ {
		value = value*10 + int64(data[pos]-'0')
	}
	return value, pos + 2, nil
}

func readBulkString(data []byte) (string, int, error) {
	pos := 1
	log.Println("e")
	//log.Println(string(data[pos:]))
	len, delta := readLength(data[pos:])
	pos += delta
	//log.Println(string(data[pos : pos+len]))
	return string(data[pos : pos+len]), pos + len + 2, nil

}

func readArray(data []byte) (interface{}, int, error) {
	// first character *
	pos := 1

	// reading the length
	count, delta := readLength(data[pos:])
	pos += delta

	var elems []interface{} = make([]interface{}, count)
	for i := range elems {
		elem, delta, err := DecodeOne(data[pos:])
		if err != nil {
			return nil, 0, err
		}
		elems[i] = elem
		pos += delta
	}
	log.Println("readArray")
	log.Println(len(elems))
	return elems, pos, nil
}

func DecodeOne(data []byte) (interface{}, int, error) {
	if len(data) == 0 {
		return nil, 0, errors.New("no data")
	}
	log.Println("DecodeOne")
	log.Println(string(data))
	switch data[0] {
	case '+':
		return readSimpleString(data)
	case '-':
		return readError(data)
	case ':':
		return readInt64(data)
	case '$':
		return readBulkString(data)
	case '*':
		return readArray(data)
	}
	return nil, 0, nil
}
func Decode(data []byte) ([]interface{}, error) {
	if len(data) == 0 {
		return nil, errors.New("no data")
	}
	log.Println("Decode")
	log.Println(data)
	var values []interface{} = make([]interface{}, 0)
	var index int = 0
	for index < len(data) {
		value, delta, err := DecodeOne(data[index:])
		if err != nil {
			return values, err
		}
		if value == nil {
			log.Println(values)
			log.Panicln("sd")

		}
		index = index + delta
		values = append(values, value)
	}
	return values, nil
}

func encodeString(v string) []byte {
	return []byte(fmt.Sprintf("$%d\r\n%s\r\n", len(v), v))
}

func Encode(value interface{}, isSimple bool) []byte {
	switch v := value.(type) {
	case string:
		if isSimple {
			return []byte(fmt.Sprintf("+%s\r\n", v))
		}
		return encodeString(v)
	case []string:
		var b []byte
		buf := bytes.NewBuffer(b)
		for _, b := range value.([]string) {
			buf.Write(encodeString(b))
		}
		return []byte(fmt.Sprintf("*%d\r\n%s", len(v), buf.Bytes()))
	case int, int8, int16, int32, int64:
		return []byte(fmt.Sprintf(":%d\r\n", v))
	case error:
		return []byte(fmt.Sprintf("-%s\r\n", v))
	default:
		return RESP_NIL
	}
}
