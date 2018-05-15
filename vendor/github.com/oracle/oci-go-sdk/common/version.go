// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated by go generate; DO NOT EDIT

package common

import (
	"bytes"
	"fmt"
	"sync"
)

const (
	major = "1"
	minor = "4"
	patch = "0"
	tag   = ""
)

var once sync.Once
var version string

// Version returns semantic version of the sdk
func Version() string {
	once.Do(func() {
		ver := fmt.Sprintf("%s.%s.%s", major, minor, patch)
		verBuilder := bytes.NewBufferString(ver)
		if tag != "" && tag != "-" {
			_, err := verBuilder.WriteString(tag)
			if err == nil {
				verBuilder = bytes.NewBufferString(ver)
			}
		}
		version = verBuilder.String()
	})
	return version
}
