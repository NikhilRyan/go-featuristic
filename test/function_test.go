package tests

import (
	"github.com/nikhilryan/go-featuristic/featuristic/services"
	"testing"
)

func TestFunctionRegistry(t *testing.T) {
	// Register functions from the package
	err := services.RegisterFunctionsFromPackage(services.FunctionPackage{})
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

	// Test Add function
	a, b := 5, 3
	value, err = services.CallFunction("Add", a, b)
	if err != nil {
		t.Fatalf("Error calling Add function: %v", err)
	}
	expectedSum := 8
	if value != expectedSum {
		t.Fatalf("Expected %d but got %d", expectedSum, value)
	}

	// Test UserGreeting function
	user := services.User{FirstName: "John", LastName: "Doe", Age: 30}
	value, err = services.CallFunction("UserGreeting", user)
	if err != nil {
		t.Fatalf("Error calling UserGreeting function: %v", err)
	}
	expectedGreeting := "Hello, John Doe, age 30!"
	if value != expectedGreeting {
		t.Fatalf("Expected %s but got %s", expectedGreeting, value)
	}

	// Test Multiply function
	value, err = services.CallFunction("Multiply", a, b)
	if err != nil {
		t.Fatalf("Error calling Multiply function: %v", err)
	}
	expectedProduct := 15
	if value != expectedProduct {
		t.Fatalf("Expected %d but got %d", expectedProduct, value)
	}
}
