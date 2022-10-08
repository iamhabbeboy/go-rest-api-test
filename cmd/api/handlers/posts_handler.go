package handlers

import (
	"net/http"
	"strconv"

	"github.com/iamhabbeboy/go-rest-api-test/cmd/api/service"
	"github.com/labstack/echo"
)

func PostIndexHandler(c echo.Context) error {
	data, err := service.GetAll()
	if err != nil {
		c.String(http.StatusBadGateway, "Unable to process data")
	}
	return c.JSON(http.StatusOK, data)
}

func PostSingleHandler(c echo.Context) error {
	id := c.Param("id")
	idx, err := strconv.Atoi(id)
	if err != nil {
		c.String(http.StatusBadGateway, "Unable to process data")
	}
	data, err := service.GetById(idx)
	if err != nil {
		c.String(http.StatusBadGateway, "Unable to process data")
	}
	return c.JSON(http.StatusOK, data)
}
