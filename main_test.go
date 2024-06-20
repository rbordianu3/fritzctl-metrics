package main

import (
	"bytes"
	"log"
	"os"
	"testing"

	"github.com/namsral/flag"
)

func TestValidateFlagsNoUsername(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer func() {
		log.SetOutput(os.Stderr)
	}()

	os.Setenv("USERNAME", "")
	os.Setenv("PASSWORD", "somepassword")
	*username = ""
	*password = "somepassword"
	*urlString = "https://fritz.box"
	flag.Parse()

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
		expected := "No username provided.\n"
		actual := buf.String()
		if actual != expected {
			t.Errorf("Unexpected log message. Expected: %q, Got: %q", expected, actual)
		}
	}()

	validateFlags()
}

func TestValidateFlagsNoPassword(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer func() {
		log.SetOutput(os.Stderr)
	}()

	os.Setenv("USERNAME", "someuser")
	os.Setenv("PASSWORD", "")
	*username = "someuser"
	*password = ""
	*urlString = "https://fritz.box"
	flag.Parse()

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
		expected := "No password provided.\n"
		actual := buf.String()
		if actual != expected {
			t.Errorf("Unexpected log message. Expected: %q, Got: %q", expected, actual)
		}
	}()

	validateFlags()
}