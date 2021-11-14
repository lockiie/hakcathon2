package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lucas/hackathon/src/models"
)

const (
	errJSON = "Error processing JSON"
	paramID = "id"
	status  = "active"
	title   = "title"
	proType = "type_product"
	name    = "name"
	father  = "father_id"
	all     = "all"
)

func tryErrorDB(c *fiber.Ctx, err string, intError string) error {
	sError := models.ErrorsDB{}
	sError.InternalMessage = intError
	sError.Message = err + intError
	return c.Status(fiber.StatusBadRequest).JSON(sError)
}
