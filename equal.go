/*
 * Copyright 2016 Xiao Li <swing1979@gmail.com>
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *  http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package assert

import (
	"fmt"
	"reflect"
	"testing"
)

func Equal(t *testing.T, expected, actual interface{}) {
	if expected == actual {
		return
	}

	msg := fmt.Sprintf("expected %v(%v), but was %v(%v)",
		expected, reflect.TypeOf(expected),
		actual, reflect.TypeOf(actual))
	Fail(t, msg)
}

func Fail(t *testing.T, msg string) {
	t.Fatal(msg)
}