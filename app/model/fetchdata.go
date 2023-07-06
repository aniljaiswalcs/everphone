package model

import (
	"errors"

	reader "github.com/aniljaiswalcs/everphone/app/reader"
)

type EmployeeRepository interface {
	GetByName(employeeName string) (reader.Employee, error)
}

type GiftRepository interface {
	GetAll() ([]reader.Gift, error)
}
type MemoryEmployeeRepository struct {
	Employees []reader.Employee
}

func (r MemoryEmployeeRepository) GetByName(employeename string) (reader.Employee, error) {
	for _, employee := range r.Employees {
		if employee.Name == employeename {
			return employee, nil
		}
	}
	return reader.Employee{}, errors.New("employee not found")
}

type MemoryGiftRepostory struct {
	Gifts []reader.Gift
}

func (r MemoryGiftRepostory) GetAll() ([]reader.Gift, error) {
	return r.Gifts, nil
}
