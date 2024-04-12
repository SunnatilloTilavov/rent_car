package handler

import (
	_ "clone/rent_car_us/api/docs"
	"clone/rent_car_us/api/models"
	"fmt"
	"net/http"
	"clone/rent_car_us/pkg/check"

	"github.com/gin-gonic/gin"
)

// CustomerLogin godoc
// @Router       /customer/login [POST]
// @Summary      Customer login
// @Description  Customer login
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        login body models.CustomerLoginRequest true "login"
// @Success      201  {object}  models.CustomerLoginResponse
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h *Handler) CustomerLogin(c *gin.Context) {
	loginReq := models.CustomerLoginRequest{}

	if err := c.ShouldBindJSON(&loginReq); err != nil {
		handleResponse(c, h.Log, "error while binding body", http.StatusBadRequest, err)
		return
	}
	fmt.Println("loginReq: ",loginReq)

	if err := check.ValidatePassword(loginReq.Password); err != nil {
		handleResponse(c,h.Log,"error while validating  old password,old password: "+loginReq.Password, http.StatusBadRequest, err.Error())
		return
	}
	
	loginResp, err := h.Services.Auth().CustomerLogin(c.Request.Context(), loginReq)
	if err != nil {
		handleResponse(c, h.Log, "unauthorized", http.StatusUnauthorized, err)
		return
	}

	handleResponse(c, h.Log, "Succes", http.StatusOK, loginResp)

}
