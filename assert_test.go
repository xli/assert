package assert_test

import (
	. "github.com/xli/assert"
)

func ExampleTrue() {
	var log Log
	True(&log, true, "msg")
	True(&log, false, "msg")
	// Output:
	// msg
	// assert_test.go:10:
	//   	var log Log
	//   	True(&log, true, "msg")
	// =>	True(&log, false, "msg")
	//   	// Output:
	//   	// msg
}

func ExampleFalse() {
	var log Log
	False(&log, false, "msg")
	False(&log, true, "msg")
	// Output:
	// msg
	// assert_test.go:24:
	//   	var log Log
	//   	False(&log, false, "msg")
	// =>	False(&log, true, "msg")
	//   	// Output:
	//   	// msg
}

func ExampleEqual() {
	var log Log
	Equal(&log, 3, 3)
	Equal(&log, "hello", "world")
	Equal(&log, 3, "3")
	// Output:
	// expected 'hello', but was 'world'
	// assert_test.go:38:
	//   	var log Log
	//   	Equal(&log, 3, 3)
	// =>	Equal(&log, "hello", "world")
	//   	Equal(&log, 3, "3")
	//   	// Output:
	//
	// expected (int), but was (string)
	// assert_test.go:39:
	//   	Equal(&log, 3, 3)
	//   	Equal(&log, "hello", "world")
	// =>	Equal(&log, 3, "3")
	//   	// Output:
	//   	// expected 'hello', but was 'world'
}

func ExampleEqual_multilinesString() {
	var log Log
	Equal(&log, "hello\nworld", "hello\nworld\n")
	// Output:
	// expected: 'hello
	// world'
	// ----
	// but was : 'hello
	// world
	// '
	// assert_test.go:60:
	//   func ExampleEqual_multilinesString() {
	//   	var log Log
	// =>	Equal(&log, "hello\nworld", "hello\nworld\n")
	//   	// Output:
	//   	// expected: 'hello
}

func ExampleNotEqual() {
	var log Log
	NotEqual(&log, "", 0)
	NotEqual(&log, "", "not")
	NotEqual(&log, "", "")
	NotEqual(&log, 1, 1)
	// Output:
	// values should be different, actual: ''
	// assert_test.go:80:
	//   	NotEqual(&log, "", 0)
	//   	NotEqual(&log, "", "not")
	// =>	NotEqual(&log, "", "")
	//   	NotEqual(&log, 1, 1)
	//   	// Output:
	//
	// values should be different, actual: 1
	// assert_test.go:81:
	//   	NotEqual(&log, "", "not")
	//   	NotEqual(&log, "", "")
	// =>	NotEqual(&log, 1, 1)
	//   	// Output:
	//   	// values should be different, actual: ''
}

func ExampleNil() {
	var log Log
	Nil(&log, nil)
	Nil(&log, "")
	Nil(&log, 0)
	Nil(&log, log)
	// Output:
	// expected nil, but was string: ''
	// assert_test.go:103:
	//   	var log Log
	//   	Nil(&log, nil)
	// =>	Nil(&log, "")
	//   	Nil(&log, 0)
	//   	Nil(&log, log)
	//
	// expected nil, but was int: 0
	// assert_test.go:104:
	//   	Nil(&log, nil)
	//   	Nil(&log, "")
	// =>	Nil(&log, 0)
	//   	Nil(&log, log)
	//   	// Output:
	//
	// expected nil, but was assert.Log: {2}
	// assert_test.go:105:
	//   	Nil(&log, "")
	//   	Nil(&log, 0)
	// =>	Nil(&log, log)
	//   	// Output:
	//   	// expected nil, but was string: ''
}

func ExampleNotNil() {
	var log Log
	NotNil(&log, nil)
	NotNil(&log, "")
	NotNil(&log, 0)
	NotNil(&log, log)
	// Output:
	// expected not nil
	// assert_test.go:134:
	//   func ExampleNotNil() {
	//   	var log Log
	// =>	NotNil(&log, nil)
	//   	NotNil(&log, "")
	//   	NotNil(&log, 0)
}

func ExampleEqual_map() {
	var log Log
	var m1 = map[string]string{
		"hello": "world",
	}
	var m2 = map[string]string{
		"hello": "world",
	}
	var m3 = map[string]string{
		"hello": "w",
	}
	Equal(&log, m1, m1)
	Equal(&log, m1, m2)
	Equal(&log, m1, m3)
	// Output:
	// expected map[hello:world], but was map[hello:w]
	// assert_test.go:161:
	//   	Equal(&log, m1, m1)
	//   	Equal(&log, m1, m2)
	// =>	Equal(&log, m1, m3)
	//   	// Output:
	//   	// expected map[hello:world], but was map[hello:w]
}
