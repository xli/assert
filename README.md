assert
================

1. Simplest assertions for your Go test.
2. Test failure also outputs assertion context.

Example:

        import (
        	"github.com/xli/assert"
        	"testing"
        )
        func TestHelloWorld(t *testing.T) {
        	assert.Equal(t, "hello", "world")
        }

Output:

        // expected 'hello', but was 'world'
        // <file name>.go:<line number>:
        //   )
        //   	func TestHelloWorld(t *testing.T) {
        // =>	Equal(t, "hello", "world")
        //   }


License: Apache License v2