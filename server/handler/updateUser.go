package handler

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
	"github.com/elias506/EchoRestAPI/models"
	repModels "github.com/elias506/EchoRestAPI/repository/models"
)

func UpdateUser(c echo.Context) error {
	cc := c.(*CustomContext)
	userID, err := strconv.Atoi(cc.Param("id"))
	if err != nil {
		fmt.Println(err)
		return err
	}
	reqUser := &models.RequestUser{}
	if err := cc.Bind(reqUser); err != nil {
		fmt.Println(err)
		return err
	}
	updatedUserID, err := repModels.IUpdateUser(cc.IUserDB, userID, reqUser)
	if err != nil {
		fmt.Println(err)
		return err
	}
	if updatedUserID == -1 {
		return cc.JSON(http.StatusNotFound, userID)
	}

	return cc.JSON(http.StatusOK, updatedUserID)
}