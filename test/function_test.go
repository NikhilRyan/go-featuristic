package tests

import (
	"fmt"
	"github.com/nikhilryan/go-featuristic/featuristic/services"
	"testing"
)

type CustomFunctions struct{}

func (cf CustomFunctions) CustomGreet(name string) string {
	return fmt.Sprintf("Greetings, %s!", name)
}

func TestFunctionRegistry(t *testing.T) {
	// Register functions from the package
	err := services.RegisterFunctionsFromPackage(services.FunctionPackage{})
	if err != nil {
		return
	}

	// Register external functions from the user's codebase
	err = services.RegisterFunctionsFromPackage(CustomFunctions{})
	if err != nil {
		return
	}

	// Test HelloWorld function
	name := "John Doe"
	value, err := services.CallFunction("HelloWorld", name)
	if err != nil {
		t.Fatalf("Error calling HelloWorld function: %v", err)
	}
	expected := "Hello, John Doe!"
	if value != expected {
		t.Fatalf("Expected %s but got %s", expected, value)
	}

	// Test CustomGreet function
	value, err = services.CallFunction("CustomGreet", name)
	if err != nil {
		t.Fatalf("Error calling CustomGreet function: %v", err)
	}
	expectedCustom := "Greetings, John Doe!"
	if value != expectedCustom {
		t.Fatalf("Expected %s but got %s", expectedCustom, value)
	}
}
