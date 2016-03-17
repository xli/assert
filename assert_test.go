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
