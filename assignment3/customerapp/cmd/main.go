package main

import (
	"customerapp/domain"
	"customerapp/mapstore"
	"fmt"
)

// CustomerController Organises the CRUD operations at UI layer.
type CustomerController struct {
	// Explicit dependency that hides dependent logic
	store domain.CustomerStore // CustomerStore value

}

func (c CustomerController) Add(cr domain.Customer) {
	err := c.store.Create(cr)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("New Customer has been created")
}

func (c CustomerController) Update(s string, cr domain.Customer) {
	err := c.store.Update(s, cr)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Customer has been Updated")
}

func (c CustomerController) GetById(s string) {
	result, err := c.store.GetById(s)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("GetById:", result)
}

func (c CustomerController) Delete(s string) {
	err := c.store.Delete(s)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Customer has been Deleted")
}

func (c CustomerController) GetAll() {
	result, err := c.store.GetAll()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("GetAll:", result)
}

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
