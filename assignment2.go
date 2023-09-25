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
	fmt.Println("------------ Add-Item ------------\n\n")
	if _, ok := data[k]; ok {
		return errors.New("Add-Item: Item already exists\n\n")
	} else {
		data[k] = v
		fmt.Printf("Add-Item:\n\n  {Key: %s, Value: %s}\n\n", k, data[k])
		return nil
	}

}

// updateItem updates the value of a given key in the data map.
func updateItem(k, v string) error {
	fmt.Println("------------ Update-Item ------------\n\n")
	if _, ok := data[k]; ok {
		fmt.Printf("(Before):\n  Key: %s \n  Value: %s\n\n", k, data[k])
		data[k] = v
		fmt.Printf("(After):\n   Key: %s \n  Value: %s\n\n", k, data[k])
		return nil
	} else {
		return errors.New("Update-Item: Item does not exist\n\n")
	}

}

// getById retrieves the value associated with the given key from the data map.
func getById(k string) error {
	fmt.Println("------------ Get-Item-ById ------------\n\n")
	if _, ok := data[k]; !ok {
		return errors.New("Get-Item-ById: Item does not exist\n\n")
	} else {
		fmt.Printf("Get-Item-ById:\n\n  {Key: %s, Value: %s}\n\n", k, data[k])
		return nil
	}

}

// getAll returns a map of strings to strings and an error.
func getAll() error {
	fmt.Println("------------ Get-All-Items ------------\n\n")
	if len(data) == 0 {
		return errors.New("Get-All-Items: No items in map\n\n")
	} else {
		fmt.Println("Get-All-Items:\n")
		for k, v := range data {
			fmt.Printf("  {Key: %s, Value: %s}\n", k, v)
		}
		fmt.Println("\n")
		return nil
	}
}

// deleteItem deletes an item from the data map based on the given key.
func deleteItem(k string) error {
	fmt.Println("------------ Delete-Item ------------\n\n")
	if _, ok := data[k]; !ok {
		return errors.New("Delete-Item: Item does not exist")
	} else {
		fmt.Printf("Item-Deleted:\n\n  {Key: %s, Value: %s}\n\n", k, data[k])
		delete(data, k)
		return nil
	}

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
