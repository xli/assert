assert
================

1. Simplest assertions for your Go test.
2. Test failure also outputs assertion context.

Example:

```go
import (
	"github.com/xli/assert"
	"testing"
)
func TestHelloWorld(t *testing.T) {
	assert.Equal(t, "hello", "world")
}
```
Output:

```
--- FAIL: TestHelloWorld (0.00s)
	assert.go:112: expected 'hello', but was 'world'
		x_test.go:9:
		  
		  func TestHelloWorld(t *testing.T) {
		=>	assert.Equal(t, "hello", "world")
		  }
```

[More examples](assert_test.go)

License: Apache License v2
