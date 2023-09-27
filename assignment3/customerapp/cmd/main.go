package main

import (
	"fmt"

	"customerapp/domain"
	"customerapp/mapstore"
)

// CustomerController Organises the CRUD operations at UI layer.
type CustomerController struct {
	// Explicit dependency that hides dependent logic
	store domain.CustomerStore // CustomerStore value

}

// Implementation of Create Method of Interface CustomerStore
func (c CustomerController) Add(cr domain.Customer) {
	err := c.store.Create(cr)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("New Customer has been created")
}

// Implementation of Update Method of Interface CustomerStore
func (c CustomerController) Update(s string, cr domain.Customer) {
	err := c.store.Update(s, cr)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Customer has been Updated")
}

// Implementation of GetById Method of Interface CustomerStore
func (c CustomerController) GetById(s string) {
	result, err := c.store.GetById(s)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("GetById:", result)
}

// Implementation of Delete Method of Interface CustomerStore
func (c CustomerController) Delete(s string) {
	err := c.store.Delete(s)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Customer has been Deleted")
}

// Implementation of GetAll Method of Interface CustomerStore
func (c CustomerController) GetAll() {
	result, err := c.store.GetAll()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("GetAll:", result)
}

// EntryPoint
func main() {
	controller := CustomerController{
		store: mapstore.NewMapStore(), // Inject the dependency
		// store : mongodb.NewMongoStore(), // with another database
	}
	customer := domain.Customer{
		ID:    "cust101",
		Name:  "JPM",
		Email: "jpm@jpm",
	}
	customer1 := domain.Customer{
		ID:    "cust102",
		Name:  "APM",
		Email: "apm@apm",
	}
	customerUpdated := domain.Customer{
		ID:    "cust101",
		Name:  "JPM",
		Email: "shankar@gmail",
	}
	// Create
	controller.Add(customer)
	controller.Add(customer1)
	// Update
	controller.Update("cust101", customerUpdated)
	// GetById
	controller.GetById("cust101")
	// GetAll
	controller.GetAll()
	// Delete
	controller.Delete("cust102")
}
