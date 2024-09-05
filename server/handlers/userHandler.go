package handlers

import (
	"echoapp/db"
	"echoapp/entities"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	*echo.Group
	db *db.DB
}

func NewUserHandler(g *echo.Group, db *db.DB) {
	h := &UserHandler{g, db}

	h.POST("/add", h.AddUser)
	h.GET("", h.GetUsers)
	h.GET("/", h.GetUserByEmail)
	h.PUT("/update", h.UpdateUser)
	h.DELETE("/delete/", h.DeleteUser)
}

func (h *UserHandler) AddUser(c echo.Context) error {
	var user entities.User

	err := c.Bind(&user)

	if err != nil {
		return c.JSON(http.StatusBadRequest, "bad  request")
	}

	err = h.db.AddUser(&user)

	if err != nil {
		return c.JSON(http.StatusBadRequest, "bad request")
	}

	return c.JSON(http.StatusOK, user)
}

func (h *UserHandler) GetUsers(c echo.Context) error {
	users, err := h.db.GetUsers()

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, users)
}

func (h *UserHandler) GetUserByEmail(c echo.Context) error {
	email := c.QueryParam("email")
	
	if email == "" {
		return c.JSON(http.StatusBadRequest, "bad request")
	}

	user, err := h.db.GetUser(email)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, user)
}

func (h *UserHandler) UpdateUser(c echo.Context) error {
	var user *entities.User

	err := c.Bind(&user)

	if err != nil {
		return c.JSON(http.StatusBadRequest, "bad request")
	}

	err = h.db.UpdateUser(user)

	if err != nil {
		return c.JSON(http.StatusNotFound, "not found")
	}

	return c.JSON(http.StatusOK, user)
}

func (h *UserHandler) DeleteUser(c echo.Context) error {
	email := c.QueryParam("email")

	if email == "" {
		return c.JSON(http.StatusBadRequest, "bad request")
	}

	err := h.db.DeleteUser(email)

	if err != nil {
		return c.JSON(http.StatusNotFound, "not found")
	}

	return c.JSON(http.StatusOK, "deleted")
}
