package config

import (
	"testing"
)

func TestUnmarshal(t *testing.T) {
	var err error

	correct := []byte(`
services:
  - name: foo
    build:
      context: context
      dockerfile: dockerfile
`)
	_, err = Unmarshal(correct)

	if err != nil {
		t.Fatal(string(err.Error()))
	}
}
