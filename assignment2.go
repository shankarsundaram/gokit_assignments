package main

import (
	"errors"
	"fmt"
)

var data map[string]string

// init initializes the data map.
func init() {
	data = make(map[string]string)
}

// addItem adds an item to the data map if it doesn't already exist.
func addItem(k, v string) error {
	fmt.Printf("------------ Add-Item ------------\n")
	if _, ok := data[k]; ok {
		return errors.New("Add-Item: Item already exists\n")
	}
	data[k] = v
	fmt.Printf("Add-Item:\n  {Key: %s, Value: %s}\n\n", k, data[k])
	return nil
}

// updateItem updates the value of a given key in the data map.
func updateItem(k, v string) error {
	fmt.Printf("------------ Update-Item ------------\n")
	if _, ok := data[k]; ok {
		fmt.Printf("(Before):\n  Key: %s \n  Value: %s\n\n", k, data[k])
		data[k] = v
		fmt.Printf("(After):\n   Key: %s \n  Value: %s\n\n", k, data[k])
		return nil
	}
	return errors.New("Update-Item: Item does not exist\n")
}

// getById retrieves the value associated with the given key from the data map.
func getById(k string) error {
	fmt.Printf("------------ Get-Item-ById ------------\n")
	if _, ok := data[k]; !ok {
		return errors.New("Get-Item-ById: Item does not exist\n")
	}
	fmt.Printf("Get-Item-ById:\n  {Key: %s, Value: %s}\n\n", k, data[k])
	return nil
}

// getAll returns a map of strings to strings and an error.
func getAll() error {
	fmt.Printf("------------ Get-All-Items ------------\n")
	if len(data) == 0 {
		return errors.New("Get-All-Items: No items in map\n")
	}
	fmt.Printf("Get-All-Items:\n")
	for k, v := range data {
		fmt.Printf("  {Key: %s, Value: %s}\n", k, v)
	}
	fmt.Printf("\n")
	return nil
}

// deleteItem deletes an item from the data map based on the given key.
func deleteItem(k string) error {
	fmt.Printf("------------ Delete-Item ------------\n")
	if _, ok := data[k]; !ok {
		return errors.New("Delete-Item: Item does not exist")
	}
	fmt.Printf("Item-Deleted:\n\n  {Key: %s, Value: %s}\n\n", k, data[k])
	delete(data, k)
	return nil
}

func main() {
	err := addItem("item1", "value1")
	if err != nil {
		fmt.Println(err)
	}
	err = updateItem("item1", "value3")
	if err != nil {
		fmt.Println(err)
	}
	err = getById("item1")
	if err != nil {
		fmt.Println(err)
	}
	err = getAll()
	if err != nil {
		fmt.Println(err)
	}
	err = deleteItem("item1")
	if err != nil {
		fmt.Println(err)
	}
}
