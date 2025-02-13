package main

import (
	"strconv"
	"syscall/js"
)

// Add function that adds values from input fields and displays the result
func add(this js.Value, p []js.Value) interface{} {
	// Fetch the values of the input fields from the DOM
	value1 := js.Global().Get("document").Call("getElementById", p[0].String()).Get("value").String()
	value2 := js.Global().Get("document").Call("getElementById", p[1].String()).Get("value").String()

	// Convert string values to integers
	int1, _ := strconv.Atoi(value1)
	int2, _ := strconv.Atoi(value2)

	// Set the result in the third input field
	js.Global().Get("document").Call("getElementById", p[2].String()).Set("value", int1+int2)
	return nil
}

// Subtract function that subtracts values from input fields and displays the result
func subtract(this js.Value, p []js.Value) interface{} {
	// Fetch the values of the input fields from the DOM
	value1 := js.Global().Get("document").Call("getElementById", p[0].String()).Get("value").String()
	value2 := js.Global().Get("document").Call("getElementById", p[1].String()).Get("value").String()

	// Convert string values to integers
	int1, _ := strconv.Atoi(value1)
	int2, _ := strconv.Atoi(value2)

	// Set the result in the third input field
	js.Global().Get("document").Call("getElementById", p[2].String()).Set("value", int1-int2)
	return nil
}

// Register the Go functions as JavaScript callbacks
func registerCallbacks() {
	js.Global().Set("add", js.FuncOf(add))           // Register the add function
	js.Global().Set("subtract", js.FuncOf(subtract)) // Register the subtract function
}

func main() {
	// Create a channel to block and keep the Go function alive
	c := make(chan struct{}, 0)

	// Print initialization message
	println("WASM Go Initialized")

	// Register the functions for JavaScript usage
	registerCallbacks()

	// Block indefinitely to keep the Go functions available
	<-c
}
