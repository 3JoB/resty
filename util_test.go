// Copyright (c) 2015-2023 Jeevanandam M (jeeva@myjeeva.com), All rights reserved.
// resty source code and usage is governed by a MIT style
// license that can be found in the LICENSE file.

package resty

import (
	"bytes"
	"mime/multipart"
	"testing"
)

func TestIsJSONType(t *testing.T) {
	for _, test := range []struct {
		input  string
		expect bool
	}{
		{input: "application/json", expect: true},
		{input: "application/xml+json", expect: true},
		{input: "application/vnd.foo+json", expect: true},

		{input: "application/json; charset=utf-8", expect: true},
		{input: "application/vnd.foo+json; charset=utf-8", expect: true},

		{input: "text/json", expect: true},
		{input: "text/xml+json", expect: true},
		{input: "text/vnd.foo+json", expect: true},

		{input: "application/foo-json", expect: false},
		{input: "application/foo.json", expect: false},
		{input: "application/vnd.foo-json", expect: false},
		{input: "application/vnd.foo.json", expect: false},
		{input: "application/json+xml", expect: false},

		{input: "text/foo-json", expect: false},
		{input: "text/foo.json", expect: false},
		{input: "text/vnd.foo-json", expect: false},
		{input: "text/vnd.foo.json", expect: false},
		{input: "text/json+xml", expect: false},
	} {
		result := IsJSONType(test.input)

		if result != test.expect {
			t.Errorf("failed on %q: want %v, got %v", test.input, test.expect, result)
		}
	}
}

func TestIsXMLType(t *testing.T) {
	for _, test := range []struct {
		input  string
		expect bool
	}{
		{input: "application/xml", expect: true},
		{input: "application/json+xml", expect: true},
		{input: "application/vnd.foo+xml", expect: true},

		{input: "application/xml; charset=utf-8", expect: true},
		{input: "application/vnd.foo+xml; charset=utf-8", expect: true},

		{input: "text/xml", expect: true},
		{input: "text/json+xml", expect: true},
		{input: "text/vnd.foo+xml", expect: true},

		{input: "application/foo-xml", expect: false},
		{input: "application/foo.xml", expect: false},
		{input: "application/vnd.foo-xml", expect: false},
		{input: "application/vnd.foo.xml", expect: false},
		{input: "application/xml+json", expect: false},

		{input: "text/foo-xml", expect: false},
		{input: "text/foo.xml", expect: false},
		{input: "text/vnd.foo-xml", expect: false},
		{input: "text/vnd.foo.xml", expect: false},
		{input: "text/xml+json", expect: false},
	} {
		result := IsXMLType(test.input)

		if result != test.expect {
			t.Errorf("failed on %q: want %v, got %v", test.input, test.expect, result)
		}
	}
}

func TestWriteMultipartFormFileReaderEmpty(t *testing.T) {
	w := multipart.NewWriter(bytes.NewBuffer(nil))
	defer func() { _ = w.Close() }()
	if err := writeMultipartFormFile(w, "foo", "bar", bytes.NewReader(nil)); err != nil {
		t.Errorf("Got unexpected error: %v", err)
	}
}

func TestWriteMultipartFormFileReaderError(t *testing.T) {
	err := writeMultipartFormFile(nil, "", "", &brokenReadCloser{})
	assertNotNil(t, err)
	assertEqual(t, "read error", err.Error())
}
