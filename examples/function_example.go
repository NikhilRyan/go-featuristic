package examples

import (
	"fmt"
	"github.com/nikhilryan/go-featuristic/featuristic/models"
	"github.com/nikhilryan/go-featuristic/featuristic/services"
)

type CustomFunctions struct{}

func (cf CustomFunctions) CustomGreet(name string) string {
	return fmt.Sprintf("Greetings, %s!", name)
}

func RunFunctionExample() {

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

	// Create a feature flag of type function
	functionFlag := models.FeatureFlag{
		Namespace: "default",
		Key:       "customGreetFunction",
		Value:     "CustomGreet", // Method name in CustomFunctions
		Type:      services.FlagTypeFunction,
	}

	err = featureFlagService.CreateFlag(&functionFlag)
	if err != nil {
		fmt.Println("Error creating function flag:", err)
		return
	}

	// Example: Get flag value with arguments
	value, err := featureFlagService.GetFunctionValue("default", "customGreetFunction", "John Doe")
	if err != nil {
		fmt.Println("Error getting function flag value:", err)
	} else {
		fmt.Println("Function flag value:", value)
	}
}
