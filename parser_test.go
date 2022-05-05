package main

import (
	"reflect"
	"testing"
)

const testData = `package main
type Gopher interface {
	Hello()
	Say(string)
	Introduce() string
}

`

func TestNewParser(t *testing.T) {
	// tests := []struct {
	// 	name string
	// 	want *Parser
	// }{
	// 	// TODO: Add test cases.
	// }
	// for _, tt := range tests {
	t.Run("Test constructor.", func(t *testing.T) {
		if got := NewParser(); reflect.DeepEqual(got, nil) {
			t.Errorf("NewParser() = %v, not want %v", got, nil)
		}
	})
	// }
}

func TestParser_Parse(t *testing.T) {
	want := "Gopher"
	t.Run("Test parse.", func(t *testing.T) {
		parser := NewParser()
		if got := parser.Parse(testData); got != want {
			t.Errorf("NewParser() = %v, want %v", got, want)
		}
	})
}
