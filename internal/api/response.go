package api

import (
	"github.com/gofiber/fiber/v2"
)

type APIError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Error   *APIError   `json:"error"`
}

func Success(c *fiber.Ctx, data interface{}, status ...int) error {
	code := fiber.StatusOK
	if len(status) > 0 {
		code = status[0]
	}
	return c.Status(code).JSON(Response{
		Success: true,
		Data:    data,
		Error:   nil,
	})
}

func Error(c *fiber.Ctx, status int, code string, message string) error {
	return c.Status(status).JSON(Response{
		Success: false,
		Data:    nil,
		Error: &APIError{
			Code:    code,
			Message: message,
		},
	})
}
