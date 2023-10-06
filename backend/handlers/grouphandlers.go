package handlers

import (
	"github.com/gofiber/fiber/v2"

	"lovelcode/utils"
)

func ApiOnly(c *fiber.Ctx) error{
	
	ct, ok := c.GetReqHeaders()["Content-Type"]
	if ct=="application/json" && ok==true{
		return c.Next()
	}
	return c.Status(400).JSON(fiber.Map{"error":"Content-Type must be application/json"})
}

func AuthRequired(c *fiber.Ctx) error{
	token := c.Cookies("token", "")
	if token==""{
		return c.Status(401).JSON(fiber.Map{"error":"authentication required"})
	}
	user, err := utils.VerifyToken(token)
	if err!=nil{
		return c.Status(401).JSON(fiber.Map{"error":"token invalid"})
	}

	c.Locals("user", user)
	return c.Next()
}