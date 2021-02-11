package handler

import (
	. "github.com/elias506/EchoRestAPI/models"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

func GetUser(c echo.Context) error {
	cc := c.(*CustomContext)

	userID, err := strconv.Atoi(cc.Param("id"))
	if err != nil {
		return err
	}
	user, err := IGetUser(cc.UserDB, userID)
	if err != nil {
		return err
	}
	return cc.JSON(http.StatusOK, user)
}
