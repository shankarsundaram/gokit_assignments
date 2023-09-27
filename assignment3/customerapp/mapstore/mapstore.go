package mapstore

import (
	"customerapp/domain"
	"errors"
)

type MapStore struct {
	store map[string]domain.Customer
}

func NewMapStore() *MapStore {
	return &MapStore{store: make(map[string]domain.Customer)}
}

// Create will inserts the record into mem datastore.
func (m *MapStore) Create(customer domain.Customer) error {
	if _, ok := m.store[customer.ID]; ok {
		return errors.New("customer record already exists")
	}
	m.store[customer.ID] = customer

	return nil
}

func (m *MapStore) GetById(s string) (domain.Customer, error) {
	if _, ok := m.store[s]; !ok {
		return domain.Customer{}, errors.New("customer record doesn't exists")
	}

	return m.store[s], nil
}

func (m *MapStore) Update(s string, customer domain.Customer) error {
	if _, ok := m.store[s]; !ok {
		return errors.New("customer record doesn't exists")
	}

	m.store[s] = customer
	return nil
}

func (m *MapStore) Delete(s string) error {
	if _, ok := m.store[s]; !ok {
		return errors.New("customer record doesn't exists")
	}
	delete(m.store, s)
	return nil
}

func (m *MapStore) GetAll() ([]domain.Customer, error) {
	var customers []domain.Customer
	for _, customer := range m.store {
		customers = append(customers, customer)
	}
	return customers, nil
}
