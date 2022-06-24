package handler

import (
	"billing"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) AddMoney(c *gin.Context) {
	//userId, err := getUserId(c)
	//if err != nil {
	//	return
	//}

	var input billing.Account
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err := h.services.Billing.AddMoney(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})

}

func (h *Handler) Balance(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	moneyAmount, err := h.services.Billing.Balance(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, moneyAmount)
}

func (h *Handler) Transfer(c *gin.Context) {
	var input billing.TransferInfo
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err := h.services.Billing.Transfer(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

func (h *Handler) GetAllAccs(c *gin.Context) {
	accounts, err := h.services.Billing.GetAllAccs()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, accounts)
}

//
//type getAllListsResponse struct {
//	Data []todo.TodoList `json:"data"`
//}
//
//func (h *Handler) getAllLists(c *gin.Context) {
//	userId, err := getUserId(c)
//	if err != nil {
//		return
//	}
//
//	lists, err := h.services.TodoList.GetAll(userId)
//	if err != nil {
//		newErrorResponse(c, http.StatusInternalServerError, err.Error())
//		return
//	}
//	c.JSON(http.StatusOK, getAllListsResponse{
//		Data: lists,
//	})
//}
//
//func (h *Handler) getListById(c *gin.Context) {
//	userId, err := getUserId(c)
//	if err != nil {
//		return
//	}
//
//	id, err := strconv.Atoi(c.Param("id"))
//	if err != nil {
//		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
//		return
//	}
//
//	list, err := h.services.TodoList.GetById(userId, id)
//	if err != nil {
//		newErrorResponse(c, http.StatusInternalServerError, err.Error())
//		return
//	}
//	c.IndentedJSON(http.StatusOK, list)
//}
//
//func (h *Handler) updateList(c *gin.Context) {
//	userId, err := getUserId(c)
//	if err != nil {
//		return
//	}
//
//	id, err := strconv.Atoi(c.Param("id"))
//	if err != nil {
//		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
//		return
//	}
//	var input todo.UpdateListInput
//	if err := c.BindJSON(&input); err != nil {
//		newErrorResponse(c, http.StatusBadRequest, err.Error())
//		return
//	}
//	if err := h.services.TodoList.Update(userId, id, input); err != nil {
//		newErrorResponse(c, http.StatusInternalServerError, err.Error())
//		return
//	}
//	c.JSON(http.StatusOK, statusResponse{
//		Status: "ok",
//	})
//}
//
//func (h *Handler) deleteList(c *gin.Context) {
//	userId, err := getUserId(c)
//	if err != nil {
//		return
//	}
//
//	id, err := strconv.Atoi(c.Param("id"))
//	if err != nil {
//		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
//		return
//	}
//
//	err = h.services.TodoList.Delete(userId, id)
//	if err != nil {
//		newErrorResponse(c, http.StatusInternalServerError, err.Error())
//		return
//	}
//	c.IndentedJSON(http.StatusOK, statusResponse{
//		Status: "ok",
//	})
//}
