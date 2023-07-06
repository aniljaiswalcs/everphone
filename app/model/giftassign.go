package model

import (
	"errors"
	"fmt"
	"sync"

	reader "github.com/aniljaiswalcs/everphone/app/reader"
)

type EmployeeGiftAssigner interface {
	AssignGift(employeeName string) (string, error)
}

type GiftAssignment struct {
	EmployeeName string
	GiftName     string
}
type GiftAssignerImpl struct {
	EmployeeRepo EmployeeRepository
	GiftRepo     GiftRepository
	Assignments  []GiftAssignment
	mutex        sync.Mutex
}

func (a *GiftAssignerImpl) AssignGift(employeename string) (string, error) {
	a.mutex.Lock()
	defer a.mutex.Unlock()

	// Check if the employee has already been assigned a gift
	for _, assignment := range a.Assignments {
		if assignment.EmployeeName == employeename {
			return fmt.Sprintf("%s: gift is assigned to you is %s", assignment.EmployeeName, assignment.GiftName), nil
		}
	}
	employee, err := a.EmployeeRepo.GetByName(employeename)
	if err != nil {
		return "", err
	}
	gifts, err := a.GiftRepo.GetAll()
	if err != nil {
		return "", err
	}

	// Find a suitable gift for the employee
	for _, gift := range gifts {
		if canAssignGiftToEmployee(gift, employee) {
			// Check if the gift has already been assigned
			if !isGiftAlreadyAssigned(gift.Name, a.Assignments) {
				// Assign the gift to the employee
				assignment := GiftAssignment{
					EmployeeName: employeename,
					GiftName:     gift.Name,
				}
				a.Assignments = append(a.Assignments, assignment)
				return gift.Name, nil
			}
		}
	}

	// call comes here meaning no direct interest and gift found
	for _, interest := range employee.Interests {
		category := getCategoryfrominterest(interest)
		if len(category) == 0 {
			continue
		} else {
			sliceofsimilarinterest := GetSliceofSimilarInterestFromCategory(category)
			for _, interest := range sliceofsimilarinterest {
				for _, gift := range gifts {
					for _, category := range gift.Categories {
						//interest match and if not assigned to anyone then assign
						if interest == category && !isGiftAlreadyAssigned(gift.Name, a.Assignments) {
							// Assign the gift to the employee
							assignment := GiftAssignment{
								EmployeeName: employeename,
								GiftName:     gift.Name,
							}
							a.Assignments = append(a.Assignments, assignment)
							return gift.Name, nil
						}
					}
				}
			}
		}
	}

	//still nothing then no one can give him gift!!
	return "", errors.New("sorry, no suitable gift found for the employee")
}

func canAssignGiftToEmployee(gift reader.Gift, employee reader.Employee) bool {
	for _, interest := range employee.Interests {
		for _, category := range gift.Categories {
			if interest == category {
				return true
			}
		}
	}
	return false
}

func isGiftAlreadyAssigned(giftName string, assignments []GiftAssignment) bool {
	for _, assignment := range assignments {
		if assignment.GiftName == giftName {
			return true
		}
	}
	return false
}
