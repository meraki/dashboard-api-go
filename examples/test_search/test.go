package main

import (
	"fmt"
	"reflect"
)

// Define a custom comparison function type.
type cmpFunc func(interface{}, interface{}) bool

// Simple comparison function that compares if two values are equal.
func simpleCmp(a, b interface{}) bool {
	return a == b
}

// getDictResult function translates the Python function to Go.
func getDictResult(result interface{}, key string, value interface{}, cmpFn cmpFunc) interface{} {
	// Check if the result is a list.
	if reflect.TypeOf(result).Kind() == reflect.Slice {
		resultSlice := reflect.ValueOf(result)

		// If the list has a single element.
		if resultSlice.Len() == 1 {
			// Check if that element is a dictionary.
			if reflect.TypeOf(resultSlice.Index(0).Interface()).Kind() == reflect.Map {
				resultMap := resultSlice.Index(0).Interface().(map[string]interface{})
				// Check if the key and value match the provided ones.
				if val, ok := resultMap[key]; ok && !cmpFn(val, value) {
					return nil
				}
				return resultMap
			}
			return nil
		}

		// Iterate over the elements of the list.
		for i := 0; i < resultSlice.Len(); i++ {
			item := resultSlice.Index(i)
			// Check if the item is a dictionary.
			if reflect.TypeOf(item.Interface()).Kind() == reflect.Map {
				itemMap := item.Interface().(map[string]interface{})
				// Check if the key and value match the provided ones.
				if val, ok := itemMap[key]; !ok || cmpFn(val, value) {
					return itemMap
				}
			}
		}
		return nil
	}

	// If the result is not a list.
	if reflect.TypeOf(result).Kind() != reflect.Map {
		return nil
	}

	// Check if the result is a dictionary.
	resultMap := result.(map[string]interface{})
	// Check if the key and value match the provided ones.
	if val, ok := resultMap[key]; ok && !cmpFn(val, value) {
		return nil
	}

	return resultMap
}

func main() {
	// Example usage
	result := []map[string]interface{}{
		{"name": "John", "age": 30},
		{"name": "Alice", "age": 25},
		{"name": "Bob", "age": 35},
	}

	key := "name"
	value := "Alice"
	result2 := getDictResult(result, key, value, simpleCmp)

	fmt.Println(result2)
}
