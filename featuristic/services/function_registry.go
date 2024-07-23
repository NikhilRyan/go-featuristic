package services

import (
	"errors"
	"reflect"
)

type FunctionRegistry struct {
	functions map[string]reflect.Value
}

var functionRegistry = &FunctionRegistry{
	functions: make(map[string]reflect.Value),
}

func RegisterFunction(name string, fn interface{}) error {
	if _, exists := functionRegistry.functions[name]; exists {
		return errors.New("function already registered")
	}
	functionRegistry.functions[name] = reflect.ValueOf(fn)
	return nil
}

func RegisterFunctionsFromPackage(pkg interface{}) error {
	pkgValue := reflect.ValueOf(pkg)
	pkgType := pkgValue.Type()
	for i := 0; i < pkgType.NumMethod(); i++ {
		method := pkgType.Method(i)
		name := method.Name
		if _, exists := functionRegistry.functions[name]; !exists {
			functionRegistry.functions[name] = method.Func
		}
	}
	return nil
}

func CallFunction(name string, args ...interface{}) (interface{}, error) {
	fn, exists := functionRegistry.functions[name]
	if !exists {
		return nil, errors.New("function not found")
	}

	fnType := fn.Type()
	if len(args) != fnType.NumIn() {
		return nil, errors.New("incorrect number of arguments")
	}

	inputs := make([]reflect.Value, len(args))
	for i, arg := range args {
		inputs[i] = reflect.ValueOf(arg)
	}

	results := fn.Call(inputs)
	if len(results) == 0 {
		return nil, nil
	}

	return results[0].Interface(), nil
}
