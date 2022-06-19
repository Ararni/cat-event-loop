package main

import (
	"reflect"
	"testing"

	"github.com/Ararni/cat-event-loop/engine"
	"github.com/stretchr/testify/assert"
)

// Must return print command
func TestParserPrint(t *testing.T) {
	commandString := "print Hello world"
	command := engine.Parse(commandString)
	examplePrint := engine.NewPrintCommand("Hello world")

	if assert.NotNil(t, command) {
		assert.IsType(t, reflect.TypeOf(examplePrint), reflect.TypeOf(command))
	}
}

// Must return cat command
func TestParserCat(t *testing.T) {
	commandString := "cat 4 6"
	command := engine.Parse(commandString)
	exampleCat := engine.NewCatCommand("4", "6")

	if assert.NotNil(t, command) {
		assert.IsType(t, reflect.TypeOf(exampleCat), reflect.TypeOf(command))
	}
}

// Must return print command (containing error message)
func TestParserDefault(t *testing.T) {
	commandString := ""
	command := engine.Parse(commandString)
	examplePrint := engine.NewPrintCommand("Hello world")

	if assert.NotNil(t, command) {
		assert.IsType(t, reflect.TypeOf(examplePrint), reflect.TypeOf(command))
	}
}
