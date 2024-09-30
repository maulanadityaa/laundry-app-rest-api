package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/maulanadityaa/laundry-app-rest-api/model/dto/request"
	"github.com/maulanadityaa/laundry-app-rest-api/model/dto/response"
	"github.com/maulanadityaa/laundry-app-rest-api/service"
	"github.com/maulanadityaa/laundry-app-rest-api/service/impl"
)

type TransactionController struct{}

var transactionService service.TransactionService = impl.NewTransactionService()

func NewTransactionController(g *gin.RouterGroup) {
	controller := new(TransactionController)

	transactionGroup := g.Group("/transactions")
	{
		transactionGroup.POST("", controller.AddTransaction)
		transactionGroup.PUT("", controller.UpdateTransaction)
		transactionGroup.GET("", controller.GetAllTransaction)
		transactionGroup.GET("/:id", controller.GetTransactionByID)
	}
}

func (TransactionController) AddTransaction(c *gin.Context) {
	var request request.TransactionRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		response.NewResponseBadRequest(c, err.Error())
		return
	}

	result, err := transactionService.AddTransaction(request)
	if err != nil {
		response.NewResponseError(c, err.Error())
		return
	}

	response.NewResponseCreated(c, result)
}

func (TransactionController) UpdateTransaction(c *gin.Context) {
	var request request.TransactionUpdateRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		response.NewResponseBadRequest(c, err.Error())
		return
	}

	result, err := transactionService.UpdateTransaction(request)
	if err != nil {
		response.NewResponseError(c, err.Error())
		return
	}

	response.NewResponseOK(c, result)
}

func (TransactionController) GetAllTransaction(c *gin.Context) {
	paging := c.DefaultQuery("paging", "1")
	rowsPerPage := c.DefaultQuery("rowsPerPage", "10")
	customerName := c.DefaultQuery("customerName", "")
	employeeName := c.DefaultQuery("employeeName", "")
	startDate := c.DefaultQuery("startDate", "")
	finishTime := c.DefaultQuery("finishTime", "")

	result, totalRows, totalPage, err := transactionService.GetAllTransaction(paging, rowsPerPage, customerName, employeeName, startDate, finishTime)
	if err != nil {
		response.NewResponseError(c, err.Error())
		return
	}

	response.NewResponseOKWithPaging(c, result, paging, rowsPerPage, totalRows, totalPage)
}

func (TransactionController) GetTransactionByID(c *gin.Context) {
	id := c.Param("id")

	result, err := transactionService.GetTransactionByID(id)
	if err != nil {
		response.NewResponseError(c, err.Error())
		return
	}

	response.NewResponseOK(c, result)
}