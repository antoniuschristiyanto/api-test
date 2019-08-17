package services

import (
	"api-test/models"
	"database/sql"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type Result map[string]interface{}

func GetUser(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, models.GetUser(db))
	}
}

func InsertUser(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {

		var user models.User

		c.Bind(&user)

		_, err := models.InsertUser(db, user.Name, user.Email, user.Handphone, user.Status)

		if err != nil {
			return c.JSON(http.StatusBadRequest, Result{
				"result":  err,
				"message": "nil",
			})
		}

		return c.JSON(http.StatusCreated, Result{
			"result":  user,
			"message": "success created user",
		})

	}
}

func UpdateHandphoneUser(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {

		id, _ := strconv.Atoi(c.Param("id"))

		var user models.User

		c.Bind(&user)

		_, err := models.UpdateHandphoneUser(db, id, user.Handphone)

		if err != nil {
			return c.JSON(http.StatusBadRequest, Result{
				"result":  err,
				"message": "nil",
			})
		}

		return c.JSON(http.StatusOK, Result{
			"result":  user.Handphone,
			"message": "success update handphone user",
		})
	}
}

func UpdateStatusUser(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {

		id, _ := strconv.Atoi(c.Param("id"))

		var user models.User

		c.Bind(&user)

		_, err := models.UpdateStatusUser(db, id, user.Status)

		if err != nil {
			return c.JSON(http.StatusBadRequest, Result{
				"result":  err,
				"message": "nil",
			})
		}

		return c.JSON(http.StatusOK, Result{
			"result":  user.Status,
			"message": "success update status user",
		})
	}
}

func DeleteUser(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))

		_, err := models.DeleteUser(db, id)

		if err != nil {
			return c.JSON(http.StatusBadRequest, Result{
				"result":  err,
				"message": "nil",
			})
		}

		return c.JSON(http.StatusOK, Result{
			"message": "success deleted user",
		})

	}
}
