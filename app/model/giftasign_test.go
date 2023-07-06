package model

import (
	"errors"
	"testing"

	reader "github.com/aniljaiswalcs/everphone/app/reader"
)

type mockEmployeeRepository struct {
	employees []reader.Employee
}

func (r mockEmployeeRepository) GetByName(name string) (reader.Employee, error) {
	for _, employee := range r.employees {
		if employee.Name == name {
			return employee, nil
		}
	}
	return reader.Employee{}, errors.New("Employee not found")
}

type mockGiftRepository struct {
	gifts []reader.Gift
}

func (r *mockGiftRepository) GetAll() ([]reader.Gift, error) {
	return r.gifts, nil
}

func TestGiftAssigner_AssignGift(t *testing.T) {
	// Test cases
	testCases := []struct {
		name          string
		employeeName  string
		employees     []reader.Employee
		gifts         []reader.Gift
		expectedGift  string
		expectedError bool
	}{
		{
			name:         "Assign gift to employee with matching interests",
			employeeName: "John",
			employees: []reader.Employee{
				{
					Name:      "John",
					Interests: []string{"music", "film"},
				},
			},
			gifts: []reader.Gift{
				{
					Name:       "Movie tickets",
					Categories: []string{"film"},
				},
				{
					Name:       "Headphones",
					Categories: []string{"music"},
				},
			},
			expectedGift:  "Movie tickets",
			expectedError: false,
		},
		{
			name:         "Assign gift to employee with matching interests",
			employeeName: "Jane",
			employees: []reader.Employee{
				{
					Name:      "Jane",
					Interests: []string{"sports"},
				},
			},
			gifts: []reader.Gift{
				{
					Name:       "Football",
					Categories: []string{"sports"},
				},
			},
			expectedGift:  "Football",
			expectedError: false,
		},
		{
			name:         "No gift to employee with no interests",
			employeeName: "Tim",
			employees: []reader.Employee{
				{
					Name:      "Tim",
					Interests: []string{"sports"},
				},
			},
			gifts: []reader.Gift{
				{
					Name:       "Rock",
					Categories: []string{"music"},
				},
			},
			expectedGift:  "",
			expectedError: true,
		},
		{
			name:         "employee not in database",
			employeeName: "Ram",
			employees: []reader.Employee{
				{
					Name:      "Jey",
					Interests: []string{"sports"},
				},
			},
			gifts: []reader.Gift{
				{
					Name:       "Rock",
					Categories: []string{"music"},
				},
			},
			expectedGift:  "",
			expectedError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			employeeRepo := mockEmployeeRepository{
				employees: tc.employees,
			}

			giftRepo := &mockGiftRepository{
				gifts: tc.gifts,
			}

			assigner := &GiftAssignerImpl{
				EmployeeRepo: employeeRepo,
				GiftRepo:     giftRepo,
				Assignments:  make([]GiftAssignment, 0),
			}

			gift, err := assigner.AssignGift(tc.employeeName)

			if (err != nil) != tc.expectedError {
				t.Errorf("Expected error: %v, got: %v", tc.expectedError, err != nil)
			}

			if gift != tc.expectedGift {
				t.Errorf("Expected gift: %s, got: %s", tc.expectedGift, gift)
			}
		})
	}
}
