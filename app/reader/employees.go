package reader

import (
	"encoding/json"
	"io/ioutil"
)

type Employee struct {
	Name      string
	Interests []string
}

type EmployeeRepository interface {
	GetByName(employeeName string) (*Employee, error)
}

func LoadEmployees(filename string) []Employee {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	var employees []Employee
	err = json.Unmarshal(data, &employees)
	if err != nil {
		panic(err)
	}

	return employees
}
