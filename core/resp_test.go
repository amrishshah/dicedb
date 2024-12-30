package core_test

import (
	"log"
	"reflect"
	"testing"

	"github.com/amrishkshah/dicedb/core"
)

func TestSimpleStringDecode(t *testing.T) {
	cases := map[string]string{
		"+OK\r\n": "OK",
	}
	for k, v := range cases {
		values, _ := core.Decode([]byte(k))
		println(v)
		println("Test ---->")
		//log.Panicln(values)
		for _, value := range values {
			if v != value {
				t.Fail()
			}
		}
	}
}

func TestErrorDecode(t *testing.T) {
	cases := map[string]string{
		"-Error Message\r\n": "Error Message",
	}
	for k, v := range cases {
		values, _ := core.Decode([]byte(k))
		for _, value := range values {
			if v != value {
				t.Fail()
			}
		}
	}
}

func TestInt64Decode(t *testing.T) {
	cases := map[string]int64{
		":123\r\n": 123,
	}
	for k, v := range cases {
		values, _ := core.Decode([]byte(k))
		for _, value := range values {
			if v != value {
				t.Fail()
			}
		}
	}
}

func TestBulkStringDecode(t *testing.T) {
	cases := map[string]string{
		"$5\r\nhello\r\n": "hello",
		"$0\r\n\r\n":      "",
	}
	for k, v := range cases {
		values, _ := core.Decode([]byte(k))
		log.Println(values)
		for _, value := range values {
			if v != value {
				t.Fail()
			}
		}
	}
}

func TestArrayDecode(t *testing.T) {
	cases := map[string][]interface{}{
		"*0\r\n":                                                   {},
		"*2\r\n$5\r\nhello\r\n$5\r\nworld\r\n":                     {"hello", "world"},
		"*3\r\n:1\r\n:2\r\n:3\r\n":                                 {int64(1), int64(2), int64(3)},
		"*5\r\n:1\r\n:2\r\n:3\r\n:4\r\n$5\r\nhello\r\n":            {int64(1), int64(2), int64(3), int64(4), "hello"},
		"*2\r\n*3\r\n:1\r\n:2\r\n:3\r\n*2\r\n+Hello\r\n-World\r\n": {[]interface{}{int64(1), int64(2), int64(3)}, []interface{}{"Hello", "World"}},
	}

	for input, expected := range cases {
		values, err := core.Decode([]byte(input))
		if err != nil {
			t.Errorf("Decode(%q) returned error: %v", input, err)
			continue
		}

		// array, ok := values.([]interface{})
		// if !ok {
		// 	t.Errorf("Decode(%q) returned %T; want []interface{}", input, values)
		// 	continue
		// }
		arr := values[0].([]interface{})
		//log.Panicln()
		if len(arr) != len(expected) {
			t.Errorf("Decode(%q) returned array of length %d; want %d", input, len(arr), len(expected))
			continue
		}

		for i := range arr {
			if !reflect.DeepEqual(expected[i], arr[i]) {
				t.Errorf("Decode(%q) element %d = %v; want %v", input, i, arr[i], expected[i])
			}
		}
	}
}
