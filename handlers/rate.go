package handlers

import (
	"github.com/go-fiber/models"
	"github.com/gofiber/fiber/v2"
)

type AddRateRequestBody struct {
	ClientID       string `json:"clientID"`
	Tokens         int    `json:"tokens"`
	Requests       int    `json:"requests"`
	ApiKey         int    `json:"apiKey"`
	TargetEndpoint string `json:"targetEndpoint"`
}

func (h handler) AddRate(c *fiber.Ctx) error {
	body := AddRateRequestBody{}

	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var rate models.Rate

	rate.ClientID = body.ClientID
	rate.Tokens = body.Tokens
	rate.Requests = body.Requests
	rate.ApiKey = body.ApiKey
	rate.TargetEndpoint = body.TargetEndpoint

	// insert new db entry
	if result := h.DB.Create(&rate); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(&rate)
}
