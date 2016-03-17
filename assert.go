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
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"runtime"
	"strings"
)

type T interface {
	Fatal(args ...interface{})
}

type Log struct {
}

func (log *Log) Fatal(args ...interface{}) {
	fmt.Println(args[0])
}

func Equal(t T, expected, actual interface{}) {
	expectedType := reflect.TypeOf(expected)
	actualType := reflect.TypeOf(actual)
	if expectedType != actualType {
		fail(t, "expected (%v), but was (%v)", expectedType, actualType)
		return
	}
	f := "expected %v, but was %v"
	if actualType.Kind() == reflect.String {
		s1, _ := expected.(string)
		s2, _ := actual.(string)
		if strings.Contains(s1+s2, "\n") {
			f = "expected: '%v'\n----\nbut was : '%v'"
		} else {
			f = "expected '%v', but was '%v'"
		}
	}
	if expected != actual {
		fail(t, f, expected, actual)
	}
}

func NotEqual(t T, unexpected, actual interface{}) {
	unexpectedType := reflect.TypeOf(unexpected)
	actualType := reflect.TypeOf(actual)
	if unexpectedType != actualType {
		return
	}
	f := "values should be different, actual: %v"
	if actualType.Kind() == reflect.String {
		f = "values should be different, actual: '%v'"
	}
	if unexpected == actual {
		fail(t, f, actual)
	}
}

func True(t T, ret bool, f string, args ...interface{}) {
	if !ret {
		fail(t, f, args...)
	}
}

func False(t T, ret bool, f string, args ...interface{}) {
	if ret {
		fail(t, f, args...)
	}
}

func fail(t T, format string, args ...interface{}) {
	t.Fatal(fmt.Sprintf(format, args...) + trace())
}

func trace() string {
	_, file, line, ok := runtime.Caller(3) // caller + fail + public func
	if !ok {
		return "\n???:??"
	}

	var buf bytes.Buffer
	finfo, err := os.Stat(file)
	if err != nil {
		buf.WriteString(fmt.Sprintf("\n%s:%d:", file, line))
	} else {
		buf.WriteString(fmt.Sprintf("\n%s:%d:", finfo.Name(), line))
	}
	source, err := ioutil.ReadFile(file)
	if err == nil {
		lines := strings.Split(string(source), "\n")
		context := 2
		from := line - context - 1
		to := line + context
		for j, l := range lines[from:to] {
			buf.WriteByte('\n')
			if j == context {
				buf.WriteString("=>")
			} else {
				buf.WriteString("  ")
			}
			buf.WriteString(l)
		}
	}
	return buf.String()
}
