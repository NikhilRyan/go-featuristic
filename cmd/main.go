package main

import (
	"flag"
	"github.com/nikhilryan/go-featuristic/examples"
	"log"
)

func main() {
	example := flag.String("example", "int", "Specify which example to run: int, float, string, intArray, floatArray, stringArray, rollout")
	flag.Parse()

	switch *example {
	case "int":
		examples.RunIntExample()
	case "float":
		examples.RunFloatExample()
	case "string":
		examples.RunStringExample()
	case "intArray":
		examples.RunIntArrayExample()
	case "floatArray":
		examples.RunFloatArrayExample()
	case "stringArray":
		examples.RunStringArrayExample()
	case "rollout":
		examples.RunRolloutExample()
	default:
		log.Fatalf("Unknown example: %s", *example)
	}
}
