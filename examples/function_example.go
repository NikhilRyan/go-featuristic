package examples

import (
	"fmt"
	"github.com/nikhilryan/go-featuristic/featuristic/models"
	"github.com/nikhilryan/go-featuristic/featuristic/services"
)

func RunFunctionExample() {

	// Register functions from the package
	err := services.RegisterFunctionsFromPackage(services.FunctionPackage{})
	if err != nil {
		return
	}

	// Create a feature flag of type function
	functionFlag := models.FeatureFlag{
		Namespace: "default",
		Key:       "userGreetingFunction",
		Value:     "UserGreeting", // Method name in FunctionPackage
		Type:      services.FlagTypeFunction,
	}

	err = featureFlagService.CreateFlag(&functionFlag)
	if err != nil {
		fmt.Println("Error creating function flag:", err)
		return
	}

	// Example: Get flag value with arguments
	user := services.User{FirstName: "John", LastName: "Doe", Age: 30}
	value, err := featureFlagService.GetFunctionValue("default", "userGreetingFunction", user)
	if err != nil {
		fmt.Println("Error getting function flag value:", err)
	} else {
		fmt.Println("Function flag value:", value)
	}
}
