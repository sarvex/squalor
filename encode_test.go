// Copyright 2012, Google Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package squalor

import (
	"bytes"
	"testing"
	"time"
)

func TestEncodeSQL(t *testing.T) {
	testCases := []struct {
		arg      interface{}
		expected string
	}{
		{nil, "NULL"},
		{(*int32)(nil), "NULL"},
		{(*string)(nil), "NULL"},
		{true, "1"},
		{false, "0"},
		{int(-1), "-1"},
		{int32(-1), "-1"},
		{int64(-1), "-1"},
		{uint(1), "1"},
		{uint32(1), "1"},
		{uint64(1), "1"},
		{1.23, "1.23"},
		{"abcd", "'abcd'"},
		{"workin' hard", "'workin\\' hard'"},
		{"\x00'\"\b\n\r\t\x1A\\", `'\0\'\"\b\n\r\t\Z\\'`},
		{[]byte(nil), "NULL"},
		{[]byte{}, "X''"},
		{[]byte("abcd"), "X'61626364'"},
		{[]byte("\x00'\"\b\n\r\t\x1A\\"), "X'002722080a0d091a5c'"},
		{time.Date(2012, time.February, 24, 23, 19, 43, 10, time.UTC), "'2012-02-24 23:19:43'"},
		{time.Date(1999, 1, 2, 3, 4, 5, 0, time.UTC), "'1999-01-02 03:04:05'"},
	}
	for _, c := range testCases {
		var buf bytes.Buffer
		if err := encodeSQLValue(&buf, c.arg); err != nil {
			t.Error(err)
			continue
		}
		encoded := buf.String()
		if encoded != c.expected {
			t.Errorf("Expected %s, but got %s", c.expected, encoded)
		}
	}
}

// Ensure dontEscape is not escaped
func TestDontEscape(t *testing.T) {
	if encodeMap[dontEscape] != dontEscape {
		t.Errorf("Encode fail: %v", encodeMap[dontEscape])
	}
	if decodeMap[dontEscape] != dontEscape {
		t.Errorf("Decode fail: %v", decodeMap[dontEscape])
	}
}