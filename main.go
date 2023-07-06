package main

import (
	"github.com/gin-gonic/gin"

	model "github.com/aniljaiswalcs/everphone/app/model"
	reader "github.com/aniljaiswalcs/everphone/app/reader"

	handlers "github.com/aniljaiswalcs/everphone/app/middleware"
)

const employeeFilePath = "./jsonfiles/employees.json"
const giftsFilePath = "./jsonfiles/gifts.json"

func main() {

	employeeRepo := &model.MemoryEmployeeRepository{
		Employees: reader.LoadEmployees(employeeFilePath),
	}

	giftRepo := &model.MemoryGiftRepostory{
		Gifts: reader.LoadGifts(giftsFilePath),
	}

	model.CreateMapOfInterestandCategory()
	assigner := &model.GiftAssignerImpl{
		EmployeeRepo: employeeRepo,
		GiftRepo:     giftRepo,
		Assignments:  make([]model.GiftAssignment, 0),
	}

	router := gin.Default()
	assignHandler := handlers.NewAssignHandler(assigner)
	router.GET("/assign-gift/:employeeName", assignHandler.HandleAssignGift)

	router.Run(":8088")
}
